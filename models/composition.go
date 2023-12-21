package models

import (
	"path/filepath"
	"sf/app/dirs"
	"sf/app/scripts"
	"sf/app/streams"
	"sf/utils"
)

type Composition struct {
	Path       string
	Frame      *Frame
	Stream     *streams.Stream
	Components []*Component
}

func NewComposition(f *Frame, st *streams.Stream) (cn *Composition) {
	cn = &Composition{
		Frame:  f,
		Stream: st,
	}
	cn.setPath()
	return
}

// Location

func (cn *Composition) setPath() {
	cn.Path = dirs.WorkspaceDir(filepath.Join(
		"composition",
		cn.Frame.Name,
	))
}

func (cn *Composition) directory(subDirs ...string) (dirPath string) {
	elem := append([]string{cn.Path}, subDirs...)
	dirPath = filepath.Join(elem...)
	return
}

// Action

func (cn *Composition) apply() {
	var err error
	defer cn.Stream.Close()
	cn.Stream.Heading("Composition apply %s", cn.Frame.Name)
	cn.setComponents()
	if err = cn.prepare(); err != nil {
		cn.Stream.Error(err.Error())
		return
	}
	if err = cn.prepareComponents(); err != nil {
		cn.Stream.Error(err.Error())
		return
	}
	if err = cn.buildComponents(); err != nil {
		cn.Stream.Error(err.Error())
		return
	}
	if err = cn.build(); err != nil {
		cn.Stream.Error(err.Error())
		return
	}
	if err = cn.start(); err != nil {
		cn.Stream.Error(err.Error())
		return
	}
	if err = cn.purge(); err != nil {
		cn.Stream.Error(err.Error())
		return
	}
}

func (cn *Composition) destroy() {
	var err error
	defer cn.Stream.Close()
	cn.Stream.Heading("Composition destroy %s", cn.Frame.Name)
	cn.setComponents()
	if err = cn.stop(); err != nil {
		cn.Stream.Error(err.Error())
		return
	}
	if err = cn.purge(); err != nil {
		cn.Stream.Error(err.Error())
		return
	}
}

// Prepare

func (cn *Composition) prepare() (err error) {
	cn.Stream.SubHeading("Prepare frame %s", cn.Frame.Name)
	cn.createFrameDirectory()
	cn.createConfigYaml()
	cn.copyFramer()
	return
}

func (cn *Composition) createFrameDirectory() {
	cn.Stream.DotPoint("create directory %s", cn.Frame.Name)
	utils.MakeDir(cn.directory())
}

func (cn *Composition) createConfigYaml() {
	cn.Stream.DotPoint("create frame file")
	utils.YamlWriteFile(cn.directory(), "frame", cn.Frame.Output())
}

func (cn *Composition) copyFramer() {
	f := cn.Frame
	fr := f.Framer
	cn.Stream.DotPoint("copy framer %s", fr.Name)
	utils.CopyFile(fr.directory("build-frame"), cn.directory("build-frame"))
	utils.CopyFile(fr.directory("start-frame"), cn.directory("start-frame"))
	utils.CopyFile(fr.directory("stop-frame"), cn.directory("stop-frame"))
	utils.CopyFile(fr.directory("purge-frame"), cn.directory("purge-frame"))
	utils.CopyDir(fr.directory("frame"), cn.directory("frame"))
}

// Build

func (cn *Composition) build() (err error) {
	cn.Stream.SubHeading("Build frame %s", cn.Frame.Name)
	if err = cn.execBuildFrameScript(); err != nil {
		return
	}
	return
}

func (cn *Composition) execBuildFrameScript() (err error) {
	filePath := cn.directory("build-frame")
	cn.Stream.DotPoint("execute build-frame")
	err = cn.Exec(filePath)
	return
}

// Start

func (cn *Composition) start() (err error) {
	cn.Stream.SubHeading("Start frame %s", cn.Frame.Name)
	if err = cn.execStartFrameScript(); err != nil {
		return
	}
	return
}

func (cn *Composition) execStartFrameScript() (err error) {
	filePath := cn.directory("start-frame")
	cn.Stream.DotPoint("execute start-frame")
	err = cn.Exec(filePath)
	return
}

// Stop

func (cn *Composition) stop() (err error) {
	cn.Stream.SubHeading("Stop frame %s", cn.Frame.Name)
	if err = cn.execStopFrameScript(); err != nil {
		return
	}
	return
}

func (cn *Composition) execStopFrameScript() (err error) {
	filePath := cn.directory("stop-frame")
	cn.Stream.DotPoint("execute stop-frame")
	err = cn.Exec(filePath)
	return
}

// Purge

func (cn *Composition) purge() (err error) {
	cn.Stream.SubHeading("Purge frame %s", cn.Frame.Name)
	if err = cn.execPurgeFrameScript(); err != nil {
		return
	}
	return
}

func (cn *Composition) execPurgeFrameScript() (err error) {
	filePath := cn.directory("purge-frame")
	cn.Stream.DotPoint("execute purge-frame")
	err = cn.Exec(filePath)
	return
}

// Components

func (cn *Composition) setComponents() {
	for _, s := range cn.Frame.Shapes {
		ct := NewComponent(cn, s, cn.Stream)
		cn.Components = append(cn.Components, ct)
	}
}

func (cn *Composition) prepareComponents() (err error) {
	for _, ct := range cn.Components {
		if err = ct.prepare(); err != nil {
			return
		}
	}
	return
}

func (cn *Composition) buildComponents() (err error) {
	for _, ct := range cn.Components {
		if err = ct.build(); err != nil {
			return
		}
	}
	return
}

// Scripts

func (cn *Composition) Exec(filePath string) (err error) {
	var isOutput bool
	if utils.IsFile(filePath) {
		if isOutput, err = scripts.Exec(filePath, cn.Stream); err != nil {
			cn.Stream.ScriptError(err)
			return
		}
		cn.Stream.ScriptSuccess(isOutput)
	} else {
		cn.Stream.ScriptMissing()
	}
	return
}
