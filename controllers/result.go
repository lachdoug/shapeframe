package controllers

import (
	"sf/app/streams"
	"sf/app/validations"
)

type Result struct {
	Payload    any
	Validation *validations.Validation
	Stream     *streams.Stream
}
