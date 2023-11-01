package cli

import "sf/cli/cliapp"

func Router(args []string) {
	cli := &cliapp.Router{
		Name:    "sf",
		Summary: "Shapeframe",
		Description: ss(
			"Use Shapeframe for provisioning and deployment",
			"  model systems",
			"  create configurations and artifacts",
			"  invoke running systems",
		),
		Commands: cs(
			list,
			create,
			destroy,
			label,
			configure,
			inspect,
			add,
			remove,
			pull,
			enter,
			exit,
			context,
			hi,
			bye,
			nuke,
		),
	}
	cli.Run(args)
}
