package models

import (
	"fmt"
	"path/filepath"
	"sf/app/streams"
	"sf/utils"
)

type Component struct {
	Name        string
	Composition *Composition
	Path        string
	Shape       *Shape
	Stream      *streams.Stream
}

func NewComponent(o *Composition, s *Shape, st *streams.Stream) (ct *Component) {
	ct = &Component{
		Name:        fmt.Sprintf("%s.%s", o.Frame.Name, s.Name),
		Composition: o,
		Shape:       s,
		Stream:      st,
	}
	ct.setPath()
	return
}

// Location

func (ct *Component) setPath() {
	ct.Path = ct.Composition.directory("shapes", ct.Shape.Name)
}

func (ct *Component) directory(subDirs ...string) (dirPath string) {
	elem := append([]string{ct.Path}, subDirs...)
	dirPath = filepath.Join(elem...)
	return
}

// Prepare

func (ct *Component) prepare() (err error) {
	ct.Stream.SubHeading("Prepare shape %s.%s", ct.Composition.Frame.Name, ct.Shape.Name)
	ct.createShapeDirectory()
	ct.createConfigYaml()
	ct.copyShaper()
	ct.copyFramer()
	return
}

func (ct *Component) createShapeDirectory() {
	ct.Stream.DotPoint("create directory %s", ct.Shape.Name)
	utils.MakeDir(ct.directory())
}

func (ct *Component) createConfigYaml() {
	ct.Stream.DotPoint("create shape file")
	utils.YamlWriteFile(ct.directory(), "shape", ct.Shape.Output())
}

func (ct *Component) copyShaper() {
	s := ct.Shape
	sr := s.Shaper
	ct.Stream.DotPoint("copy shaper %s", sr.Name)
	utils.CopyFile(sr.directory("build-shape"), ct.directory("build-shape"))
	utils.CopyDir(sr.directory("shape"), ct.directory("shape"))
}

func (ct *Component) copyFramer() {
	f := ct.Composition.Frame
	fr := f.Framer
	ct.Stream.DotPoint("copy framer %s", fr.Name)
	utils.CopyFile(fr.directory("build-frame-shape"), ct.directory("build-frame-shape"))
	utils.CopyFile(fr.directory("build-frame-shape"), ct.directory("build-frame-shape"))
	utils.CopyDir(fr.directory("frame-shape"), ct.directory("frame-shape"))
}

// Build

func (ct *Component) build() (err error) {
	ct.Stream.SubHeading("Build shape %s.%s", ct.Composition.Frame.Name, ct.Shape.Name)
	if err = ct.execBuildShapeScript(); err != nil {
		return
	}
	if err = ct.execBuildFrameShapeScript(); err != nil {
		return
	}
	return
}

func (ct *Component) execBuildShapeScript() (err error) {
	filePath := ct.directory("build-shape")
	ct.Stream.DotPoint("execute build shape")
	err = ct.Composition.Exec(filePath)
	return
}

func (ct *Component) execBuildFrameShapeScript() (err error) {
	filePath := ct.directory("build-frame-shape")
	ct.Stream.DotPoint("execute build-frame-shape")
	err = ct.Composition.Exec(filePath)
	return
}
