package tuiform

// import (
// 	"sf/app/logs"

// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

// type Submit struct {
// 	Buttons   *Buttons
// 	IsFocus   bool
// 	IsEnabled bool
// }

// func newSubmit(btns *Buttons) (sb *Submit) {
// 	logs.Log("btns", btns)
// 	sb = &Submit{Buttons: btns}
// 	return
// }

// func (sb *Submit) Init() tea.Cmd { return nil }

// func (sb *Submit) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
// 	m = sb
// 	if sb.IsFocus {
// 		switch msg := msg.(type) {
// 		case tea.KeyMsg:
// 			switch msg.Type {
// 			case tea.KeyEnter:
// 				if sb.IsEnabled {
// 					c = sb.enter()
// 				}
// 				// case tea.KeyTab:
// 				// 	c = sb.next()
// 				// 	return
// 				// case tea.KeyShiftTab:
// 				// 	c = sb.previous()
// 				// 	return
// 			}
// 		case error:
// 			panic(msg)
// 		}
// 	}
// 	return
// }

// func (sb *Submit) View() (v string) {
// 	style := lipgloss.NewStyle().
// 		BorderStyle(lipgloss.RoundedBorder()).
// 		PaddingLeft(1).PaddingRight(1)

// 	if sb.IsFocus {
// 		style.
// 			BorderForeground(lipgloss.Color("15"))
// 	} else {
// 		style.
// 			Foreground(lipgloss.Color("7")).
// 			BorderForeground(lipgloss.Color("8"))
// 	}

// 	if !sb.IsEnabled {
// 		style.Foreground(lipgloss.Color("9"))
// 	}

// 	v = style.Render("Submit")
// 	return
// }

// func (sb *Submit) width() int { return 0 }
// func (sb *Submit) resize()    {}

// func (sb *Submit) enter() (c tea.Cmd) {
// 	c = sb.Buttons.Form.submit()
// 	return
// }

// // func (sb *Submit) next() (c tea.Cmd) {
// // 	c = sb.Buttons.next()
// // 	return
// // }

// // func (sb *Submit) previous() (c tea.Cmd) {
// // 	c = sb.Buttons.previous()
// // 	return
// // }

// func (sb *Submit) Focus(aspect string) (c tea.Cmd) {
// 	if !sb.IsEnabled {
// 		if aspect == "next" {
// 			c = func() tea.Msg { return tea.KeyTab }
// 		} else {
// 			c = func() tea.Msg { return tea.KeyShiftTab }
// 		}
// 		return
// 	}
// 	sb.IsFocus = true
// 	return
// }

// func (sb *Submit) Blur() {
// 	sb.IsFocus = false
// }

// func (sb *Submit) depend() {
// 	sb.IsEnabled = sb.Buttons.Form.validity() == ""
// }

// func (sb *Submit) set(string, string) {}
// func (sb *Submit) value() string      { return "" }
// func (sb *Submit) validity() string   { return "" }

// func (sb *Submit) shown() []string { return nil }
