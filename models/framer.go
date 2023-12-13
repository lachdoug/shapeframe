package models

import (
	"fmt"
	"path/filepath"
	"sf/app/errors"
	"sf/app/validations"
	"sf/utils"
)

type Framer struct {
	Workspace *Workspace
	Path      string
	Name      string
	About     string
	Frame     []*FormComponent
	Shape     []*FormComponent
}

type FramerInspector struct {
	Name  string
	About string
	URI   string
}

// Construction

func NewFramer(w *Workspace, path string, name string) (fr *Framer) {
	fr = &Framer{
		Workspace: w,
		Path:      path,
		Name:      name,
	}
	return
}

// Inspection

func (fr *Framer) Inspect() (fri *FramerInspector) {
	var err error
	var uri string
	if uri, err = fr.URI(); err != nil {
		uri = ""
	}
	fri = &FramerInspector{
		URI:   uri,
		Name:  fr.Name,
		About: fr.About,
	}
	return
}

// Location

func (fr *Framer) directory(subDirs ...string) (dirPath string) {
	elem := append([]string{fr.Path, "framers", fr.Name}, subDirs...)
	dirPath = filepath.Join(elem...)
	return
}

// Data

func (fr *Framer) URI() (uri string, err error) {
	var gruri string
	if gruri, err = utils.GitURI(fr.Path); err != nil {
		err = errors.ErrorWrap(err, "framer URI")
		return
	}
	uri = gruri + "#" + fr.Name
	return
}

func (fr *Framer) Load() (err error) {
	if err = utils.YamlReadFile(fr.directory(), "framer", fr); err != nil {
		err = errors.ErrorWrapf(err, "load %s", filepath.Join(fr.Path, fr.Name, "framer.yaml"))
		return
	}
	err = fr.validation()
	return
}

func (fr *Framer) validation() (err error) {
	vn := validations.NewValidation()
	if fr.About == "" {
		vn.Add("root", "must have an about")
	}
	for i, fmc := range fr.Frame {
		fmc.validation(fmt.Sprintf("frame %d", i), vn)
	}
	for i, fmc := range fr.Shape {
		fmc.validation(fmt.Sprintf("shape %d", i), vn)
	}
	if vn.IsInvalid() {
		err = errors.ErrorWrapf(errors.ValidationError(vn.Maps()), "framer %s", fr.Name)
		return
	}
	return
}
