package models

import (
	"sf/app/errors"
	"sf/database/queries"
	"sf/utils"
	"slices"
)

type FrameLoader struct {
	Frame              *Frame
	Loads              []string
	Framer             bool
	Configuration      bool
	WorkspaceLoads     []string
	ShapeLoads         []string
	ConfigurationLoads []string
	Preloads           []string
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
	err = fl.assign()
	return
}

func (fl *FrameLoader) dependencies() {
	primaries := primaryLoads(fl.Loads)
	if slices.Contains(primaries, "Configuration") {
		fl.Loads = append(fl.Loads,
			"Workspace.Framers",
			"Framer",
			"Configuration.Form",
		)
	}
	if slices.Contains(primaries, "Framer") {
		fl.Loads = append(fl.Loads,
			"Workspace.Framers",
		)
	}
}

func (fl *FrameLoader) settle() {
	utils.UniqStrings(&fl.Loads)
	for _, load := range fl.Loads {
		switch primaryLoad(load) {
		case "Shapes":
			databaseAssociation(load, "Shapes", &fl.Preloads, &fl.ShapeLoads)
		case "Framer":
			abstractAssociation(load, "Framer", &fl.Framer)
		case "Configuration":
			abstractAssociation(load, "Configuration", &fl.Configuration, &fl.ConfigurationLoads)
		case "Workspace":
			databaseAssociation(load, "Workspace", &fl.Preloads, &fl.WorkspaceLoads)
		default:
			fl.Preloads = append(fl.Preloads, load)
		}
	}
}

func (fl *FrameLoader) query() {
	if len(fl.Preloads) > 0 {
		utils.UniqStrings(&fl.Preloads)
		queries.Load(fl.Frame, fl.Frame.ID, fl.Preloads...)
	}
}

func (fl *FrameLoader) assign() (err error) {
	if err = fl.loadWorkspace(); err != nil {
		return
	}
	if err = fl.loadShapes(); err != nil {
		return
	}
	if err = fl.loadFramer(); err != nil {
		return
	}
	if err = fl.loadConfiguration(); err != nil {
		return
	}
	return
}

func (fl *FrameLoader) loadWorkspace() (err error) {
	if len(fl.WorkspaceLoads) > 0 {
		err = fl.Frame.Workspace.Load(fl.WorkspaceLoads...)
	}
	return
}

func (fl *FrameLoader) loadShapes() (err error) {
	if len(fl.ShapeLoads) > 0 {
		for _, s := range fl.Frame.Shapes {
			if err = s.Load(fl.ShapeLoads...); err != nil {
				return
			}
		}
	}
	return
}

func (fl *FrameLoader) loadFramer() (err error) {
	if fl.Framer {
		if err = fl.SetFramer(); err != nil {
			return
		}
	}
	return
}

func (fl *FrameLoader) loadConfiguration() (err error) {
	if fl.Configuration {
		if err = fl.setConfiguration(); err != nil {
			return
		}
		err = fl.Frame.Configuration.load(fl.ConfigurationLoads...)
	}
	return
}

func (fl *FrameLoader) setConfiguration() (err error) {
	c := NewConfiguration(fl.Frame.ID, "frame", "frame", fl.Frame.Framer.Frame)
	queries.Lookup(c)
	fl.Frame.Configuration = c
	return
}

func (fl *FrameLoader) SetFramer() (err error) {
	framer := fl.Frame.Workspace.FindFramer(fl.Frame.FramerName)
	if framer == nil {
		err = errors.Errorf(
			"framer %s does not exist in workspace %s",
			fl.Frame.FramerName,
			fl.Frame.Workspace.Name,
		)
		return
	}
	fl.Frame.Framer = framer
	return
}
