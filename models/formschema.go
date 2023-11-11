package models

import (
	"fmt"
	"sf/app"
	"sf/utils"
)

type FormSchema struct {
	Kind       string
	Name       string
	Properties []map[string]any
}

func NewFormSchema(
	kind string,
	name string,
	properties []map[string]any) (fs *FormSchema) {
	fs = &FormSchema{
		Kind:       kind,
		Name:       name,
		Properties: properties,
	}
	return
}

func (fs *FormSchema) findProperty(key string) (property map[string]any) {
	for _, property := range fs.Properties {
		if property["key"] == key {
			return property
		}
	}
	property = nil
	return
}

func (fs *FormSchema) Validation() (err error) {
	msgs := []string{}
	var assertionOk bool
	for _, property := range fs.Properties {
		if property["key"] == nil {
			msgs = append(
				msgs,
				fmt.Sprintf("property must have a key: %s", utils.JsonMarshal(property)),
			)
		}
		if property["required"] != nil {
			_, assertionOk = property["required"].(bool)
			if !assertionOk {
				msgs = append(
					msgs,
					fmt.Sprintf("required attribute must be a boolean: %s", utils.JsonMarshal(property)),
				)
			}
		}
		if property["label"] != nil {
			_, assertionOk = property["label"].(string)
			if !assertionOk {
				msgs = append(
					msgs,
					fmt.Sprintf("label attribute must be a string: %s", utils.JsonMarshal(property)),
				)
			}
		}
	}
	if len(msgs) > 0 {
		errMsg := fmt.Sprintf("%s %s configuration schema validation:", fs.Kind, fs.Name)
		for _, msg := range msgs {
			errMsg = errMsg + "\n  - " + msg
		}
		err = app.Error(errMsg)
	}
	return
}
