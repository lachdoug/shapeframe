package cli

import (
	"sf/app/io"
	"sf/cli/cliapp"
	"sf/controllers"
)

func orchestrate() (command any) {
	command = &cliapp.Command{
		Name:    "orchestrate",
		Summary: "Orchestrate a frame",
		Aliases: ss("o"),
		Usage: ss(
			"sf orchestrate [options]",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Is required if -workspace flag is set",
			"  Uses frame context when not provided",
		),
		Flags: ss(
			"string", "frame", "Frame name",
			"string", "workspace", "Workspace name",
		),
		Handler:    orchestrateHandler,
		Controller: controllers.FrameOrchestrationsCreate,
		Viewer:     orchestrateViewer,
	}
	return
}

func orchestrateHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.FrameOrchestrationsCreateParams{
			Workspace: context.StringFlag("workspace"),
			Frame:     context.StringFlag("frame"),
		},
	}
	return
}

func orchestrateViewer(result *controllers.Result) (output string, err error) {
	r := result.Payload.(*controllers.FrameOrchestrationsCreateResult)

	// Stream orchestrate output
	io.Printf("Orchestrate %s\n", r.Frame)
	if err = result.Stream.Print(); err != nil {
		return
	}

	output, err = cliapp.View("framecompositions/create")(result)
	return
}
