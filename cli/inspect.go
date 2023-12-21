package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func inspect() (command any) {
	command = &cliapp.Command{
		Name:    "inspect",
		Aliases: ss("i"),
		Summary: "Inspect workspace",
		Usage: ss(
			"sf inspect [options]",
		),
		Controller: controllers.WorkspaceInspectsRead,
		Viewer: cliapp.View(
			"workspaceinspects/read",
			"workspaceinspects/frames",
			"workspaceinspects/frame",
			"workspaceinspects/frame/shape",
			"workspaceinspects/frame/shapes",
			"workspaceinspects/frame/relations",
			"workspaceinspects/frame/children",
			"workspaceinspects/frame/ancestors",
			"workspaceinspects/frame/descendents",
			"workspaceinspects/repositories",
			"workspaceinspects/repository",
			"workspaceinspects/directories",
			"workspaceinspects/directory",
			"configurations/configuration",
			"configurations/settings",
			"configurations/setting",
			"gitrepos/gitrepo",
			"gitrepos/branches",
			"gitrepos/framers",
			"gitrepos/framer",
			"gitrepos/shapers",
			"gitrepos/shaper",
		),
	}
	return
}
