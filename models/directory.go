package models

import (
	"sf/app"
	"sf/database/queries"
	"time"

	"gorm.io/plugin/soft_delete"
)

type Directory struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   soft_delete.DeletedAt `gorm:"softDelete:nano;index:idx_nondeleted_workspace_directory,unique"`
	Workspace   *Workspace            `gorm:"foreignkey:WorkspaceID"`
	WorkspaceID uint                  `gorm:"index:idx_nondeleted_workspace_directory,unique"`
	Path        string                `gorm:"index:idx_nondeleted_workspace_directory,unique"`
	GitRepo     *GitRepo              `gorm:"-"`
	Framers     []*Framer             `gorm:"-"`
	Shapers     []*Shaper             `gorm:"-"`
}

type DirectoryInspector struct {
	Path    string
	GitRepo *GitRepoInspector
}

// Construction

func NewDirectory(w *Workspace, path string) (d *Directory) {
	d = &Directory{
		Workspace: w,
		Path:      path,
	}
	return
}

func CreateDirectory(w *Workspace, path string) (d *Directory, err error) {
	d = NewDirectory(w, path)
	if d.IsExists() {
		err = app.Error("directory %s already exists in workspace %s", path, w.Name)
		return
	}
	d.Create()
	return
}

func ResolveDirectory(w *Workspace, path string, loads ...string) (d *Directory, err error) {
	if w == nil {
		err = app.Error("no workspace")
		return
	}
	if len(w.Directories) == 0 {
		err = app.Error("no directories exist in workspace %s", w.Name)
		return
	}
	d = w.FindDirectory(path)
	if d == nil {
		err = app.Error("directory %s does not exist in workspace %s", path, w.Name)
		return
	}
	if len(loads) > 0 {
		if err = d.Load(loads...); err != nil {
			return
		}
	}
	return
}

// Inspection

func (d *Directory) Inspect() (di *DirectoryInspector) {
	di = &DirectoryInspector{
		Path:    d.Path,
		GitRepo: d.GitRepoInspect(),
	}
	return
}

func (d *Directory) GitRepoInspect() (gri *GitRepoInspector) {
	gri = d.GitRepo.Inspect()
	return
}

// Data

func (d *Directory) IsExists() (is bool) {
	if d.Workspace.FindDirectory(d.Path) != nil {
		is = true
		return
	}
	return
}

func (d *Directory) Load(loads ...string) (err error) {
	dl := NewDirectoryLoader(d, loads)
	err = dl.load()
	return
}

func (d *Directory) Create() {
	queries.Create(d)
}

func (d *Directory) Destroy() {
	queries.Delete(d)
}
