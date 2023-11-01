package cliapp

import (
	"fmt"
	"sf/app"
	"strings"

	ucli "github.com/urfave/cli/v2"
)

func uApp(router *Router) (uapp *ucli.App) {
	uapp = &ucli.App{
		Name:                 router.Name,
		Usage:                router.Summary,
		UsageText:            router.Usage,
		Description:          strings.Join(router.Description, "\n"),
		Flags:                uFlags(router.Flags),
		EnableBashCompletion: true,
		CommandNotFound:      commandNotFound,
		OnUsageError:         usageError,
		HideHelp:             true,
		Commands:             uCommands(router.Commands),
	}
	return
}

func uCommands(commands []func() any) (ucommands []*ucli.Command) {
	for _, command := range commands {
		ucommands = append(ucommands, uCommand("command", command))
	}
	return
}

func uCommand(kind string, command func() any) (ucommand *ucli.Command) {
	switch cc := command().(type) {
	case *Command:
		ucommand = &ucli.Command{
			Name:        cc.Name,
			Usage:       cc.Summary,
			UsageText:   strings.Join(cc.Usage, "\n"),
			Description: strings.Join(cc.Description, "\n"),
			Aliases:     cc.Aliases,
			Flags:       uFlags(cc.Flags),
			HideHelp:    true,
			Action:      uAction(kind, cc.Parametizer, cc.Controller, cc.Viewer),
		}
	case *CommandSet:
		ucommand = &ucli.Command{
			Name:        cc.Name,
			Usage:       cc.Summary,
			UsageText:   cc.Usage,
			Description: strings.Join(cc.Description, "\n"),
			Aliases:     cc.Aliases,
			Flags:       uFlags(cc.Flags),
			HideHelp:    true,
			Subcommands: uSubcommands(cc.Commands),
		}
	}
	return
}

func uSubcommands(commands []func() any) (usubcommands []*ucli.Command) {
	for _, command := range commands {
		usubcommands = append(usubcommands, uCommand("subcommand", command))
	}
	return
}

func uAction(
	kind string,
	parametizer func(*Context) ([]byte, *app.Validation, error),
	controller func([]byte) ([]byte, *app.Validation, error),
	viewer func(map[string]any) (string, error),
) (ufunction func(*ucli.Context) error) {
	ufunction = func(ucontext *ucli.Context) (err error) {
		cliAction := &Action{
			Parametizer: parametizer,
			Controller:  controller,
			Viewer:      viewer,
		}
		cliContext := &Context{
			UContext: ucontext,
		}
		cliAction.run(cliContext)
		return
	}
	return
}

func uFlags(flags []string) (uflags []ucli.Flag) {
	uflags = append(uflags, uFlag("bool", "help", "Show help"))
	for i := 0; i < len(flags); i += 3 {
		uflags = append(uflags, uFlag(flags[i], flags[i+1], flags[i+2]))
	}
	return
}

func uFlag(kind string, name string, about string) (uflag ucli.Flag) {
	switch kind {
	case "bool":
		uflag = &ucli.BoolFlag{
			Name:  name,
			Usage: about,
		}
	case "string":
		uflag = &ucli.StringFlag{
			Name:  name,
			Usage: about,
		}
	}
	return
}

var commandNotFound func(*ucli.Context, string) = func(ucontext *ucli.Context, command string) {
	fmt.Fprintf(ucontext.App.Writer, "Incorrect Usage: %q is not a valid command\n", command)
}

var usageError func(*ucli.Context, error, bool) error = func(ucontext *ucli.Context, err error, isSubcommand bool) error {
	fmt.Fprintf(ucontext.App.Writer, "Incorrect Usage: %s\n", err)
	return nil
}
