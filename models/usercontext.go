package models

import (
	"sf/database/queries"

	"gorm.io/gorm"
)

type UserContext struct {
	gorm.Model
	WorkspaceID uint
	FrameID     uint
	ShapeID     uint
	Workspace   *Workspace `gorm:"foreignkey:WorkspaceID"`
	Frame       *Frame     `gorm:"foreignkey:FrameID"`
	Shape       *Shape     `gorm:"foreignkey:ShapeID"`
	Workspaces  []*Workspace
}

// Construction

func NewUserContext(loads ...string) (uc *UserContext) {
	uc = &UserContext{
		Model: gorm.Model{ID: uint(1)},
	}
	return
}

func ResolveUserContext(loads ...string) (uc *UserContext) {
	uc = NewUserContext()
	uc.Load(loads...)
	return
}

// Data

func (uc *UserContext) Load(preloads ...string) {
	queries.Load(uc, uc.ID, preloads...)
}

func (uc *UserContext) Clear(assocName string) {
	queries.Clear(uc, assocName)
}

// Record

func (uc *UserContext) Save() {
	queries.Save(uc)
}

// Associations

func (uc *UserContext) WorkspaceName() (n string) {
	if uc.Workspace != nil {
		n = uc.Workspace.Name
	}
	return
}

func (uc *UserContext) FrameName() (n string) {
	if uc.Frame != nil {
		n = uc.Frame.Name
	}
	return
}

func (uc *UserContext) ShapeName() (n string) {
	if uc.Shape != nil {
		n = uc.Shape.Name
	}
	return
}

func (uc *UserContext) FindWorkspace(name string) (w *Workspace) {
	for _, w = range uc.Workspaces {
		if w.Name == name {
			return
		}
	}
	w = nil
	return
}
