package models

import (
	"regexp"
	"sf/app/validations"
	"sf/database/queries"
	"sf/utils"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Configuration struct {
	gorm.Model   `yaml:"-"`
	OwnerID      uint              `yaml:"-"`
	OwnerType    string            `yaml:"-"`
	ConfigType   string            `yaml:"-"`
	SettingsJSON datatypes.JSON    `yaml:"-"`
	Settings     map[string]string `gorm:"-"`
	Info         []map[string]any  `gorm:"-"`
	Form         []*FormComponent  `gorm:"-"`
}

type ConfigurationInspector struct {
	Settings     map[string]string
	SettingsMaps []map[string]string
	Info         []map[string]any
}

func NewConfiguration(
	ownerId uint,
	ownerType string,
	configType string,
	form []*FormComponent,
) (c *Configuration) {
	c = &Configuration{
		OwnerID:    ownerId,
		OwnerType:  ownerType,
		ConfigType: configType,
		Form:       form,
	}
	return
}

func (c *Configuration) Inspect() (ci *ConfigurationInspector) {
	ci = &ConfigurationInspector{
		Settings:     c.Settings,
		SettingsMaps: c.SettingsMaps(),
		Info:         c.Info,
	}
	return
}

func (c *Configuration) load(loads ...string) (err error) {
	cl := NewConfigurationLoader(c, loads)
	cl.load()
	return
}

func (c *Configuration) Update(updates map[string]string) (vn *validations.Validation) {
	c.Settings = updates
	if vn = c.validation(); vn.IsInvalid() {
		return
	}
	c.SettingsJSON = utils.JsonMarshal(c.Settings)
	c.Save()
	return
}

func (c *Configuration) validation() (vn *validations.Validation) {
	vn = validations.NewValidation()
	for _, fmc := range c.Form {
		c.componentValidation(fmc, vn)
	}
	return
}

func (c *Configuration) componentValidation(fmc *FormComponent, vn *validations.Validation) {
	if c.isDependMatch(fmc.Depend) {
		if fmc.Type == "row" {
			c.rowValidation(fmc, vn)
		} else {
			c.fieldValidation(fmc, vn)
		}
	}
}

func (c *Configuration) isDependMatch(depend *FormDepend) (is bool) {
	if depend == nil {
		is = true
		return
	}
	var r *regexp.Regexp
	var err error
	if r, err = regexp.Compile(depend.Pattern); err != nil {
		panic(err)
	}
	is = r.MatchString(c.Settings[depend.Key])
	return
}

func (c *Configuration) rowValidation(fmr *FormComponent, vn *validations.Validation) {
	for _, fmc := range fmr.Cols {
		c.componentValidation(fmc, vn)
	}
}

func (c *Configuration) fieldValidation(fmc *FormComponent, vn *validations.Validation) {
	v := c.Settings[fmc.Key]
	fmc.ValueValidation(v, vn)
}

func (c *Configuration) SettingsMaps() (ms []map[string]string) {
	ms = []map[string]string{}
	for k, v := range c.Settings {
		m := map[string]string{}
		m[k] = v
		ms = append(ms, m)
	}
	return
}

// Record

func (c *Configuration) Save() {
	queries.Save(c)
}
