package models

import (
	"sf/app/errors"
	"sf/app/validations"
	"sf/database/queries"
	"sf/utils"
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

func CreateDirectory(w *Workspace, path string) (d *Directory, vn *validations.Validation, err error) {
	d = NewDirectory(w, path)
	if vn = d.Validation(); vn.IsValid() {
		if d.IsExists() {
			err = errors.Errorf("directory %s already exists in workspace %s", path, w.Name)
			return
		}
		d.Save()
	}
	return
}

func ResolveDirectory(w *Workspace, path string, loads ...string) (d *Directory, err error) {
	if w == nil {
		err = errors.Error("no workspace")
		return
	}
	if len(w.Directories) == 0 {
		err = errors.Errorf("no directories exist in workspace %s", w.Name)
		return
	}
	d = w.FindDirectory(path)
	if d == nil {
		err = errors.Errorf("directory %s does not exist in workspace %s", path, w.Name)
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

func (d *Directory) Inspect() (di *DirectoryInspector, err error) {
	var gri *GitRepoInspector
	if gri, err = d.GitRepo.Inspect(); err != nil {
		return
	}
	di = &DirectoryInspector{
		Path:    d.Path,
		GitRepo: gri,
	}
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

func (d *Directory) Validation() (vn *validations.Validation) {
	vn = validations.NewValidation()
	if d.Path == "" {
		vn.Add("Path", "must not be blank")
	}
	if !utils.IsDir(d.Path) {
		vn.Add("Path", "must be a valid directory path")
	}
	return
}

// Record

func (d *Directory) Save() {
	queries.Save(d)
}

func (d *Directory) Delete() {
	queries.Delete(d)
}
