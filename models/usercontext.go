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

func UserContextNew() (uc *UserContext) {
	uc = &UserContext{}
	uc.Singleton()
	return
}

// Data

func (uc *UserContext) Singleton() {
	uc.Load()
	if uc.ID == uint(0) {
		queries.Create(uc)
	}
}

func (uc *UserContext) Load(preloads ...string) {
	queries.Load(uc, uint(1), preloads...)
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

func (uc *UserContext) Lookup(assocName string, key string, value string, model any) {
	queries.Lookup(uc, assocName, key, value, model)
}

func (uc *UserContext) WorkspacesLoad() (ws []*Workspace) {
	ws = []*Workspace{}
	for _, w := range uc.Workspaces {
		w.Load()
		ws = append(ws, w)
	}
	return
}

func (uc *UserContext) Frames() (fs []*Frame) {
	fs = []*Frame{}
	for _, w := range uc.Workspaces {
		fs = append(fs, w.Frames...)
	}
	return
}

func (uc *UserContext) Shapes() (ss []*Shape) {
	ss = []*Shape{}
	for _, w := range uc.Workspaces {
		ss = append(ss, w.Shapes()...)
	}
	return
}

func (uc *UserContext) Framers() (frs []*Framer, err error) {
	frs = []*Framer{}
	for _, w := range uc.WorkspacesLoad() {
		var wfrs []*Framer
		if wfrs, err = w.Framers(); err != nil {
			return
		}
		frs = append(frs, wfrs...)
	}
	return
}

func (uc *UserContext) Shapers() (frs []*Shaper, err error) {
	frs = []*Shaper{}
	for _, w := range uc.WorkspacesLoad() {
		var wsrs []*Shaper
		if wsrs, err = w.Shapers(); err != nil {
			return
		}
		frs = append(frs, wsrs...)
	}
	return
}

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

func (uc *UserContext) WorkspaceFind(name string) (w *Workspace) {
	w = &Workspace{}
	uc.Lookup("Workspaces", "name", name, w)
	if w.ID == uint(0) {
		w = nil
	}
	return
}
