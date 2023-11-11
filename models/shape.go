package models

import (
	"sf/app"
	"sf/database/queries"
	"sf/utils"
	"time"

	"gorm.io/plugin/soft_delete"
)

type Shape struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     soft_delete.DeletedAt `gorm:"softDelete:nano;index:idx_nondeleted_frame_shape,unique"`
	FrameID       uint                  `gorm:"index:idx_nondeleted_frame_shape,unique"`
	Frame         *Frame                `gorm:"foreignkey:FrameID"`
	Name          string                `gorm:"index:idx_nondeleted_frame_shape,unique"`
	About         string
	ShaperName    string
	Configuration *Configuration `gorm:"polymorphic:Owner;"`
	Shaper        *Shaper        `gorm:"-"`
}

type ShapeReader struct {
	Name          string
	About         string
	Workspace     string
	Frame         string
	Shaper        string
	Configuration []map[string]any
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

func CreateShape(f *Frame, shaper string, name string, about string) (s *Shape, vn *app.Validation, err error) {
	s = NewShape(f, name)
	s.ShaperName = shaper
	if err = s.Load("Shaper"); err != nil {
		return
	}
	if name == "" {
		s.Name = s.Shaper.Name
	}
	if about == "" {
		s.About = s.Shaper.About
	}
	if s.IsExists() {
		err = app.Error("shape %s already exists in frame %s workspace %s", name, f.Name, f.Workspace.Name)
		return
	}
	if vn = s.Validation(); vn.IsValid() {
		s.Save()
	}
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
		Configuration: s.Configuration.Details(),
	}
	return
}

// Read

func (s *Shape) Read() (si *ShapeReader) {
	si = &ShapeReader{
		Name:          s.Name,
		About:         s.About,
		Workspace:     s.Frame.Workspace.Name,
		Frame:         s.Frame.Name,
		Shaper:        s.ShaperName,
		Configuration: s.Configuration.Details(),
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
}

func (s *Shape) Validation() (vn *app.Validation) {
	vn = &app.Validation{}
	if s.ShaperName == "" {
		vn.Add("Shaper", "must not be blank")
	}
	if s.Name == "" {
		vn.Add("Name", "must not be blank")
	}
	if !utils.IsValidName(s.Name) {
		vn.Add("Name", "must contain word characters, digits, hyphens and underscores only")
	}
	return
}

func (s *Shape) Update(update map[string]any) (vn *app.Validation) {
	s.Assign(update)
	if vn = s.Validation(); vn.IsValid() {
		s.Save()
	}
	return
}

// Record

func (s *Shape) Save() {
	queries.Save(s)
}

func (s *Shape) Destroy() {
	queries.Delete(s)
}

// Configuration

func (s *Shape) configurationFormSchema() (schema *FormSchema) {
	schema = s.Shaper.configurationFormSchema()
	return
}

func (s *Shape) readID() (id uint) {
	id = s.ID
	return
}

func (s *Shape) readType() (t string) {
	t = "Shape"
	return
}
