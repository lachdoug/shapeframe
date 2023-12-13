package tuiform

import (
	"sf/models"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Row struct {
	Form           *Form
	Parent         Parenter
	ComponentModel *models.FormComponent
	Components     []Componenter
	Width          int
	// IsFocus        bool
	// FocusIndex     int
	IsVisible bool
}

func NewRow(fm *Form, pt Parenter, fmcm *models.FormComponent) (r *Row) {
	r = &Row{Form: fm, Parent: pt, ComponentModel: fmcm}
	return
}

func (r *Row) setWidth() {
	r.Width = r.Parent.width()
}

func (r *Row) setComponents() {
	for _, fmc := range r.ComponentModel.Cols {
		r.Components = append(r.Components, r.Form.newControl(r, fmc))
	}
}

func (r *Row) Init() (c tea.Cmd) {
	r.setWidth()
	r.setComponents()
	for _, fmc := range r.Components {
		fmc.Init()
	}
	return
}

func (r *Row) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = r
	// switch msg := msg.(type) {
	// case tea.KeyMsg:
	// 	switch msg.Type {
	// 	case tea.KeyEnter, tea.KeyTab, tea.KeyShiftTab:
	// 		_, c = r.Components[r.FocusIndex].Update(msg)
	// 		return
	// 	}
	// }
	c = r.updateComponents(msg)
	return
}

func (r *Row) updateComponents(msg tea.Msg) (c tea.Cmd) {
	cs := []tea.Cmd{}
	for _, fmc := range r.Components {
		_, c = fmc.Update(msg)
		cs = append(cs, c)
	}
	c = tea.Batch(cs...)
	return
}

func (r *Row) View() (v string) {
	if r.IsVisible {
		lines := []string{}
		for _, fmc := range r.Components {
			lines = append(lines, fmc.View())
		}
		v = lipgloss.JoinHorizontal(lipgloss.Top, lines...)
	}
	return
}

func (r *Row) width() (w int) {
	w = r.Width
	return
}

func (r *Row) resize() {
	r.setWidth()
	for _, fmc := range r.Components {
		fmc.resize()
	}
}

// func (r *Row) enter() (c tea.Cmd) { return nil }

// func (r *Row) next() (c tea.Cmd) {
// 	r.Blur()
// 	r.FocusIndex++
// 	if r.FocusIndex == len(r.Components) {
// 		r.FocusIndex = len(r.Components) - 1
// 		c = r.Parent.next()
// 	} else {
// 		c = r.Focus("next")
// 	}
// 	return
// }

// func (r *Row) previous() (c tea.Cmd) {
// 	r.Blur()
// 	r.FocusIndex--
// 	if r.FocusIndex < 0 {
// 		r.FocusIndex = 0
// 		c = r.Parent.previous()
// 	} else {
// 		c = r.Focus("previous")
// 	}
// 	return
// }

func (r *Row) FocusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{}
	for _, fmc := range r.Components {
		fc = append(fc, fmc.FocusChain()...)
	}
	return
}

// func (r *Row) Focus(aspect string) (c tea.Cmd) {
// 	if !r.IsVisible {
// 		if aspect == "next" {
// 			c = func() tea.Msg {return tea.KeyTab}
// 		} else {
// 			c = func() tea.Msg {return tea.KeyShiftTab}
// 		}
// 		return
// 	}
// 	if slices.Contains(on, "first") {
// 		r.FocusIndex = 0
// 	} else if slices.Contains(on, "last") {
// 		r.FocusIndex = len(r.Components) - 1
// 	}
// 	r.IsFocus = true
// 	c = r.Components[r.FocusIndex].Focus(on...)
// 	return
// }

// func (r *Row) Blur() {
// 	r.Components[r.FocusIndex].Blur()
// 	r.IsFocus = false
// }

func (r *Row) depend() {
	r.IsVisible = r.Form.isDependMatch(r.ComponentModel.Depend)
	if r.IsVisible {
		for _, fmc := range r.Components {
			fmc.depend()
		}
	}
}

func (r *Row) set(key string, value string) {
	r.Parent.set(key, value)
}

// func (r *Row) value() (vy string) { return "" }

func (r *Row) validity() (vy string) {
	vy = ""
	for _, fmc := range r.Components {
		vy = vy + fmc.validity()
	}
	return
}

func (r *Row) shown() (ks []string) {
	if r.IsVisible {
		ks = []string{}
		for _, fmc := range r.Components {
			ks = append(ks, fmc.shown()...)
		}
	}
	return
}
