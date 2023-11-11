package controllers

import (
	"sf/app"
	"sf/utils"
)

func ParamsFor[T any](jparams []byte) (params *T) {
	params = new(T)
	utils.JsonUnmarshal(jparams, &params)
	return
}

func jbodyFor(result any, opts ...any) (j []byte) {
	body := &app.Body{Result: result}

	if len(opts) > 0 && opts[0] != nil {
		vmap := opts[0].(*app.Validation).Map()
		body.Invalid = vmap
	}
	if len(opts) > 1 {
		st := opts[1].(*utils.Stream).Identifier
		body.Stream = st
	}
	j = utils.JsonMarshal(body)
	return
}
