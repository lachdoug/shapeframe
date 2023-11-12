package models

import (
	"sf/app"
	"sf/database/queries"
	"sf/utils"

	"golang.org/x/exp/slices"
)

type ShapeLoader struct {
	Shape              *Shape
	Loads              []string
	Shaper             bool
	Configuration      bool
	ConfigurationLoads []string
	FrameLoads         []string
	Preloads           []string
}

func NewShapeLoader(s *Shape, loads []string) (sl *ShapeLoader) {
	sl = &ShapeLoader{
		Shape: s,
		Loads: loads,
	}
	return
}

func (sl *ShapeLoader) load() (err error) {
	sl.dependencies()
	sl.settle()
	sl.query()
	err = sl.assign()
	return
}

func (sl *ShapeLoader) dependencies() {
	primaries := primaryLoads(sl.Loads)
	if slices.Contains(primaries, "Configuration") {
		sl.Loads = append(sl.Loads,
			"Frame.Workspace.Shapers",
			"Shaper",
			"Configuration.Form",
		)
	}
	if slices.Contains(primaries, "Shaper") {
		sl.Loads = append(sl.Loads,
			"Frame.Workspace.Shapers",
		)
	}
}

func (sl *ShapeLoader) settle() {
	utils.UniqStrings(&sl.Loads)
	for _, load := range sl.Loads {
		switch primaryLoad(load) {
		case "Shaper":
			abstractAssociation(load, "Shaper", &sl.Shaper)
		case "Configuration":
			abstractAssociation(load, "Configuration", &sl.Configuration, &sl.ConfigurationLoads)
		case "Frame":
			databaseAssociation(load, "Frame", &sl.Preloads, &sl.FrameLoads)
		default:
			sl.Preloads = append(sl.Preloads, load)
		}
	}
}

func (sl *ShapeLoader) query() {
	utils.UniqStrings(&sl.Preloads)
	queries.Load(sl.Shape, sl.Shape.ID, sl.Preloads...)
}

func (sl *ShapeLoader) assign() (err error) {
	if err = sl.loadFrame(); err != nil {
		return
	}
	if err = sl.loadShaper(); err != nil {
		return
	}
	if err = sl.loadConfiguration(); err != nil {
		return
	}
	return
}

func (sl *ShapeLoader) loadFrame() (err error) {
	if len(sl.FrameLoads) > 0 {
		err = sl.Shape.Frame.Load(sl.FrameLoads...)
	}
	return
}

func (sl *ShapeLoader) loadShaper() (err error) {
	if sl.Shaper {
		if err = sl.setShaper(); err != nil {
			return
		}
	}
	return
}

func (sl *ShapeLoader) loadConfiguration() (err error) {
	if sl.Configuration {
		if err = sl.setConfiguration(); err != nil {
			return
		}
		if len(sl.ConfigurationLoads) > 0 {
			if err = sl.Shape.Configuration.Load(sl.ConfigurationLoads...); err != nil {
				return
			}
		}
	}
	return
}

func (sl *ShapeLoader) setConfiguration() (err error) {
	c := &Configuration{
		OwnerID:   sl.Shape.ID,
		OwnerType: "Shape",
	}
	queries.Lookup(c)
	if c.ID == 0 {
		c = NewConfiguration(sl.Shape)
	}
	c.Owner = sl.Shape
	sl.Shape.Configuration = c
	return
}

func (sl *ShapeLoader) setShaper() (err error) {
	shaper := sl.Shape.Frame.Workspace.FindShaper(sl.Shape.ShaperName)
	if shaper == nil {
		err = app.Error(
			"shaper %s does not exist in workspace %s",
			sl.Shape.ShaperName,
			sl.Shape.Frame.Workspace.Name,
		)
		return
	}
	sl.Shape.Shaper = shaper
	return
}
