package models

import (
	"sf/app"
	"sf/database/queries"
	"sf/utils"
	"strings"

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
	if err = sl.assign(); err != nil {
		return
	}
	return
}

func (sl *ShapeLoader) dependencies() {
	if slices.Contains(sl.Loads, "Configuration") {
		sl.Loads = append(sl.Loads,
			"Shaper",
			"Configuration.Form",
		)
	}
	if slices.Contains(sl.Loads, "Shaper") {
		sl.Loads = append(sl.Loads,
			"Frame.Workspace.Shapers",
		)
	}
	utils.UniqStrings(&sl.Loads)
}

func (sl *ShapeLoader) settle() {
	for _, load := range sl.Loads {
		elem := strings.SplitN(load, ".", 2)
		switch elem[0] {
		case "Shaper":
			sl.Shaper = true
		case "Configuration":
			sl.Configuration = true
			if len(elem) > 1 {
				sl.ConfigurationLoads = append(sl.ConfigurationLoads, elem[1])
			}
		case "Frame":
			sl.Preloads = append(sl.Preloads, "Frame")
			if len(elem) > 1 {
				sl.FrameLoads = append(sl.FrameLoads, elem[1])
			}
		default:
			sl.Preloads = append(sl.Preloads, load)
		}
	}
	utils.UniqStrings(&sl.Preloads)
}

func (sl *ShapeLoader) query() {
	queries.Load(sl.Shape, sl.Shape.ID, sl.Preloads...)
}

func (sl *ShapeLoader) assign() (err error) {
	if len(sl.FrameLoads) > 0 {
		if err = sl.LoadFrame(); err != nil {
			return
		}
	}
	if sl.Shaper {
		if err = sl.SetShaper(); err != nil {
			return
		}
	}
	if sl.Configuration {
		if err = sl.SetConfiguration(); err != nil {
			return
		}
	}
	if len(sl.ConfigurationLoads) > 0 {
		if err = sl.LoadConfiguration(); err != nil {
			return
		}
	}
	return
}

func (sl *ShapeLoader) SetConfiguration() (err error) {
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

func (sl *ShapeLoader) SetShaper() (err error) {
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

func (sl *ShapeLoader) LoadFrame() (err error) {
	err = sl.Shape.Frame.Load(sl.FrameLoads...)
	return
}

func (sl *ShapeLoader) LoadConfiguration() (err error) {
	err = sl.Shape.Configuration.Load(sl.ConfigurationLoads...)
	return
}
