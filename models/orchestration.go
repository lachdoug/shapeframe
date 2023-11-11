package models

import (
	"sf/app"
	"sf/utils"
)

type Orchestration struct {
	Frame       *Frame
	Stream      *utils.Stream
	Deployments []*Deployment
}

func NewOrchestration(f *Frame, st *utils.Stream) (o *Orchestration) {
	o = &Orchestration{
		Frame:  f,
		Stream: st,
	}
	return
}

// Action

func (o *Orchestration) apply() {
	var err error
	o.SetDeployments()
	o.write("Orchestrating %s\n", o.Frame.Name)
	o.write("Configuration:\n")
	for _, setting := range o.Frame.Configuration.Details() {
		o.write("  %s: %s\n", setting["Key"], setting["Value"])
	}

	if err = o.build(); err != nil {
		o.error(err, "apply orchestration")
		o.close()
		return
	}
	o.close()
}

func (o *Orchestration) SetDeployments() {
	for _, s := range o.Frame.Shapes {
		d := NewDeployment(s, o.Stream)
		o.Deployments = append(o.Deployments, d)
	}
}

func (o *Orchestration) build() (err error) {
	for _, d := range o.Deployments {
		if err = d.build(); err != nil {
			return
		}
	}
	return
}

// Stream

func (o *Orchestration) write(format string, a ...any) {
	o.Stream.Writef(format, a...)
}

func (o *Orchestration) error(err error, format string, a ...any) {
	err = app.ErrorWrapf(err, format, a...)
	o.Stream.Error(err)
}

func (o *Orchestration) close() {
	o.Stream.Close()
}
