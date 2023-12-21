package cli

import "sf/cli/cliapp"

func sf() (sf *cliapp.Router) {
	sf = &cliapp.Router{
		Name:    "sf",
		Summary: "Shapeframe",
		Flags: ss(
			"bool", "debug", "Execute command in debug mode",
			"string", "directory", "Target workspace directory for command execution",
		),
		Description: ss(
			"Use Shapeframe for provisioning and deployment",
			"  Model systems",
			"  Create configurations and artifacts",
			"  Invoke running systems",
		),
		Commands: cs(
			initialize,
			get,
			list,
			label,
			configure,
			inspect,
			add,
			remove,
			pull,
			checkout,
			// join,
			context,
			orchestrate,
			terminalUserInterface,
		),
	}
	return
}
