package tuisupport

type Apper interface {
	MatchRoute(string) (bool, []string)
}
