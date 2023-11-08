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
	DeletedAt     soft_delete.DeletedAt `gorm:"softDelete:nano;index:idx_nondeleted_usercontext_workspace,unique"`
	UserContext   *UserContext          `gorm:"foreignkey:UserContextID"`
	UserContextID uint                  `gorm:"index:idx_nondeleted_usercontext_workspace,unique"`
	Name          string                `gorm:"index:idx_nondeleted_usercontext_workspace,unique"`
	About         string
	Frames        []*Frame
	Directories   []*Directory
	Repositories  []*Repository `gorm:"-"`
	Framers       []*Framer     `gorm:"-"`
	Shapers       []*Shaper     `gorm:"-"`
}

type WorkspaceInspector struct {
	Name         string
	About        string
	Frames       []*FrameInspector
	Repositories []*RepositoryInspector
	Directories  []*DirectoryInspector
}

// Construction

func NewWorkspace(uc *UserContext, name string) (w *Workspace) {
	w = &Workspace{
		UserContext: uc,
		Name:        name,
	}
	return
}

func CreateWorkspace(uc *UserContext, name string, about string) (w *Workspace, err error) {
	w = NewWorkspace(uc, name)
	w.About = about
	if w.IsExists() {
		err = app.Error("workspace %s already exists", w.Name)
		return
	}
	w.Create()
	return
}

func ResolveWorkspace(uc *UserContext, name string, loads ...string) (w *Workspace, err error) {
	if uc == nil {
		err = app.Error("no user context")
		return
	}
	if name == "" {
		w = uc.Workspace
		if w == nil {
			err = app.Error("no workspace context")
			return
		}
	} else {
		if len(uc.Workspaces) == 0 {
			err = app.Error("no workspaces exist")
			return
		}
		w = uc.FindWorkspace(name)
		if w == nil {
			err = app.Error("workspace %s does not exist", name)
			return
		}
	}
	if len(loads) > 0 {
		if err = w.Load(loads...); err != nil {
			return
		}
	}
	return
}

// Inspection

func (w *Workspace) Inspect() (wi *WorkspaceInspector) {
	ris := w.RepositoriesInspect()
	dis := w.DirectoriesInspect()
	fis := w.FramesInspect()
	wi = &WorkspaceInspector{
		Name:         w.Name,
		About:        w.About,
		Frames:       fis,
		Repositories: ris,
		Directories:  dis,
	}
	return
}

func (w *Workspace) FramesInspect() (fis []*FrameInspector) {
	fis = []*FrameInspector{}
	for _, f := range w.Frames {
		fi := f.Inspect()
		fis = append(fis, fi)
	}
	return
}

func (w *Workspace) RepositoriesInspect() (ris []*RepositoryInspector) {
	ris = []*RepositoryInspector{}
	for _, r := range w.Repositories {
		ri := r.Inspect()
		ris = append(ris, ri)
	}
	return
}

func (w *Workspace) DirectoriesInspect() (dis []*DirectoryInspector) {
	dis = []*DirectoryInspector{}
	for _, d := range w.Directories {
		di := d.Inspect()
		dis = append(dis, di)
	}
	return
}

// Data

func (w *Workspace) IsExists() (is bool) {
	if w.UserContext.FindWorkspace(w.Name) != nil {
		is = true
		return
	}
	return
}

func (w *Workspace) directory() (dirPath string) {
	dirPath = utils.DataDir(filepath.Join("workspaces", w.Name))
	return
}

func (w *Workspace) Load(loads ...string) (err error) {
	wl := NewWorkspaceLoader(w, loads)
	err = wl.load()
	return
}

func (w *Workspace) Assign(params map[string]any) {
	if params["Name"] != nil {
		w.Name = params["Name"].(string)
	}
	if params["About"] != nil {
		w.About = params["About"].(string)
	}
}

func (w *Workspace) Create() {
	queries.Create(w)
}

func (w *Workspace) Save() {
	queries.Update(w)
}

func (w *Workspace) Destroy() {
	queries.Delete(w)
}

// Associations

func (w *Workspace) GitRepoDirectoryFor(uri string) (dirPath string) {
	elem := []string{}
	elem = append(elem, w.directory(), "repos")
	elem = append(elem, strings.Split(uri, "/")...)
	dirPath = filepath.Join(elem...)
	return
}

func (w *Workspace) FindDirectory(path string) (d *Directory) {
	for _, d = range w.Directories {
		if d.Path == path {
			return
		}
	}
	d = nil
	return
}

func (w *Workspace) FindRepository(path string) (r *Repository) {
	for _, r = range w.Repositories {
		if r.Path == path {
			return
		}
	}
	r = nil
	return
}

func (w *Workspace) FindFrame(name string) (f *Frame) {
	for _, f = range w.Frames {
		if f.Name == name {
			return
		}
	}
	f = nil
	return
}

func (w *Workspace) FindFramer(name string) (fr *Framer) {
	for _, fr = range w.Framers {
		if fr.Name == name {
			return
		}
	}
	fr = nil
	return
}

func (w *Workspace) FindShaper(name string) (sr *Shaper) {
	for _, sr = range w.Shapers {
		if sr.Name == name {
			return
		}
	}
	sr = nil
	return
}
