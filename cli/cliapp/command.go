package cliapp

type Command struct {
	Name        string
	Summary     string
	Usage       []string
	Description []string
	Aliases     []string
	Flags       []string
	Parametizer func(*Context) ([]byte, error)
	Controller  func([]byte) ([]byte, error)
	Viewer      func(map[string]any) (string, error)
}
