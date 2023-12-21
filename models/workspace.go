package models

import (
	"path/filepath"
	"sf/app/dirs"
	"sf/app/errors"
	"sf/app/validations"
	"sf/database/queries"
	"sf/utils"
	"strings"

	"gorm.io/gorm"
)

type Workspace struct {
	gorm.Model
	Name         string `gorm:"index:idx_nondeleted_usercontext_workspace,unique"`
	About        string
	FrameID      uint
	ShapeID      uint
	Frame        *Frame `gorm:"foreignkey:FrameID"`
	Shape        *Shape `gorm:"foreignkey:ShapeID"`
	Frames       []*Frame
	Directories  []*Directory
	Repositories []*Repository `gorm:"-"`
	Framers      []*Framer     `gorm:"-"`
	Shapers      []*Shaper     `gorm:"-"`
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

func NewWorkspace() (w *Workspace) {
	w = &Workspace{}
	return
}

func CreateWorkspace(path string, name string, about string) (w *Workspace, vn *validations.Validation, err error) {
	if path, err = filepath.Abs(path); err != nil {
		err = errors.ErrorWrap(err, "resolve directory")
	}
	if name == "" {
		name = filepath.Base(path)
	}

	w = NewWorkspace()
	w.Name = name
	w.About = about

	if vn = w.Validation(); vn.IsValid() {
		w.Save()
	}
	return
}

func ResolveWorkspace(loads ...string) (w *Workspace, err error) {
	w = NewWorkspace()
	if err = w.Load(loads...); err != nil {
		return
	}
	if w.ID == 0 {
		err = errors.Error("workspace is not initalized")
		return
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

func (w *Workspace) Directory() (dirPath string) {
	dirPath = dirs.WorkspaceDir(".")
	return
}

func (w *Workspace) Load(loads ...string) (err error) {
	wl := NewWorkspaceLoader(w, loads)
	err = wl.load()
	return
}

func (w *Workspace) Clear(assocName string) {
	queries.Clear(w, assocName)
}

// Record

func (w *Workspace) Save() {
	queries.Save(w)
}

func (w *Workspace) Assign(params map[string]string) {
	if params["Name"] != "" {
		w.Name = params["Name"]
	}
	if params["About"] != "" {
		w.About = params["About"]
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

func (w *Workspace) Update(updates map[string]string) (vn *validations.Validation) {
	w.Assign(updates)
	if vn = w.Validation(); vn.IsValid() {
		w.Save()
	}
	return
}

// Associations

func (w *Workspace) GitRepoDirectoryFor(uri string) (dirPath string) {
	elem := []string{}
	elem = append(elem, w.Directory(), "repos")
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

func (w *Workspace) FrameName() (n string) {
	if w.Frame != nil {
		n = w.Frame.Name
	}
	return
}

func (w *Workspace) ShapeName() (n string) {
	if w.Shape != nil {
		n = w.Shape.Name
	}
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
