package models

import (
	"slices"
)

type ConfigurationLoader struct {
	Configuration *Configuration
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
	if err = cl.assign(); err != nil {
		return
	}
	return
}

func (cl *ConfigurationLoader) assign() (err error) {
	if slices.Contains(cl.Loads, "Form") {
		if err = cl.SetForm(); err != nil {
			return
		}
	}
	return
}

func (cl *ConfigurationLoader) SetForm() (err error) {
	schema := cl.Configuration.FormSchema()
	settings := cl.Configuration.Settings()
	cl.Configuration.Form = NewForm(schema, settings)
	return
}
