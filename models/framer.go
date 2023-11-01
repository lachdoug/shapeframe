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
	Config    map[any]any
}

type FramerInspector struct {
	Name  string
	About string
	URI   string
}

// Construction

func FramerNew(w *Workspace, path string, name string) (fr *Framer) {
	if w == nil {
		panic("Framer Workspace is <nil>")
	}
	fr = &Framer{
		Workspace: w,
		Path:      path,
		Name:      name,
	}
	return
}

// Inspection

func (fr *Framer) Inspect() (fri *FramerInspector, err error) {
	var uri string
	if uri, err = fr.URI(); err != nil {
		return
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
	if gruri, err = utils.GitRepoURI(fr.Path); err != nil {
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
		err = app.Error(err, "load framer %s in %s", fr.Name, fr.Path)
	}
	return
}

func (fr *Framer) ConfigurationValidate(name string, values map[string]any) (err error) {
	c := NewConfiguration("frame", name, fr.Config, values)
	err = c.Validate()
	return
}
