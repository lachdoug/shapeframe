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
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   soft_delete.DeletedAt `gorm:"index:idx_nondeleted_workspace_frame,unique"`
	Workspace   *Workspace            `gorm:"foreignkey:WorkspaceID"`
	WorkspaceID uint                  `gorm:"index:idx_nondeleted_workspace_frame,unique"`
	Name        string                `gorm:"index:idx_nondeleted_workspace_frame,unique"`
	About       string
	Shapes      []*Shape
	ConfigJson  datatypes.JSON
	FramerName  string
	Framer      *Framer `gorm:"-"`
}

type FrameInspector struct {
	Name   string
	About  string
	Shapes []*ShapeInspector
	Framer string
}

// Construction

func FrameNew(w *Workspace, name string) (f *Frame) {
	if w == nil {
		panic("Frame Workspace is <nil>")
	}
	f = &Frame{
		Workspace: w,
		Name:      utils.StringTidy(name),
	}
	return
}

// Inspection

func (f *Frame) Inspect() (fi *FrameInspector) {
	fi = &FrameInspector{
		Name:   f.Name,
		About:  f.About,
		Shapes: f.ShapesInspect(),
		Framer: f.FramerName,
	}
	return
}

func (f *Frame) ShapesInspect() (sis []*ShapeInspector) {
	sis = []*ShapeInspector{}
	for _, s := range f.Shapes {
		s.Load()
		sis = append(sis, s.Inspect())
	}
	return
}

// Data

func (f *Frame) IsExists() (is bool) {
	if f.Workspace.FrameFind(f.Name) != nil {
		is = true
		return
	}
	return
}

func (f *Frame) Load(preloads ...string) (err error) {
	queries.Load(f, f.ID, preloads...)
	return
}

func (f *Frame) Assign(params map[string]any) {
	if params["FramerName"] != nil {
		f.FramerName = utils.StringTidy(params["FramerName"].(string))
	}
	if params["About"] != nil {
		f.About = utils.StringTidy(params["About"].(string))
	}
	if params["Config"] != nil {
		f.ConfigJson = datatypes.JSON(utils.JsonMarshal(params["Config"]))
	}
}

func (f *Frame) Create() (err error) {
	if f.Name == "" {
		f.Name = f.Framer.Name
	}
	if f.About == "" {
		f.About = f.Framer.About
	}
	if f.IsExists() {
		err = app.Error(nil, "frame %s already exists in workspace %s", f.Name, f.Workspace.Name)
		return
	}
	queries.Create(f)
	return
}

func (f *Frame) Save() (err error) {
	queries.Update(f)
	return
}

func (f *Frame) SaveConfiguration() (err error) {
	if err = f.ConfigurationValidate(); err != nil {
		return
	}
	queries.Update(f)
	return
}

func (f *Frame) Destroy() {
	queries.Delete(f)
}

// Configuration

func (f *Frame) ConfigValues() (config map[string]any) {
	j := f.ConfigJson.String()
	if j != "" {
		utils.JsonUnmarshal([]byte(string(f.ConfigJson)), &config)
	}
	return
}

func (f *Frame) ConfigurationValidate() (err error) {
	err = f.Framer.ConfigurationValidate(f.Name, f.ConfigValues())
	return
}

// Associations

func (f *Frame) SetFramer() (err error) {
	var fr *Framer
	if fr, err = f.FramerFind(); err != nil {
		return
	} else if fr == nil {
		err = app.Error(nil, "failed to locate framer %s in workspace %s", f.FramerName, f.Workspace.Name)
		return
	} else {
		fr.Load()
		f.Framer = fr
	}
	return
}

func (f *Frame) FramerFind() (sr *Framer, err error) {
	w := f.Workspace
	sr, err = w.FramerFind(f.FramerName)
	return
}

func (f *Frame) ShapeFind(name string) (s *Shape) {
	s = &Shape{}
	f.Lookup("Shapes", "name", name, s)
	if s.ID == uint(0) {
		s = nil
	}
	return
}

func (f *Frame) Lookup(assocName string, key string, value string, model any) {
	queries.Lookup(f, assocName, key, value, model)
}
