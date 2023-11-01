package models

type Bye struct {
	Tone string
	Name string
}

func NewBye(name string, tone string) (bye *Bye) {
	bye = &Bye{
		Tone: tone,
		Name: name,
	}
	return
}
