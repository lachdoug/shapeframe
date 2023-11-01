package cliapp

type CommandSet struct {
	Name        string
	Summary     string
	Usage       string
	Description []string
	Aliases     []string
	Flags       []string
	Commands    []func() any
}
