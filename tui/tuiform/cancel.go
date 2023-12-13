package tuiform

// import (
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

// type Cancel struct {
// 	Buttons *Buttons
// 	IsFocus bool
// }

// func newCancel(btns *Buttons) (cb *Cancel) {
// 	cb = &Cancel{Buttons: btns}
// 	return
// }

// func (cb *Cancel) Init() tea.Cmd { return nil }

// func (cb *Cancel) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
// 	m = cb
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.Type {
// 		case tea.KeyEnter:
// 			if cb.IsFocus {
// 				c = cb.enter()
// 			}
// 		}
// 	}
// 	return
// }

// func (cb *Cancel) View() (v string) {
// 	style := lipgloss.NewStyle().
// 		BorderStyle(lipgloss.RoundedBorder()).
// 		PaddingLeft(1).PaddingRight(1).
// 		Foreground(lipgloss.Color("9"))

// 	if cb.IsFocus {
// 		style.
// 			BorderForeground(lipgloss.Color("15"))
// 	} else {
// 		style.
// 			Foreground(lipgloss.Color("7")).
// 			BorderForeground(lipgloss.Color("8"))
// 	}

// 	v = style.Render("Cancel")
// 	return
// }

// func (cb *Cancel) width() int { return 0 }
// func (cb *Cancel) resize()    {}

// func (cb *Cancel) enter() (c tea.Cmd) {
// 	c = cb.Buttons.Form.cancel()
// 	return
// }

// // func (cb *Cancel) next() (c tea.Cmd) {
// // 	c = cb.Buttons.next()
// // 	return
// // }

// // func (cb *Cancel) previous() (c tea.Cmd) {
// // 	c = cb.Buttons.previous()
// // 	return
// // }

// func (cb *Cancel) Focus(string) (c tea.Cmd) {
// 	cb.IsFocus = true
// 	return
// }

// func (cb *Cancel) Blur() {
// 	cb.IsFocus = false
// }

// func (cb *Cancel) depend()            {}
// func (cb *Cancel) set(string, string) {}
// func (cb *Cancel) value() string      { return "" }
// func (cb *Cancel) validity() string   { return "" }

// func (cb *Cancel) shown() []string { return nil }
