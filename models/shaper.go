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
	Config    map[any]any
}

type ShaperInspector struct {
	Name  string
	About string
	URI   string
}

// Construction

func ShaperNew(w *Workspace, path string, name string) (sr *Shaper) {
	if w == nil {
		panic("Shaper Workspace is <nil>")
	}
	sr = &Shaper{
		Workspace: w,
		Path:      path,
		Name:      name,
	}
	return
}

// Inspection

func (sr *Shaper) Inspect() (sri *ShaperInspector, err error) {
	var uri string
	if uri, err = sr.URI(); err != nil {
		return
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
		err = app.Error(err, "load shaper %s in %s", sr.Name, sr.Path)
	}
	return
}

func (sr *Shaper) ConfigurationValidate(name string, values map[string]any) (err error) {
	c := NewConfiguration("shape", name, sr.Config, values)
	err = c.Validate()
	return
}
