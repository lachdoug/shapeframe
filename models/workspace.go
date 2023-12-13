package models

import (
	"path/filepath"
	"sf/app/errors"
	"sf/app/validations"
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
	UserContextID uint                  `gorm:"index:idx_nondeleted_usercontext_workspace,unique"`
	UserContext   *UserContext          `gorm:"foreignkey:UserContextID"`
	Name          string                `gorm:"index:idx_nondeleted_usercontext_workspace,unique"`
	About         string
	Frames        []*Frame
	Directories   []*Directory
	Repositories  []*Repository `gorm:"-"`
	Framers       []*Framer     `gorm:"-"`
	Shapers       []*Shaper     `gorm:"-"`
}

type WorkspaceReader struct {
	Name         string
	About        string
	Frames       []string
	Repositories []string
	Directories  []string
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

func CreateWorkspace(uc *UserContext, name string, about string) (w *Workspace, vn *validations.Validation, err error) {
	w = NewWorkspace(uc, name)
	w.About = about
	if w.IsExists() {
		err = errors.Errorf("workspace %s already exists", w.Name)
		return
	}
	if vn = w.Validation(); vn.IsValid() {
		w.Save()
	}
	return
}

func ResolveWorkspace(uc *UserContext, name string, loads ...string) (w *Workspace, err error) {
	if uc == nil {
		err = errors.Error("no user context")
		return
	}
	if name == "" {
		w = uc.Workspace
		if w == nil {
			err = errors.Error("no workspace context")
			return
		}
	} else {
		if len(uc.Workspaces) == 0 {
			err = errors.Error("no workspaces exist")
			return
		}
		w = uc.FindWorkspace(name)
		if w == nil {
			err = errors.Errorf("workspace %s does not exist", name)
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

// Read

func (w *Workspace) Read() (wr *WorkspaceReader) {
	wr = &WorkspaceReader{
		Name:         w.Name,
		About:        w.About,
		Frames:       w.FrameNames(),
		Repositories: w.RepositoryUris(),
		Directories:  w.DirectoryPaths(),
	}
	return
}

// Inspection

func (w *Workspace) Inspect() (wi *WorkspaceInspector, err error) {
	var fis []*FrameInspector
	var ris []*RepositoryInspector
	var dis []*DirectoryInspector
	if ris, err = w.RepositoriesInspect(); err != nil {
		return
	}
	if dis, err = w.DirectoriesInspect(); err != nil {
		return
	}
	if fis, err = w.FramesInspect(); err != nil {
		return
	}
	wi = &WorkspaceInspector{
		Name:         w.Name,
		About:        w.About,
		Frames:       fis,
		Repositories: ris,
		Directories:  dis,
	}
	return
}

func (w *Workspace) FramesInspect() (fis []*FrameInspector, err error) {
	var fi *FrameInspector
	fis = []*FrameInspector{}
	for _, f := range w.Frames {
		if fi, err = f.Inspect(); err != nil {
			return
		}
		fis = append(fis, fi)
	}
	return
}

func (w *Workspace) RepositoriesInspect() (ris []*RepositoryInspector, err error) {
	ris = []*RepositoryInspector{}
	for _, r := range w.Repositories {
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
	for _, d := range w.Directories {
		var di *DirectoryInspector
		if di, err = d.Inspect(); err != nil {
			return
		}
		dis = append(dis, di)
	}
	return
}

// Data

func (w *Workspace) directory() (dirPath string) {
	dirPath = utils.DataDir(filepath.Join("workspaces", w.Name))
	return
}

func (w *Workspace) IsExists() (is bool) {
	if w.UserContext.FindWorkspace(w.Name) != nil {
		is = true
		return
	}
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

func (w *Workspace) Validation() (vn *validations.Validation) {
	vn = validations.NewValidation()
	if w.Name == "" {
		vn.Add("Name", "must not be blank")
	}
	if !utils.IsValidName(w.Name) {
		vn.Add("Name", "must contain word characters, digits, hyphens and underscores only")
	}
	return
}

func (w *Workspace) Update(updates map[string]any) (vn *validations.Validation) {
	w.Assign(updates)
	if vn = w.Validation(); vn.IsValid() {
		w.Save()
	}
	return
}

// Record

func (w *Workspace) Save() {
	queries.Save(w)
}

func (w *Workspace) Delete() {
	utils.RemoveDir(w.directory())
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

func (w *Workspace) FindRepository(uri string) (r *Repository) {
	for _, r = range w.Repositories {
		if r.URI == uri {
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

func (w *Workspace) FrameNames() (ns []string) {
	ns = []string{}
	for _, f := range w.Frames {
		ns = append(ns, f.Name)
	}
	return
}

func (w *Workspace) DirectoryPaths() (ps []string) {
	ps = []string{}
	for _, d := range w.Directories {
		ps = append(ps, d.Path)
	}
	return
}

func (w *Workspace) RepositoryUris() (us []string) {
	us = []string{}
	for _, r := range w.Repositories {
		us = append(us, r.URI)
	}
	return
}
