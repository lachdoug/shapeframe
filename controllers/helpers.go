package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

func paramsFor[T any](jparams []byte) (params *T) {
	params = new(T)
	utils.JsonUnmarshal(jparams, &params)
	return
}

func jbodyFor(result any, opts ...any) (j []byte) {
	body := &app.Body{Result: result}
	if len(opts) > 0 {
		st := opts[0].(*models.Stream).Identifier
		body.Stream = st
	}
	j = utils.JsonMarshal(body)
	return
}
