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
	ID         uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  soft_delete.DeletedAt `gorm:"index:idx_nondeleted_frame_shape,unique"`
	Frame      *Frame                `gorm:"foreignkey:FrameID"`
	FrameID    uint                  `gorm:"index:idx_nondeleted_frame_shape,unique"`
	Name       string                `gorm:"index:idx_nondeleted_frame_shape,unique"`
	About      string
	ShaperName string
	ConfigJson datatypes.JSON
	Shaper     *Shaper `gorm:"-"`
}

type ShapeInspector struct {
	Name   string
	About  string
	Shaper string
}

// Construction

func ShapeNew(f *Frame, name string) (s *Shape) {
	if f == nil {
		panic("Shape Frame is <nil>")
	}
	s = &Shape{
		Frame: f,
		Name:  utils.StringTidy(name),
	}
	return
}

// Inspection

func (s *Shape) Inspect() (si *ShapeInspector) {
	si = &ShapeInspector{
		Name:   s.Name,
		About:  s.About,
		Shaper: s.ShaperName,
	}
	return
}

// Data

func (s *Shape) IsExists() (is bool) {
	if s.Frame.ShapeFind(s.Name) != nil {
		is = true
		return
	}
	return
}

func (s *Shape) Load(preloads ...string) (err error) {
	queries.Load(s, s.ID, preloads...)
	return
}

func (s *Shape) Assign(params map[string]any) {
	if params["ShaperName"] != nil {
		s.ShaperName = utils.StringTidy(params["ShaperName"].(string))
	}
	if params["About"] != nil {
		s.About = utils.StringTidy(params["About"].(string))
	}
	if params["Config"] != nil {
		s.ConfigJson = datatypes.JSON(utils.JsonMarshal(params["Config"]))
	}
}

func (s *Shape) Create() (err error) {
	if s.Name == "" {
		s.Name = s.Shaper.Name
	}
	if s.About == "" {
		s.About = s.Shaper.About
	}
	if s.IsExists() {
		err = app.Error(nil, "shape %s already exists in frame %s", s.Name, s.Frame.Name)
		return
	}
	queries.Create(s)
	return
}

func (s *Shape) Save() (err error) {
	queries.Update(s)
	return
}

func (s *Shape) SaveConfiguration() (err error) {
	if err = s.ConfigurationValidate(); err != nil {
		return
	}
	queries.Update(s)
	return
}

func (s *Shape) Destroy() {
	queries.Delete(s)
}

// Configuration

func (s *Shape) ConfigValues() (config map[string]any) {
	j := s.ConfigJson.String()
	if j != "" {
		utils.JsonUnmarshal([]byte(string(s.ConfigJson)), &config)
	}
	return
}

func (s *Shape) ConfigurationValidate() (err error) {
	err = s.Shaper.ConfigurationValidate(s.Name, s.ConfigValues())
	return
}

// Associations

func (s *Shape) SetShaper() (err error) {
	var sr *Shaper
	if sr, err = s.ShaperFind(); err != nil {
		return
	} else if sr == nil {
		err = app.Error(nil, "locate shaper %s: no such shaper in workspace %s", s.ShaperName, s.Frame.Workspace.Name)
		return
	} else {
		sr.Load()
		s.Shaper = sr
	}
	return
}

func (s *Shape) ShaperFind() (sr *Shaper, err error) {
	f := s.Frame
	w := f.Workspace
	sr, err = w.ShaperFind(s.ShaperName)
	return
}
