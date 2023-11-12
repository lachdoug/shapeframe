package cli

import "sf/cli/cliapp"

func Router(args []string) {
	cli := &cliapp.Router{
		Name:    "sf",
		Summary: "Shapeframe",
		Description: ss(
			"Use Shapeframe for provisioning and deployment",
			"  Model systems",
			"  Create configurations and artifacts",
			"  Invoke running systems",
		),
		Commands: cs(
			get,
			list,
			label,
			configure,
			inspect,
			add,
			remove,
			pull,
			// checkout,
			join,
			context,
			orchestrate,
			nuke,
		),
	}
	cli.Run(args)
}
