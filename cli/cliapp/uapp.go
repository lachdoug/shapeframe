package cliapp

import (
	"fmt"
	"os"
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
			Name:         cc.Name,
			Usage:        cc.Summary,
			UsageText:    strings.Join(cc.Usage, "\n"),
			Description:  strings.Join(cc.Description, "\n"),
			Aliases:      cc.Aliases,
			Flags:        uFlags(cc.Flags),
			OnUsageError: usageError,
			HideHelp:     true,
			Action:       uAction(kind, cc.Parametizer, cc.Controller, cc.Viewer),
		}
	case *CommandSet:
		ucommand = &ucli.Command{
			Name:         cc.Name,
			Usage:        cc.Summary,
			UsageText:    cc.Usage,
			Description:  strings.Join(cc.Description, "\n"),
			Aliases:      cc.Aliases,
			Flags:        uFlags(cc.Flags),
			OnUsageError: usageError,
			HideHelp:     true,
			Subcommands:  uSubcommands(cc.Commands),
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
	parametizer func(*Context) ([]byte, error),
	controller func([]byte) ([]byte, error),
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
		err = cliAction.run(cliContext)
		return
	}
	return
}

func uFlags(flags []string) (uflags []ucli.Flag) {
	uflags = append(uflags, uFlag("bool", "help h ?", "Show help"))
	for i := 0; i < len(flags); i += 3 {
		uflags = append(uflags, uFlag(flags[i], flags[i+1], flags[i+2]))
	}
	return
}

func uFlag(kind string, namers string, about string) (uflag ucli.Flag) {
	names := strings.Split(namers, " ")
	name := names[0]
	aliases := names[1:]
	switch kind {
	case "bool":
		uflag = &ucli.BoolFlag{
			Name:    name,
			Aliases: aliases,
			Usage:   about,
		}
	case "string":
		uflag = &ucli.StringFlag{
			Name:    name,
			Aliases: aliases,
			Usage:   about,
		}
	}
	return
}

var commandNotFound func(*ucli.Context, string) = func(ucontext *ucli.Context, command string) {
	_, _ = fmt.Fprintf(ucontext.App.Writer, "Incorrect Usage: %q is not a valid command\n", command)
	os.Exit(1)
}

var usageError func(*ucli.Context, error, bool) error = func(ucontext *ucli.Context, err error, isSubcommand bool) error {
	_, _ = fmt.Fprintf(ucontext.App.Writer, "Incorrect Usage: %s\n", err)
	os.Exit(1)
	return nil
}
