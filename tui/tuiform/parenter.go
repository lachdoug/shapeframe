package tuiform

type Parenter interface {
	// Init() tea.Cmd
	// Update(tea.Msg) (tea.Model, tea.Cmd)
	// View() string
	// // Focus(...string) tea.Cmd
	// // Blur()
	// FocusChain() []tuisupport.Focuser
	width() int
	// resize()
	// enter() tea.Cmd
	// // next() tea.Cmd
	// // previous() tea.Cmd
	// depend()
	set(string, string)
	// // value() string
	// validity() string
	// shown() []string
}
