package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
	"sf/utils"
)

func configureShape() (command any) {
	command = &cliapp.Command{
		Name:        "shape",
		Summary:     "Configure a shape",
		Aliases:     ss("s"),
		Parametizer: configureShapeParams,
		Controller:  controllers.ShapeConfigurationsUpdate,
		Viewer:      configureShapeViewer,
	}
	return
}

func configureShapeParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	uc := models.UserContextNew()
	uc.Load("Workspace", "Frame", "Shape")

	w := uc.Workspace
	if w == nil {
		err = app.Error(nil, "no workspace context")
		return
	}
	f := uc.Frame
	if f == nil {
		err = app.Error(nil, "no frame context")
		return
	}
	s := uc.Shape
	if s == nil {
		err = app.Error(nil, "no shape context")
		return
	}

	s.Load("Frame.Workspace.Directories")
	// cs := s.ConfigSchema()
	// c := s.Config()

	// fmt.Println("CONFIG SCHEMA", cs)
	// fmt.Println("CONFIG", c)

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Frame":     f.Name,
		"Name":      s.Name,
		"Config": map[string]string{
			"color": "red",
			"foo":   "bar",
		},
	})
	return
}

func configureShapeViewer(body map[string]any) (output string, err error) {
	result := body["Result"].(map[string]any)
	result["ConfigYaml"] = string(utils.YamlMarshal(result["Config"]))
	body["Result"] = result
	output, err = cliapp.View("shapeconfigurations/update")(body)
	return
}
