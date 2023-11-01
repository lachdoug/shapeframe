package models

import (
	"fmt"
	"sf/database/queries"
	"time"

	"gorm.io/plugin/soft_delete"
)

type Directory struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   soft_delete.DeletedAt `gorm:"index:idx_nondeleted_workspace_directory,unique"`
	Workspace   *Workspace            `gorm:"foreignkey:WorkspaceID"`
	WorkspaceID uint                  `gorm:"index:idx_nondeleted_workspace_directory,unique"`
	Path        string                `gorm:"index:idx_nondeleted_workspace_directory,unique"`
	GitRepo     *GitRepo              `gorm:"-"`
}

type DirectoryInspector struct {
	Path    string
	GitRepo *GitRepoInspector
}

// Construction

func DirectoryNew(w *Workspace, path string) (d *Directory) {
	if w == nil {
		panic("Directory Workspace is <nil>")
	}
	d = &Directory{
		Workspace: w,
		Path:      path,
		GitRepo:   GitRepoNew(w, path),
	}
	return
}

// Inspection

func (d *Directory) Inspect() (di *DirectoryInspector, err error) {
	var gri *GitRepoInspector
	if gri, err = d.GitRepoInspect(); err != nil {
		return
	}
	di = &DirectoryInspector{
		Path:    d.Path,
		GitRepo: gri,
	}
	return
}

func (d *Directory) GitRepoInspect() (gri *GitRepoInspector, err error) {
	if d.GitRepo.isExists() {
		gri, err = d.GitRepo.Inspect()
	}
	return
}

// Data

func (d *Directory) IsExists() (is bool) {
	if d.Workspace.DirectoryFind(d.Path) != nil {
		is = true
		return
	}
	return
}

func (d *Directory) Load(preloads ...string) {
	queries.Load(d, d.ID, preloads...)
	fmt.Println("GitRepo", d.GitRepo)
	d.GitRepo.load()
}

// func (d *Directory) GitRepoLoad() {
// 	gr := GitRepoNew(d.Workspace, d.Path)
// 	gr.load()
// 	d.GitRepo = gr
// }

func (d *Directory) Create() (err error) {
	queries.Create(d)
	return
}

func (d *Directory) Destroy() {
	queries.Delete(d)
}

// Associations

func (d *Directory) Framers() (frs []*Framer, err error) {
	if d.GitRepo != nil {
		frs, err = d.GitRepo.Framers()
	}
	return
}

func (d *Directory) Shapers() (frs []*Shaper, err error) {
	if d.GitRepo != nil {
		frs, err = d.GitRepo.Shapers()
	}
	return
}
