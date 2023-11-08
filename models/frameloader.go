package models

import (
	"sf/app"
	"sf/database/queries"
	"sf/utils"
	"strings"

	"golang.org/x/exp/slices"
)

type FrameLoader struct {
	Frame          *Frame
	Loads          []string
	Framer         bool
	Configuration  bool
	ShapeLoads     []string
	WorkspaceLoads []string
	Preloads       []string
}

func NewFrameLoader(f *Frame, loads []string) (fl *FrameLoader) {
	fl = &FrameLoader{
		Frame: f,
		Loads: loads,
	}
	return
}

func (fl *FrameLoader) load() (err error) {
	fl.dependencies()
	fl.settle()
	fl.query()
	if err = fl.assign(); err != nil {
		return
	}
	return
}

func (fl *FrameLoader) dependencies() {
	if slices.Contains(fl.Loads, "Configuration") {
		fl.Loads = append(fl.Loads,
			"Framer",
		)
	}
	if slices.Contains(fl.Loads, "Framer") {
		fl.Loads = append(fl.Loads,
			"Workspace.Framers",
		)
	}
	utils.UniqStrings(&fl.Loads)
}

func (fl *FrameLoader) settle() {
	for _, load := range fl.Loads {
		elem := strings.SplitN(load, ".", 2)
		switch elem[0] {
		case "Shapes":
			fl.Preloads = append(fl.Preloads, "Shapes")
			if len(elem) > 1 {
				switch elem[1] {
				case "Configuration":
					fl.ShapeLoads = append(fl.ShapeLoads, elem[1])
				default:
					fl.Preloads = append(fl.Preloads, load)
				}
			}
		case "Framer":
			fl.Framer = true
		case "Configuration":
			fl.Configuration = true
		case "Workspace":
			fl.Preloads = append(fl.Preloads, "Workspace")
			if len(elem) > 1 {
				switch elem[1] {
				case "Shapers", "Framers":
					fl.WorkspaceLoads = append(fl.WorkspaceLoads, elem[1])
				default:
					fl.Preloads = append(fl.Preloads, load)
				}
			}
		default:
			fl.Preloads = append(fl.Preloads, load)
		}
	}
}

func (fl *FrameLoader) query() {
	queries.Load(fl.Frame, fl.Frame.ID, fl.Preloads...)
}

func (fl *FrameLoader) assign() (err error) {
	if len(fl.WorkspaceLoads) > 0 {
		if err = fl.LoadWorkspace(); err != nil {
			return
		}
	}
	if len(fl.ShapeLoads) > 0 {
		if err = fl.LoadShapes(); err != nil {
			return
		}
	}
	if fl.Framer {
		if err = fl.SetFramer(); err != nil {
			return
		}
	}
	if fl.Configuration {
		if err = fl.SetConfiguration(); err != nil {
			return
		}
	}
	return
}

func (fl *FrameLoader) LoadWorkspace() (err error) {
	err = fl.Frame.Workspace.Load(fl.WorkspaceLoads...)
	return
}

func (fl *FrameLoader) LoadShapes() (err error) {
	for _, d := range fl.Frame.Shapes {
		if err = d.Load(fl.ShapeLoads...); err != nil {
			return
		}
	}
	return
}

func (fl *FrameLoader) SetFramer() (err error) {
	framer := fl.Frame.Workspace.FindFramer(fl.Frame.FramerName)
	if framer == nil {
		err = app.Error(
			"framer %s does not exist in workspace %s",
			fl.Frame.FramerName,
			fl.Frame.Workspace.Name,
		)
		return
	}
	fl.Frame.Framer = framer
	return
}

func (fl *FrameLoader) SetConfiguration() (err error) {
	schema := fl.Frame.Framer.ConfigurationFormSchema()
	if err = schema.Validate(); err != nil {
		return
	}
	settings := fl.Frame.ConfigurationSettings()
	fl.Frame.Configuration = NewForm("frame", fl.Frame.Name, schema, settings)
	return
}
