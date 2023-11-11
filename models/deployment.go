package models

import (
	"path/filepath"
	"sf/utils"
)

type Deployment struct {
	Shape  *Shape
	Stream *utils.Stream
}

func NewDeployment(s *Shape, st *utils.Stream) (o *Deployment) {
	o = &Deployment{
		Shape:  s,
		Stream: st,
	}
	return
}

// Build

func (d *Deployment) build() (err error) {
	d.write("Build %s\n", d.Shape.Name)
	d.write("Configuration:\n")
	for _, setting := range d.Shape.Configuration.Details() {
		d.write("  %s: %s\n", setting["Key"], setting["Value"])
	}

	if err = d.createBuildDirectory(); err != nil {
		return
	}
	if err = d.createConfigYaml(); err != nil {
		return
	}
	return
}

// Location

func (d *Deployment) createBuildDirectory() (err error) {
	utils.MakeDir(d.directory())
	return
}

func (d *Deployment) directory() (dirPath string) {
	dirPath = utils.TempDir(filepath.Join("build", d.Shape.Name))
	return
}

func (d *Deployment) createConfigYaml() (err error) {
	utils.YamlWriteFile(d.directory(), "config", d.Shape.Configuration.Settings())
	return
}

// Stream

func (d *Deployment) write(format string, a ...any) {
	d.Stream.Writef(format, a...)
}
