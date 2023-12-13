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
	gorm.Model
	OwnerID      uint
	OwnerType    string
	ConfigType   string
	SettingsJSON datatypes.JSON
	Settings     map[string]string `gorm:"-"`
	Form         []*FormComponent  `gorm:"-"`
	// CipherKey    string                   `gorm:"-"`
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

func (c *Configuration) load(loads ...string) (err error) {
	cl := NewConfigurationLoader(c, loads)
	cl.load()
	return
}

func (c *Configuration) Update(updates map[string]string) (vn *validations.Validation) {
	c.Settings = updates
	c.loadFormComponents()
	if vn = c.validation(); vn.IsInvalid() {
		return
	}
	c.SettingsJSON = utils.JsonMarshal(c.Settings)
	c.Save()
	return
}

func (c *Configuration) loadFormComponents() {
	for _, fmc := range c.Form {
		fmc.Load()
	}
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

func (c *Configuration) Info() (ms []map[string]string) {
	ci := NewConfigurationInfo(c)
	for _, fmc := range c.Form {
		c.componentInfo(fmc, ci)
	}
	ms = ci.Maps
	return
}

func (c *Configuration) componentInfo(fmc *FormComponent, ci *ConfigurationInfo) {
	if c.isDependMatch(fmc.Depend) {
		if fmc.Type == "row" {
			c.rowInfo(fmc, ci)
		} else {
			ci.Add(fmc)
		}
	}
}

func (c *Configuration) rowInfo(fmr *FormComponent, ci *ConfigurationInfo) {
	for _, fmc := range fmr.Cols {
		c.componentInfo(fmc, ci)
	}
}

func (c *Configuration) Populate(vs []string) (u map[string]string, err error) {
	// u = map[string]string{}
	// if len(vs) != len(c.Properties) {
	// 	err = errors.Errorf(
	// 		"received %d configuration settings when %d expected",
	// 		len(vs),
	// 		len(c.Properties),
	// 	)
	// 	return
	// }
	// for i, cp := range c.Properties {
	// 	v := ""
	// 	if i < len(vs) {
	// 		v = vs[i]
	// 	}
	// 	u[cp.Key] = v
	// }
	return
}

// Record

func (c *Configuration) Save() {
	queries.Save(c)
}
