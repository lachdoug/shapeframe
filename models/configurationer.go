package models

type Configurationer interface {
	configurationFormSchema() *FormSchema
	readID() uint
	readType() string
}
