package models

import "strings"

type ConfigurationInfo struct {
	Configuration *Configuration
	Maps          []map[string]string
}

func NewConfigurationInfo(c *Configuration) (ci *ConfigurationInfo) {
	ci = &ConfigurationInfo{Configuration: c}
	return
}

func (ci *ConfigurationInfo) Add(fmc *FormComponent) {
	setting := ci.Configuration.Settings[fmc.Key]
	ci.Maps = append(
		ci.Maps,
		map[string]string{
			"Type":    fmc.Type,
			"Key":     fmc.Key,
			"Setting": setting,
			"Options": strings.Join(ci.optionValues(fmc), " "),
		},
	)
}

func (ci *ConfigurationInfo) optionValues(fmc *FormComponent) (vals []string) {
	vals = []string{}
	for _, opt := range fmc.Options {
		vals = append(vals, opt.Value)
	}
	return
}
