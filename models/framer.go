package models

import (
	"path/filepath"
	"sf/app"
	"sf/utils"
)

type Framer struct {
	Workspace *Workspace
	Path      string
	Name      string
	About     string
	Config    []map[string]any
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

// Data

func (fr *Framer) URI() (uri string, err error) {
	var gruri string
	if gruri, err = utils.GitURI(fr.Path); err != nil {
		err = app.ErrorWrapf(err, "framer URI")
		return
	}
	uri = gruri + "#" + fr.Name
	return
}

func (fr *Framer) directory() (dirPath string) {
	dirPath = filepath.Join(fr.Path, "framers", fr.Name)
	return
}

func (fr *Framer) Load() (err error) {
	if err = utils.YamlReadFile(fr.directory(), "framer", fr); err != nil {
		err = app.ErrorWrapf(err, "load framer %s in %s", fr.Name, fr.Path)
	}
	return
}

// Configuration

func (fr *Framer) configurationFormSchema() (schema *FormSchema) {
	schema = NewFormSchema("framer", fr.Name, fr.Config)
	return
}
