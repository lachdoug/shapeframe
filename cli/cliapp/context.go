package cliapp

import (
	ucli "github.com/urfave/cli/v2"
)

type Context struct {
	UContext *ucli.Context
}

func (context *Context) Argument(i int) (value string) {
	value = context.UContext.Args().Get(i)
	return
}

func (context *Context) Arguments() (values []string) {
	values = context.UContext.Args().Slice()
	return
}

func (context *Context) BoolFlag(name string) (value bool) {
	value = context.UContext.Bool(name)
	return
}

func (context *Context) StringFlag(name string) (value string) {
	value = context.UContext.String(name)
	return
}

func (context *Context) IsSet(name string) (is bool) {
	is = context.UContext.IsSet(name)
	return
}
