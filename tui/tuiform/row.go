package tuiform

import (
	"sf/models"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Row struct {
	Form           *Form
	ComponentModel *models.FormComponent
	Components     []Componenter
	IsVisible      bool
}

func NewRow(fm *Form, fmcm *models.FormComponent) (r *Row) {
	r = &Row{Form: fm, ComponentModel: fmcm}
	return
}

func (r *Row) setWidth(w int) {
	for _, fmc := range r.Components {
		fmc.setWidth(w)
	}
}

func (r *Row) setComponents() {
	for _, model := range r.ComponentModel.Cols {
		r.Components = append(r.Components, r.Form.newControl(model))
	}
}

func (r *Row) Init() (c tea.Cmd) {
	r.setComponents()
	for _, fmc := range r.Components {
		fmc.Init()
	}
	return
}

func (r *Row) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = r
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

func (r *Row) FocusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{}
	for _, fmc := range r.Components {
		fc = append(fc, fmc.FocusChain()...)
	}
	return
}

func (r *Row) depend() {
	r.IsVisible = r.Form.isDependMatch(r.ComponentModel.Depend)
	if r.IsVisible {
		for _, fmc := range r.Components {
			fmc.depend()
		}
	}
}

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

func (r *Row) isFocus() (is bool) {
	for _, fmc := range r.Components {
		if fmc.isFocus() {
			is = true
			return
		}
	}
	return
}
