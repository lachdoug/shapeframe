package models

import (
	"sf/utils"

	"golang.org/x/exp/slices"
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

func (cl *ConfigurationLoader) load() {
	cl.settle()
	cl.assign()
	return
}

func (cl *ConfigurationLoader) settle() {
	if slices.Contains(cl.Loads, "Form") {
		cl.Form = true
	}
}

func (cl *ConfigurationLoader) assign() {
	cl.loadSettings()
	cl.loadForm()
}

func (cl *ConfigurationLoader) loadSettings() {
	ss := map[string]string{}
	json := cl.Configuration.SettingsJSON.String()
	if json != "" {
		utils.JsonUnmarshal([]byte(json), &ss)
	}
	cl.Configuration.Settings = ss
}

func (cl *ConfigurationLoader) loadForm() {
	if cl.Form {
		cl.Configuration.loadFormComponents()
	}
}
