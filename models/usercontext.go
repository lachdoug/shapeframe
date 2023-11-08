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
	// Frames      []*Frame  `gorm:"-"`
	// Shapes      []*Shape  `gorm:"-"`
	// Framers     []*Framer `gorm:"-"`
	// Shapers     []*Shaper `gorm:"-"`
}

type UserContextChange struct {
	Exit  *UserContextInspector
	Enter *UserContextInspector
}

type UserContextInspector struct {
	Workspace string
	Frame     string
	Shape     string
}

// Construction

func NewUserContext(loads ...string) (uc *UserContext) {
	uc = &UserContext{
		Model: gorm.Model{ID: uint(1)},
	}
	return
}

// Resolve user context
func ResolveUserContext(loads ...string) (uc *UserContext) {
	uc = NewUserContext()
	uc.Load(loads...)
	return
}

// Data

func (uc *UserContext) Load(preloads ...string) {
	queries.Load(uc, uc.ID, preloads...)
}

func (uc *UserContext) Save() {
	queries.Update(uc)
}

func (uc *UserContext) Clear(assocName string) {
	queries.Clear(uc, assocName)
}

// Inspection

func (uc *UserContext) Inspect() (uci *UserContextInspector) {
	uci = &UserContextInspector{
		Workspace: uc.WorkspaceName(),
		Frame:     uc.FrameName(),
		Shape:     uc.ShapeName(),
	}
	return
}

// Associations

func (uc *UserContext) WorkspaceName() (n string) {
	if uc.Workspace == nil {
		n = ""
	} else {
		n = uc.Workspace.Name
	}
	return
}

func (uc *UserContext) FrameName() (n string) {
	if uc.Frame == nil {
		n = ""
	} else {
		n = uc.Frame.Name
	}
	return
}

func (uc *UserContext) ShapeName() (n string) {
	if uc.Shape == nil {
		n = ""
	} else {
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
