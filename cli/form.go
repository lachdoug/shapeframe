package cli

import (
	"sf/cli/prompting"
	"sf/models"
)

type Form struct {
	Model *models.Form
}

type FormControl struct {
	Key   string
	Value string
	Label string
}

func (cf *Form) prompts() (answers map[string]any, err error) {
	answers = map[string]any{}
	for _, c := range cf.controls() {
		key := c.Key
		answer := ""
		if answer, err = cf.prompt(c); err != nil {
			return
		}
		answers[key] = answer
	}
	return
}

func (cf *Form) prompt(control *FormControl) (answer string, err error) {
	suggest := control.Value
	question := control.Label
	answer, err = prompting.Prompt(question+"?", suggest)
	return
}

func (cf *Form) controls() (controls []*FormControl) {
	for _, property := range cf.Model.Schema.Properties {
		key := property["key"].(string)
		value := cf.Model.ValueOrDefault(key)
		label := cf.Model.LabelWithDetail(key)
		controls = append(controls, &FormControl{
			Key:   key,
			Label: label,
			Value: value,
		})
	}
	return
}
