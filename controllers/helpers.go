package controllers

// import (
// 	"sf/app/streams"
// 	"sf/app/validations"
// 	"sf/utils"
// )

// func ParamsFor[T any](jparams []byte) (params *T) {
// 	params = new(T)
// 	utils.JsonUnmarshal(jparams, &params)
// 	return
// }

// func resultFor(payload any, opts ...any) (j []byte) {
// 	result := &Result{Payload: payload}
// 	if len(opts) > 0 && opts[0] != nil {
// 		result.Validation = opts[0].(*validations.Validation)
// 	}
// 	if len(opts) > 1 {
// 		result.Stream = opts[1].(*streams.Stream)
// 	}
// 	return
// }

// // func bodyFor(payload any, opts ...any) (b *app.Body) {
// // 	body := &Result{Payload: payload}
// // 	if len(opts) > 0 && opts[0] != nil {
// // 		body.Invalid = opts[0].(*validations.Validation).Maps()
// // 	}
// // 	if len(opts) > 1 {
// // 		st := opts[1].(*streams.Stream).Identifier
// // 		body.Stream = st
// // 	}
// // 	return
// // }
