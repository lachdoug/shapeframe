package models

import (
	"slices"
)

type ConfigurationLoader struct {
	Configuration *Configuration
	Form          bool
	Loads         []string
	Preloads      []string
}

func NewConfigurationLoader(c *Configuration, loads []string) (cl *ConfigurationLoader) {
	cl = &ConfigurationLoader{
		Configuration: c,
		Loads:         loads,
	}
	return
}

func (cl *ConfigurationLoader) load() (err error) {
	cl.settle()
	err = cl.assign()
	return
}

func (cl *ConfigurationLoader) settle() {
	if slices.Contains(cl.Loads, "Form") {
		cl.Form = true
	}
}

func (cl *ConfigurationLoader) assign() (err error) {
	if err = cl.loadForm(); err != nil {
		return
	}
	return
}

func (cl *ConfigurationLoader) loadForm() (err error) {
	if slices.Contains(cl.Loads, "Form") {
		if err = cl.setForm(); err != nil {
			return
		}
	}
	return
}

func (cl *ConfigurationLoader) setForm() (err error) {
	schema := cl.Configuration.FormSchema()
	settings := cl.Configuration.Settings()
	cl.Configuration.Form = NewForm(schema, settings)
	return
}
