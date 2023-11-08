package models

import (
	"sf/app"
	"sf/database/queries"
	"sf/utils"
	"time"

	"gorm.io/datatypes"
	"gorm.io/plugin/soft_delete"
)

type Frame struct {
	ID                uint `gorm:"primaryKey"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         soft_delete.DeletedAt `gorm:"softDelete:nano;index:idx_nondeleted_workspace_frame,unique"`
	Workspace         *Workspace            `gorm:"foreignkey:WorkspaceID"`
	WorkspaceID       uint                  `gorm:"index:idx_nondeleted_workspace_frame,unique"`
	Name              string                `gorm:"index:idx_nondeleted_workspace_frame,unique"`
	About             string
	Shapes            []*Shape
	ConfigurationJson datatypes.JSON
	FramerName        string
	Framer            *Framer `gorm:"-"`
	Configuration     *Form   `gorm:"-"`
}

type FrameInspector struct {
	Name          string
	About         string
	Shapes        []*ShapeInspector
	Framer        string
	Configuration []map[string]any
}

// Construction

func NewFrame(w *Workspace, name string) (f *Frame) {
	f = &Frame{
		Workspace: w,
		Name:      name,
	}
	return
}

func CreateFrame(w *Workspace, framer string, name string, about string) (f *Frame, err error) {
	f = NewFrame(w, name)
	f.FramerName = framer
	f.About = about
	if err = f.Load("Framer"); err != nil {
		return
	}
	f.Label()
	if f.IsExists() {
		err = app.Error("frame %s already exists in workspace %s", name, w.Name)
		return
	}
	f.Create()
	return
}

func ResolveFrame(uc *UserContext, w *Workspace, name string, loads ...string) (f *Frame, err error) {
	if name == "" {
		if uc == nil {
			err = app.Error("no user context")
			return
		}
		f = uc.Frame
		if f == nil {
			err = app.Error("no frame context")
			return
		}
	} else {
		if w == nil {
			err = app.Error("no workspace")
			return
		}
		if len(w.Frames) == 0 {
			err = app.Error("no frames exist in workspace %s", w.Name)
			return
		}
		f = w.FindFrame(name)
		if f == nil {
			err = app.Error("frame %s does not exist in workspace %s", name, w.Name)
			return
		}
	}
	if len(loads) > 0 {
		if err = f.Load(loads...); err != nil {
			return
		}
	}
	return
}

// Inspection

func (f *Frame) Inspect() (fi *FrameInspector) {
	fi = &FrameInspector{
		Name:          f.Name,
		About:         f.About,
		Shapes:        f.ShapesInspect(),
		Framer:        f.FramerName,
		Configuration: f.Configuration.SettingsDetail(),
	}
	return
}

func (f *Frame) ShapesInspect() (sis []*ShapeInspector) {
	sis = []*ShapeInspector{}
	for _, s := range f.Shapes {
		sis = append(sis, s.Inspect())
	}
	return
}

// Data

func (f *Frame) IsExists() (is bool) {
	if f.Workspace.FindFrame(f.Name) != nil {
		is = true
		return
	}
	return
}

func (f *Frame) Load(loads ...string) (err error) {
	fl := NewFrameLoader(f, loads)
	err = fl.load()
	return
}

func (f *Frame) Assign(params map[string]any) {
	if params["Name"] != nil {
		f.Name = params["Name"].(string)
	}
	if params["About"] != nil {
		f.About = params["About"].(string)
	}
	if params["Configuration"] != nil {
		f.ConfigurationJson = datatypes.JSON(utils.JsonMarshal(params["Configuration"]))
	}
}

func (f *Frame) Label() {
	if f.Name == "" {
		f.Name = f.Framer.Name
	}
	if f.About == "" {
		f.About = f.Framer.About
	}
}

func (f *Frame) Create() {
	queries.Create(f)
}

func (f *Frame) Save() {
	queries.Update(f)
}

func (f *Frame) Destroy() {
	queries.Delete(f)
}

// Orchestration

func (f *Frame) Orchestrate(st *Stream) {
	o := NewOrchestration(f, st)
	go o.apply()
}

// Configuration

func (f *Frame) ConfigurationSettings() (settings map[string]any) {
	j := f.ConfigurationJson.String()
	if j != "" {
		settings = map[string]any{}
		utils.JsonUnmarshal([]byte(string(f.ConfigurationJson)), &settings)
	}
	return
}

// Associations

func (f *Frame) FindShape(name string) (s *Shape) {
	for _, s = range f.Shapes {
		if s.Name == name {
			return
		}
	}
	s = nil
	return
}
