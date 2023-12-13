package models

import (
	"fmt"
	"sf/app/errors"
	"sf/app/streams"
	"sf/app/validations"
	"sf/database/queries"
	"sf/utils"
	"strings"
	"time"

	"gorm.io/plugin/soft_delete"
)

type Frame struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     soft_delete.DeletedAt `gorm:"softDelete:nano;index:idx_nondeleted_workspace_frame,unique"`
	WorkspaceID   uint                  `gorm:"index:idx_nondeleted_workspace_frame,unique"`
	Workspace     *Workspace            `gorm:"foreignkey:WorkspaceID"`
	Name          string                `gorm:"index:idx_nondeleted_workspace_frame,unique"`
	About         string
	Parent        *Frame
	FramerName    string
	ParentID      *uint
	Children      []*Frame `gorm:"foreignkey:ParentID"`
	Shapes        []*Shape
	Configuration *Configuration `gorm:"polymorphic:Owner;"`
	Framer        *Framer        `gorm:"-"`
}

type FrameInspector struct {
	Name          string
	About         string
	Framer        string
	Configuration []map[string]string
	Relations     *FrameRelationsInspector
	Shapes        []*ShapeInspector
}

type FrameRelationsInspector struct {
	Parent      string
	Children    []string
	Ancestors   []string
	Descendents []string
}

type FrameReader struct {
	Name          string
	About         string
	Workspace     string
	Parent        string
	Framer        string
	Configuration []map[string]string
	Shapes        []string
}

type FrameOutput struct {
	Identifier string
	Workspace  string
	Name       string
	About      string
	Config     map[string]string
}

// Construction

func NewFrame(w *Workspace, name string) (f *Frame) {
	f = &Frame{
		Workspace: w,
		Name:      name,
	}
	return
}

func CreateFrame(w *Workspace, framer string, name string, about string) (f *Frame, vn *validations.Validation, err error) {
	f = NewFrame(w, name)
	f.FramerName = framer
	if vn = f.FramerValidation(); vn.IsInvalid() {
		return
	}
	if err = f.Load("Framer"); err != nil {
		return
	}
	if name == "" {
		f.Name = f.Framer.Name
	}
	if about == "" {
		f.About = f.Framer.About
	}
	if f.IsExists() {
		err = errors.Errorf("frame %s already exists in workspace %s", name, w.Name)
		return
	}
	if vn = f.Validation(); vn.IsValid() {
		f.Save()
	}
	return
}

func ResolveFrame(uc *UserContext, w *Workspace, name string, loads ...string) (f *Frame, err error) {
	if name == "" {
		if uc == nil {
			err = errors.Error("no user context")
			return
		}
		f = uc.Frame
		if f == nil {
			err = errors.Error("no frame context")
			return
		}
	} else {
		if w == nil {
			err = errors.Error("no workspace")
			return
		}
		if len(w.Frames) == 0 {
			err = errors.Errorf("no frames exist in workspace %s", w.Name)
			return
		}
		f = w.FindFrame(name)
		if f == nil {
			err = errors.Errorf("frame %s does not exist in workspace %s", name, w.Name)
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

// Identification

func (f *Frame) identifier() (iden string) {
	iden = strings.ToLower(fmt.Sprintf(
		"%s.%s",
		f.Name,
		f.Workspace.Name,
	))
	return
}

// Inspection

func (f *Frame) Inspect() (fi *FrameInspector, err error) {
	var ri *FrameRelationsInspector
	if ri, err = f.RelationsInspect(); err != nil {
		return
	}
	fi = &FrameInspector{
		Name:          f.Name,
		About:         f.About,
		Framer:        f.FramerName,
		Configuration: f.Configuration.Info(),
		Relations:     ri,
		Shapes:        f.ShapesInspect(),
	}
	return
}

func (f *Frame) RelationsInspect() (ri *FrameRelationsInspector, err error) {
	var ans []string
	var dns []string
	if ans, err = f.AncestorNames(); err != nil {
		return
	}
	if dns, err = f.DescendentNames(); err != nil {
		return
	}
	ri = &FrameRelationsInspector{
		Parent:      f.ParentName(),
		Children:    f.ChildNames(),
		Ancestors:   ans,
		Descendents: dns,
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

// Read

func (f *Frame) Read() (fr *FrameReader) {
	fr = &FrameReader{
		Name:          f.Name,
		About:         f.About,
		Workspace:     f.Workspace.Name,
		Parent:        f.ParentName(),
		Framer:        f.FramerName,
		Configuration: f.Configuration.Info(),
		Shapes:        f.ShapeNames(),
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
}

func (f *Frame) FramerValidation() (vn *validations.Validation) {
	vn = validations.NewValidation()
	if f.FramerName == "" {
		vn.Add("Framer", "must not be blank")
	}
	return
}

func (f *Frame) Validation() (vn *validations.Validation) {
	vn = validations.NewValidation()
	if f.Name == "" {
		vn.Add("Name", "must not be blank")
	}
	if !utils.IsValidName(f.Name) {
		vn.Add("Name", "must contain word characters, digits, hyphens and underscores only")
	}
	return
}

func (f *Frame) Update(updates map[string]any) (vn *validations.Validation) {
	f.Assign(updates)
	if vn = f.Validation(); vn.IsValid() {
		f.Save()
	}
	return
}

// Record

func (f *Frame) Save() {
	queries.Save(f)
}

func (f *Frame) Delete() {
	queries.Delete(f)
}

// Composition

func (f *Frame) Apply(st *streams.Stream) {
	o := NewComposition(f, st)
	go o.apply()
}

func (f *Frame) Destroy(st *streams.Stream) {
	o := NewComposition(f, st)
	go o.destroy()
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

func (f *Frame) IsCircular() (is bool, err error) {
	var as []*Frame
	if as, err = f.Ancestors(); err != nil {
		return
	}
	for _, a := range as {
		if f.ID == a.ID {
			is = true
			return
		}
	}
	return
}

func (f *Frame) ParentName() (n string) {
	if f.Parent == nil {
		n = ""
	} else {
		n = f.Parent.Name
	}
	return
}

func (f *Frame) ShapeNames() (ns []string) {
	ns = []string{}
	for _, s := range f.Shapes {
		ns = append(ns, s.Name)
	}
	return
}

func (f *Frame) ChildNames() (ns []string) {
	ns = []string{}
	for _, fc := range f.Children {
		ns = append(ns, fc.Name)
	}
	return
}

func (f *Frame) AncestorNames() (ns []string, err error) {
	var as []*Frame
	ns = []string{}
	if as, err = f.Ancestors(); err != nil {
		return
	}
	for _, a := range as {
		ns = append(ns, a.Name)
	}
	return
}

func (f *Frame) DescendentNames() (ns []string, err error) {
	var ds []*Frame
	ns = []string{}
	if ds, err = f.Descendents(); err != nil {
		return
	}
	for _, d := range ds {
		ns = append(ns, d.Name)
	}
	return
}

func (f *Frame) Ancestors() (as []*Frame, err error) {
	as = []*Frame{}
	var pas []*Frame
	if f.Parent != nil {
		p := f.Parent
		if err = p.Load("Parent"); err != nil {
			return
		}
		if pas, err = p.Ancestors(); err != nil {
			return
		}
		as = append(as, p)
		as = append(as, pas...)
	}
	return
}

func (f *Frame) Descendents() (ds []*Frame, err error) {
	ds = []*Frame{}
	var cds []*Frame
	for _, c := range f.Children {
		if err = c.Load("Children"); err != nil {
			return
		}
		if cds, err = c.Descendents(); err != nil {
			return
		}
		ds = append(ds, c)
		ds = append(ds, cds...)
	}
	return
}

// Output

func (f *Frame) Output() (o *FrameOutput) {
	o = &FrameOutput{
		Identifier: f.identifier(),
		Workspace:  f.Workspace.Name,
		Name:       f.Name,
		About:      f.About,
		Config:     f.Configuration.Settings,
	}
	return
}
