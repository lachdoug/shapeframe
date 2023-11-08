package models

import (
	"fmt"
	"sf/app"
)

type Form struct {
	Kind     string
	Name     string
	Schema   *FormSchema
	Settings map[string]any
}

func NewForm(
	kind string,
	name string,
	schema *FormSchema,
	settings map[string]any) (c *Form) {
	c = &Form{
		Kind:     kind,
		Name:     name,
		Schema:   schema,
		Settings: settings,
	}
	return
}

func (c *Form) SettingsDetail() (settings []map[string]any) {
	settings = []map[string]any{}
	if c.Settings == nil {
		return
	}
	for _, property := range c.Schema.Properties {
		key := property["key"].(string)
		settings = append(settings, map[string]any{
			"Label": c.LabelWithDetail(key),
			"Key":   key,
			"Value": c.Value(key),
		})
	}
	return
}

func (c *Form) Value(key string) (value string) {
	value = ""
	if c.Settings[key] != nil {
		value = c.Settings[key].(string)
	}
	return
}

func (c *Form) PropertyLabelsInfo() (args []string) {
	for _, property := range c.Schema.Properties {
		label := c.LabelWithDetail(property["key"].(string))
		args = append(args, label)
	}
	return
}

func (c *Form) ValueOrDefault(key string) (value string) {
	value = ""
	if c.Settings[key] != nil {
		value = c.Settings[key].(string)
	} else {
		property := c.Schema.findProperty(key)
		if property["default"] != nil {
			value = fmt.Sprintf("%v", property["default"])
		}
	}
	return
}

func (c *Form) FieldType(key string) (fType string) {
	property := c.Schema.findProperty(key)
	if property["type"] == nil {
		fType = "string"
	} else {
		fType = property["type"].(string)
	}
	return
}

func (c *Form) IsRequired(key string) (is bool) {
	property := c.Schema.findProperty(key)
	if property["required"] == nil {
		is = false
	} else {
		is = property["required"].(bool)
	}
	return
}

func (c *Form) LabelWithDetail(key string) (label string) {
	property := c.Schema.findProperty(key)
	fType := c.FieldType(key)
	required := c.IsRequired(key)
	if property["label"] == nil {
		info := fType
		if required {
			info = info + " required"
		}
		label = fmt.Sprintf("%s (%s)", key, info)
	} else {
		info := key + " " + fType
		if required {
			info = info + " required"
		}
		label = fmt.Sprintf("%s (%s)", property["label"].(string), info)
	}
	return
}

func (c *Form) SettingsForValues(values []string) (settings map[string]any, err error) {
	settings = map[string]any{}
	for i, property := range c.Schema.Properties {
		key := property["key"].(string)
		value := ""
		if i < len(values) {
			value = values[i]
		}
		settings[key] = value
	}
	return
}

func (c *Form) Validate() (vn *app.Validation) {
	vn = &app.Validation{}
	for _, property := range c.Schema.Properties {
		key := property["key"].(string)
		required := false
		if property["required"] != nil {
			required = property["required"].(bool)
		}
		value := c.Value(key)
		if required && value == "" {
			vn.Add(key, "must not be blank")
		}
	}
	if vn.IsInvalid() {
		return
	} else {
		vn = nil
	}
	return
}
