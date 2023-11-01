package models

import (
	"path/filepath"
	"sf/app"
	"sf/database/queries"
	"sf/utils"
	"strings"
	"time"

	"gorm.io/plugin/soft_delete"
)

type Workspace struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     soft_delete.DeletedAt `gorm:"index:idx_nondeleted_usercontext_workspace,unique"`
	UserContext   *UserContext          `gorm:"foreignkey:UserContextID"`
	UserContextID uint                  `gorm:"index:idx_nondeleted_usercontext_workspace,unique"`
	Name          string                `gorm:"index:idx_nondeleted_usercontext_workspace,unique"`
	About         string
	Frames        []*Frame
	Directories   []*Directory
}

type WorkspaceInspector struct {
	Name         string
	About        string
	Frames       []*FrameInspector
	Repositories []*RepositoryInspector
	Directories  []*DirectoryInspector
}

// Construction

func WorkspaceNew(uc *UserContext, name string) (w *Workspace) {
	if uc == nil {
		panic("Workspace UserContext is <nil>")
	}
	w = &Workspace{UserContext: uc, Name: name}
	return
}

// Inspection

func (w *Workspace) Inspect() (wi *WorkspaceInspector, err error) {
	var ris []*RepositoryInspector
	var dis []*DirectoryInspector
	if ris, err = w.RepositoriesInspect(); err != nil {
		return
	}
	if dis, err = w.DirectoriesInspect(); err != nil {
		return
	}
	wi = &WorkspaceInspector{
		Name:         w.Name,
		About:        w.About,
		Frames:       w.FramesInspect(),
		Repositories: ris,
		Directories:  dis,
	}
	return
}

func (w *Workspace) FramesInspect() (fis []*FrameInspector) {
	fis = []*FrameInspector{}
	for _, f := range w.Frames {
		fis = append(fis, f.Inspect())
	}
	return
}

func (w *Workspace) RepositoriesInspect() (ris []*RepositoryInspector, err error) {
	ris = []*RepositoryInspector{}
	for _, r := range w.RepositoriesWithGitRepos() {
		var ri *RepositoryInspector
		if ri, err = r.Inspect(); err != nil {
			return
		}
		ris = append(ris, ri)
	}
	return
}

func (w *Workspace) DirectoriesInspect() (dis []*DirectoryInspector, err error) {
	dis = []*DirectoryInspector{}
	for _, d := range w.DirectoriesWithGitRepos() {
		var di *DirectoryInspector
		if di, err = d.Inspect(); err != nil {
			return
		}
		dis = append(dis, di)
	}
	return
}

// Data

func (w *Workspace) IsExists() (is bool) {
	if w.UserContext.WorkspaceFind(w.Name) != nil {
		is = true
		return
	}
	return
}

func (w *Workspace) directory() (dirPath string) {
	dirPath = utils.DataDir(filepath.Join("workspaces", w.Name))
	return
}

func (w *Workspace) Load(preloads ...string) (err error) {
	queries.Load(w, w.ID, preloads...)
	return
}

func (w *Workspace) Assign(params map[string]any) {
	if params["Name"] != nil {
		w.Name = utils.StringTidy(params["Name"].(string))
	}
	if params["About"] != nil {
		w.About = utils.StringTidy(params["About"].(string))
	}
}

func (w *Workspace) Create() (err error) {
	if w.IsExists() {
		err = app.Error(nil, "workspace %s already exists", w.Name)
		return
	}
	queries.Create(w)
	return
}

func (w *Workspace) Save() (err error) {
	queries.Update(w)
	return
}

func (w *Workspace) Destroy() {
	queries.Delete(w)
}

// Associations

func (w *Workspace) Lookup(assocName string, key string, value string, model any) {
	queries.Lookup(w, assocName, key, value, model)
}

func (w *Workspace) GitRepoDirectoryFor(uri string) (dirPath string) {
	elem := []string{}
	elem = append(elem, w.directory(), "repos")
	elem = append(elem, strings.Split(uri, "/")...)
	dirPath = filepath.Join(elem...)
	return
}

func (w *Workspace) RepositoriesWithGitRepos() (rs []*Repository) {
	for _, dp := range w.RepositoryDirs() {
		r := RepositoryNew(w, dp)
		r.Load()
		rs = append(rs, r)
	}
	return
}

func (w *Workspace) RepositoryDirs() (dps []string) {
	dps = utils.GitRepoDirs(filepath.Join(w.directory(), "repos"))
	return
}

func (w *Workspace) DirectoriesWithGitRepos() (ds []*Directory) {
	for _, d := range w.Directories {
		d = DirectoryNew(w, d.Path)
		d.Load()
		ds = append(ds, d)
	}
	return
}

func (w *Workspace) Shapes() (ss []*Shape) {
	ss = []*Shape{}
	for _, f := range w.Frames {
		ss = append(ss, f.Shapes...)
	}
	return
}

func (w *Workspace) Framers() (frs []*Framer, err error) {
	frs = []*Framer{}
	var dfrs []*Framer
	var rfrs []*Framer
	for _, d := range w.DirectoriesWithGitRepos() {
		if dfrs, err = d.Framers(); err != nil {
			return
		}
		frs = append(frs, dfrs...)
	}
	for _, r := range w.RepositoriesWithGitRepos() {
		if rfrs, err = r.Framers(); err != nil {
			return
		}
		frs = append(frs, rfrs...)
	}
	return
}

func (w *Workspace) Shapers() (srs []*Shaper, err error) {
	srs = []*Shaper{}
	var dsrs []*Shaper
	var rsrs []*Shaper
	for _, d := range w.DirectoriesWithGitRepos() {
		if dsrs, err = d.Shapers(); err != nil {
			return
		}
		srs = append(srs, dsrs...)
	}
	for _, r := range w.RepositoriesWithGitRepos() {
		if rsrs, err = r.Shapers(); err != nil {
			return
		}
		srs = append(srs, rsrs...)
	}
	return
}

func (w *Workspace) DirectoryFind(path string) (d *Directory) {
	d = &Directory{}
	w.Lookup("Directories", "path", path, d)
	if d.ID == uint(0) {
		d = nil
	}
	return
}

func (w *Workspace) RepositoryFind(path string) (r *Repository) {
	r = RepositoryNew(w, path)
	// r.Load()
	if !r.IsExists() {
		r = nil
	}
	return
}

func (w *Workspace) FrameFind(name string) (f *Frame) {
	f = &Frame{}
	w.Lookup("Frames", "name", name, f)
	if f.ID == uint(0) {
		f = nil
	}
	return
}

func (w *Workspace) FramerFind(name string) (fr *Framer, err error) {
	var frs []*Framer
	if frs, err = w.Framers(); err != nil {
		return
	}
	for _, fr = range frs {
		if fr.Name == name {
			break
		} else {
			fr = nil
		}
	}
	return
}

func (w *Workspace) ShaperFind(name string) (sr *Shaper, err error) {
	var srs []*Shaper
	if srs, err = w.Shapers(); err != nil {
		return
	}
	for _, sr = range srs {
		if sr.Name == name {
			break
		} else {
			sr = nil
		}
	}
	return
}
