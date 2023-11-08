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
			list,
			label,
			configure,
			inspect,
			add,
			remove,
			pull,
			enter,
			exit,
			context,
			orchestrate,
			nuke,
		),
	}
	cli.Run(args)
}
