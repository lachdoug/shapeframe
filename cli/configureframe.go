package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
	"sf/utils"
)

func configureFrame() (command any) {
	command = &cliapp.Command{
		Name:        "frame",
		Summary:     "Configure a frame",
		Aliases:     ss("f"),
		Parametizer: configureFrameParams,
		Controller:  controllers.FrameConfigurationsUpdate,
		Viewer:      configureFrameViewer,
	}
	return
}

func configureFrameParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	uc := models.UserContextNew()
	uc.Load("Workspace", "Frame")

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

	//	f.Load("Workspace.Directories")

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Name":      f.Name,
		"Config": map[string]string{
			"Color": "red",
			"Foo":   "bar",
		},
	})
	return
}

func configureFrameViewer(body map[string]any) (output string, err error) {
	result := body["Result"].(map[string]any)
	result["ConfigYaml"] = string(utils.YamlMarshal(result["Config"]))
	body["Result"] = result
	output, err = cliapp.View("frameconfigurations/update")(body)
	return
}
