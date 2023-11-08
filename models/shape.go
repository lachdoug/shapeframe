package models

import (
	"sf/app"
	"sf/database/queries"
	"sf/utils"
	"time"

	"gorm.io/datatypes"
	"gorm.io/plugin/soft_delete"
)

type Shape struct {
	ID                uint `gorm:"primaryKey"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         soft_delete.DeletedAt `gorm:"softDelete:nano;index:idx_nondeleted_frame_shape,unique"`
	Frame             *Frame                `gorm:"foreignkey:FrameID"`
	FrameID           uint                  `gorm:"index:idx_nondeleted_frame_shape,unique"`
	Name              string                `gorm:"index:idx_nondeleted_frame_shape,unique"`
	About             string
	ConfigurationJson datatypes.JSON
	ShaperName        string
	Shaper            *Shaper `gorm:"-"`
	Configuration     *Form   `gorm:"-"`
}

type ShapeInspector struct {
	Name          string
	About         string
	Shaper        string
	Configuration []map[string]any
}

// Construction

func NewShape(f *Frame, name string) (s *Shape) {
	s = &Shape{
		Frame: f,
		Name:  name,
	}
	return
}

func CreateShape(f *Frame, shaper string, name string, about string) (s *Shape, err error) {
	s = NewShape(f, name)
	s.ShaperName = shaper
	s.About = about
	if err = s.Load("Shaper"); err != nil {
		return
	}
	s.Label()
	if s.IsExists() {
		err = app.Error("shape %s already exists in frame %s workspace %s", name, f.Name, f.Workspace.Name)
		return
	}
	s.Create()
	return
}

func ResolveShape(uc *UserContext, f *Frame, name string, loads ...string) (s *Shape, err error) {
	if name == "" {
		if uc == nil {
			err = app.Error("no user context")
			return
		}
		s = uc.Shape
		if s == nil {
			err = app.Error("no shape context")
			return
		}
	} else {
		if f == nil {
			err = app.Error("no frame")
			return
		}
		if len(f.Shapes) == 0 {
			err = app.Error("no shapes exist in frame %s", f.Name)
			return
		}
		s = f.FindShape(name)
		if s == nil {
			err = app.Error("shape %s does not exist in frame %s", name, f.Name)
			return
		}
	}
	if len(loads) > 0 {
		if err = s.Load(loads...); err != nil {
			return
		}
	}
	return
}

// Inspection

func (s *Shape) Inspect() (si *ShapeInspector) {
	si = &ShapeInspector{
		Name:          s.Name,
		About:         s.About,
		Shaper:        s.ShaperName,
		Configuration: s.Configuration.SettingsDetail(),
	}
	return
}

// Data

func (s *Shape) IsExists() (is bool) {
	if s.Frame.FindShape(s.Name) != nil {
		is = true
		return
	}
	return
}

func (s *Shape) Load(loads ...string) (err error) {
	sl := NewShapeLoader(s, loads)
	err = sl.load()
	return
}

func (s *Shape) Assign(params map[string]any) {
	if params["Name"] != nil {
		s.Name = params["Name"].(string)
	}
	if params["About"] != nil {
		s.About = params["About"].(string)
	}
	if params["Configuration"] != nil {
		s.ConfigurationJson = datatypes.JSON(utils.JsonMarshal(params["Configuration"]))
	}
}

func (s *Shape) Label() {
	if s.Name == "" {
		s.Name = s.Shaper.Name
	}
	if s.About == "" {
		s.About = s.Shaper.About
	}
}

func (s *Shape) Create() {
	queries.Create(s)
}

func (s *Shape) Save() {
	queries.Update(s)
}

func (s *Shape) Destroy() {
	queries.Delete(s)
}

// Configuration

func (s *Shape) ConfigurationSettings() (settings map[string]any) {
	j := s.ConfigurationJson.String()
	if j != "" {
		settings = map[string]any{}
		utils.JsonUnmarshal([]byte(string(s.ConfigurationJson)), &settings)
	}
	return
}
