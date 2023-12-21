package models

import (
	"sf/utils"

	"golang.org/x/exp/slices"
)

type ConfigurationLoader struct {
	Configuration *Configuration
	Form          bool
	Info          bool
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
}

func (cl *ConfigurationLoader) settle() {
	if slices.Contains(cl.Loads, "Form") {
		cl.Form = true
	}
	if slices.Contains(cl.Loads, "Info") {
		cl.Form = true
		cl.Info = true
	}
}

func (cl *ConfigurationLoader) assign() {
	cl.loadSettings()
	cl.loadForm()
	cl.loadInfo()
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
		for _, fmc := range cl.Configuration.Form {
			fmc.Load()
		}
	}
}

func (cl *ConfigurationLoader) loadInfo() {
	if cl.Info {
		for _, fmc := range cl.Configuration.Form {
			if fmc.Type == "row" {
				cl.loadFormRowInfo(fmc)
			} else {
				cl.loadFormComponentInfo(fmc)
			}
		}
	}
}

func (cl *ConfigurationLoader) loadFormRowInfo(fmr *FormComponent) {
	if cl.Configuration.isDependMatch(fmr.Depend) {
		for _, fmc := range fmr.Cols {
			cl.loadFormComponentInfo(fmc)
		}
	}
}

func (cl *ConfigurationLoader) loadFormComponentInfo(fmc *FormComponent) {
	if cl.Configuration.isDependMatch(fmc.Depend) {
		cl.Configuration.Info = append(cl.Configuration.Info, cl.settingInfo(fmc))
	}
}

func (cl *ConfigurationLoader) settingInfo(fmc *FormComponent) (info map[string]any) {
	info = map[string]any{
		"Type":    fmc.Type,
		"Key":     fmc.Key,
		"Value":   cl.Configuration.Settings[fmc.Key],
		"Label":   fmc.Label,
		"Options": cl.settingInfoOptions(fmc),
	}
	return
}

func (cl *ConfigurationLoader) settingInfoOptions(fmc *FormComponent) (opts map[string]string) {
	opts = map[string]string{}
	for _, fmcOpt := range fmc.Options {
		opts[fmcOpt.Value] = fmcOpt.Label
	}
	return
}
