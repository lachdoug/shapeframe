package models

import (
	"fmt"
	"path/filepath"
	"sf/app/errors"
	"sf/app/validations"
	"sf/utils"
)

type Shaper struct {
	Workspace *Workspace
	Path      string
	Name      string
	About     string
	Shape     []*FormComponent
	Connect   []*ShaperConnect
	// ConfigurationFormSchema *FormSchema
	Build   *ShaperBuild
	Start   []string
	Ports   [][]string
	Volumes [][]string
}

type ShaperBuild struct {
	On string
	Do [][]string
}

type ShaperInspector struct {
	Name  string
	About string
	URI   string
}

// Construction

func NewShaper(w *Workspace, path string, name string) (sr *Shaper) {
	sr = &Shaper{
		Workspace: w,
		Path:      path,
		Name:      name,
	}
	return
}

// Inspection

func (sr *Shaper) Inspect() (sri *ShaperInspector) {
	var err error
	var uri string
	if uri, err = sr.URI(); err != nil {
		uri = ""
	}
	sri = &ShaperInspector{
		URI:   uri,
		Name:  sr.Name,
		About: sr.About,
	}
	return
}

// Location

func (sr *Shaper) directory(subDirs ...string) (dirPath string) {
	elem := append([]string{sr.Path, "shapers", sr.Name}, subDirs...)
	dirPath = filepath.Join(elem...)
	return
}

// Data

func (sr *Shaper) URI() (uri string, err error) {
	var gruri string
	if gruri, err = utils.GitURI(sr.Path); err != nil {
		err = errors.ErrorWrap(err, "shaper URI")
		return
	}
	uri = gruri + "#" + sr.Name
	return
}

func (sr *Shaper) load(loads ...string) (err error) {
	if err = utils.YamlReadFile(sr.directory(), "shaper", sr); err != nil {
		err = errors.ErrorWrapf(err, "load %s", filepath.Join(sr.Path, sr.Name, "shaper.yaml"))
		return
	}
	err = sr.validation()
	return
}

func (sr *Shaper) validation() (err error) {
	vn := validations.NewValidation()
	if sr.About == "" {
		vn.Add("root", "must have an about")
	}
	if len(sr.Start) == 0 {
		vn.Add("root", "must have a start instruction")
	}
	for i, fmc := range sr.Shape {
		fmc.validation(fmt.Sprintf("shape %d", i), vn)
	}
	if vn.IsInvalid() {
		err = errors.ErrorWrapf(errors.ValidationError(vn.Maps()), "shaper %s", sr.Name)
		return
	}
	return
}
