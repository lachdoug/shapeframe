package models

import (
	"sf/app"
	"sf/database/queries"
	"sf/utils"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Configuration struct {
	gorm.Model
	OwnerID   uint
	OwnerType string
	JSON      datatypes.JSON
	Owner     Configurationer `gorm:"-"`
	Form      *Form           `gorm:"-"`
}

func NewConfiguration(owner Configurationer) (c *Configuration) {
	c = &Configuration{
		Owner:     owner,
		OwnerID:   owner.readID(),
		OwnerType: owner.readType(),
	}
	return
}

func (c *Configuration) IsSet() (is bool) {
	if c.JSON.String() != "" {
		is = true
	}
	return
}

func (c *Configuration) Load(loads ...string) (err error) {
	cl := NewConfigurationLoader(c, loads)
	err = cl.load()
	return
}

func (c *Configuration) Update(update map[string]any) (vn *app.Validation, err error) {
	c.JSON = datatypes.JSON(utils.JsonMarshal(update))
	c.Save()
	if err = c.Load("Form"); err != nil {
		return
	}
	vn = c.Form.Validation()
	return
}

func (c *Configuration) Settings() (settings map[string]any) {
	if c.IsSet() {
		settings = map[string]any{}
		utils.JsonUnmarshal([]byte(c.JSON.String()), &settings)
	}
	return
}

func (c *Configuration) FormSchema() (schema *FormSchema) {
	schema = c.Owner.configurationFormSchema()
	return
}

func (c *Configuration) Details() (details []map[string]any) {
	details = c.Form.Details()
	return
}

func (c *Configuration) Save() {
	queries.Save(c)
}
