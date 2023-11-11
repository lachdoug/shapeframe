package models

import (
	"path/filepath"
	"sf/app"
	"sf/utils"
)

type Shaper struct {
	Workspace *Workspace
	Path      string
	Name      string
	About     string
	Config    []map[string]any
	Connect   []*ShaperConnect
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

// Data

func (sr *Shaper) URI() (uri string, err error) {
	var gruri string
	if gruri, err = utils.GitRepoURI(sr.Path); err != nil {
		err = app.ErrorWrapf(err, "shaper URI")
		return
	}
	uri = gruri + "#" + sr.Name
	return
}

func (sr *Shaper) directory() (dirPath string) {
	dirPath = filepath.Join(sr.Path, "shapers", sr.Name)
	return
}

func (sr *Shaper) Load() (err error) {
	if err = utils.YamlReadFile(sr.directory(), "shaper", sr); err != nil {
		err = app.ErrorWrapf(err, "load shaper %s in %s", sr.Name, sr.Path)
	}
	return
}

// Configuration

func (sr *Shaper) configurationFormSchema() (schema *FormSchema) {
	schema = NewFormSchema("shaper", sr.Name, sr.Config)
	return
}
