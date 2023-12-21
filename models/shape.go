package models

import (
	"fmt"
	"sf/app/errors"
	"sf/app/validations"
	"sf/database/queries"
	"sf/utils"
	"strings"
	"time"

	"gorm.io/plugin/soft_delete"
)

type Shape struct {
	ID                 uint `gorm:"primaryKey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          soft_delete.DeletedAt `gorm:"softDelete:nano;index:idx_nondeleted_frame_shape,unique"`
	FrameID            uint                  `gorm:"index:idx_nondeleted_frame_shape,unique"`
	Frame              *Frame                `gorm:"foreignkey:FrameID"`
	Name               string                `gorm:"index:idx_nondeleted_frame_shape,unique"`
	About              string
	ShaperName         string
	ShapeConfiguration *ShapeConfiguration `gorm:"-"`
	Shaper             *Shaper             `gorm:"-"`
}

type ShapeReader struct {
	Name          string
	About         string
	Workspace     string
	Frame         string
	Shaper        string
	Configuration *ShapeConfiguration
}

type ShapeInspector struct {
	Name          string
	About         string
	Shaper        string
	Configuration *ShapeConfigurationInspector
}

type ShapeOutput struct {
	Identifier string
	Workspace  string
	Frame      string
	Name       string
	About      string
	Settings   *ShapeOutputSettings
	Build      *ShapeOutputBuild
	Start      []string
	Ports      [][]string
	Volumes    [][]string
}

type ShapeOutputBuild struct {
	On string
	Do [][]string
}

type ShapeOutputSettings struct {
	Shape map[string]string
	Frame map[string]string
}

// Construction

func NewShape(f *Frame, name string) (s *Shape) {
	s = &Shape{
		Frame: f,
		Name:  name,
	}
	return
}

func CreateShape(f *Frame, shaper string, name string, about string) (s *Shape, vn *validations.Validation, err error) {
	s = NewShape(f, name)
	s.ShaperName = shaper
	s.About = about
	if vn = s.ShaperValidation(); vn.IsInvalid() {
		return
	}
	if err = s.Load("Shaper"); err != nil {
		return
	}
	if s.Name == "" {
		s.Name = s.Shaper.Name
	}
	if s.About == "" {
		s.About = s.Shaper.About
	}
	if s.IsExists() {
		err = errors.Errorf("shape %s already exists in frame %s", s.Name, f.Name)
		return
	}
	if vn = s.Validation(); vn.IsValid() {
		s.Save()
	}
	return
}

func ResolveShape(w *Workspace, f *Frame, name string, loads ...string) (s *Shape, err error) {
	if name == "" {
		if w == nil {
			err = errors.Error("no workspaces")
			return
		}
		s = w.Shape
		if s == nil {
			err = errors.Error("no shape context")
			return
		}
	} else {
		if f == nil {
			err = errors.Error("no frame")
			return
		}
		if len(f.Shapes) == 0 {
			err = errors.Errorf("no shapes exist in frame %s", f.Name)
			return
		}
		s = f.FindShape(name)
		if s == nil {
			err = errors.Errorf("shape %s does not exist in frame %s", name, f.Name)
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

// Identification

func (s *Shape) identifier() (iden string) {
	iden = strings.ToLower(fmt.Sprintf(
		"%s.%s.%s",
		s.Name,
		s.Frame.Name,
		s.Frame.Workspace.Name,
	))
	return
}

// Inspection

func (s *Shape) Inspect() (si *ShapeInspector) {
	si = &ShapeInspector{
		Name:          s.Name,
		About:         s.About,
		Shaper:        s.ShaperName,
		Configuration: s.ShapeConfiguration.Inspect(),
	}
	return
}

// Read

func (s *Shape) Read() (sr *ShapeReader) {
	sr = &ShapeReader{
		Name:          s.Name,
		About:         s.About,
		Workspace:     s.Frame.Workspace.Name,
		Frame:         s.Frame.Name,
		Shaper:        s.ShaperName,
		Configuration: s.ShapeConfiguration,
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

func (s *Shape) Assign(params map[string]string) {
	if params["Name"] != "" {
		s.Name = params["Name"]
	}
	if params["About"] != "" {
		s.About = params["About"]
	}
}

func (s *Shape) ShaperValidation() (vn *validations.Validation) {
	vn = validations.NewValidation()
	if s.ShaperName == "" {
		vn.Add("Shaper", "must not be blank")
	}
	return
}

func (s *Shape) Validation() (vn *validations.Validation) {
	vn = validations.NewValidation()
	if s.Name == "" {
		vn.Add("Name", "must not be blank")
	}
	if !utils.IsValidName(s.Name) {
		vn.Add("Name", "must contain word characters, digits, hyphens and underscores only")
	}
	return
}

func (s *Shape) Update(updates map[string]string) (vn *validations.Validation) {
	s.Assign(updates)
	if vn = s.Validation(); vn.IsValid() {
		s.Save()
	}
	return
}

// Record

func (s *Shape) Save() {
	queries.Save(s)
}

func (s *Shape) Delete() {
	queries.Delete(s)
}

// Output

func (s *Shape) Output() (o *ShapeOutput) {
	o = &ShapeOutput{
		Identifier: s.identifier(),
		Workspace:  s.Frame.Workspace.Name,
		Frame:      s.Frame.Name,
		Name:       s.Name,
		About:      s.About,
		Settings: &ShapeOutputSettings{
			Shape: s.ShapeConfiguration.Shape.Settings,
			Frame: s.ShapeConfiguration.Frame.Settings,
		},
		Start:   s.Shaper.Start,
		Ports:   s.Shaper.Ports,
		Volumes: s.Shaper.Volumes,
	}
	if s.Shaper.Build != nil {
		o.Build = &ShapeOutputBuild{
			On: s.Shaper.Build.On,
			Do: s.Shaper.Build.Do,
		}
	} else {
		o.Build = &ShapeOutputBuild{
			On: "",
			Do: [][]string{},
		}
	}
	return
}
