package tuiform

import (
	"fmt"
	"math"
	"regexp"
	"sf/models"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/slices"
)

type Form struct {
	ID              string
	Title           string
	ComponentModels []*models.FormComponent
	Answers         map[string]string
	SubmitCallback  func(map[string]string) tea.Cmd
	CancelCallback  func() tea.Cmd
	Components      []Componenter
	// FocusIndex      int
	Width    int
	Validity string
	IsActive bool
	// Cancelled  bool
}

func NewForm(
	id string,
	title string,
	components []*models.FormComponent,
	answers map[string]string,
	submitCallback func(map[string]string) tea.Cmd,
	cancelCallback func() tea.Cmd,
) (fm *Form) {
	fm = &Form{
		ID:              fmt.Sprintf("%s-form", id),
		Title:           title,
		ComponentModels: components,
		Answers:         answers,
		SubmitCallback:  submitCallback,
		CancelCallback:  cancelCallback,
	}
	return
}

func (fm *Form) setAnswers() {
	if fm.Answers == nil {
		fm.Answers = map[string]string{}
	}
}

func (fm *Form) setComponents() {
	for _, fmcm := range fm.ComponentModels {
		fm.Components = append(fm.Components, fm.newControl(fm, fmcm))
	}
	fm.Components = append(fm.Components, newButtons(fm))
}

func (fm *Form) newControl(pt Parenter, fmcm *models.FormComponent) (c Componenter) {
	switch fmcm.Type {
	case "row":
		c = NewRow(fm, pt, fmcm)
	default:
		c = NewField(fm, pt, fmcm)
	}
	return
}

func (fm *Form) Init() (c tea.Cmd) {
	fm.setAnswers()
	fm.setComponents()
	for _, fmc := range fm.Components {
		fmc.Init()
	}
	fm.depend()
	fm.IsActive = true
	return
}

func (fm *Form) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = fm
	if !fm.IsActive {
		return
	}
	c = fm.updateComponents(msg)
	fm.depend()
	return
}

func (fm *Form) updateComponents(msg tea.Msg) (c tea.Cmd) {
	cs := []tea.Cmd{}
	for _, fmc := range fm.Components {
		_, c = fmc.Update(msg)
		cs = append(cs, c)
	}
	c = tea.Batch(cs...)
	return
}

func (fm *Form) View() (v string) {
	headingStyle := lipgloss.NewStyle().
		Bold(true)

	lines := []string{headingStyle.Render(fm.Title)}
	for _, fmc := range fm.Components {
		lines = append(lines, fmc.View())
	}
	lines = slices.DeleteFunc(lines, func(s string) bool { return s == "" })
	v = lipgloss.JoinVertical(lipgloss.Left, lines...)
	return
}

func (fm *Form) set(k string, v string) {
	fm.Answers[k] = v
	fm.depend()
}

func (fm *Form) width() (w int) {
	w = fm.Width
	return
}

func (fm *Form) SetSize(w int, h int) {
	x := int(math.Mod(float64(w), 12))
	fm.Width = w - x
	for _, fmc := range fm.Components {
		fmc.resize()
	}
}

// func (fm *Form) enter() (c tea.Cmd) { return nil }

func (fm *Form) FocusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{}
	for _, fmc := range fm.Components {
		fc = append(fc, fmc.FocusChain()...)
	}
	return
}

func (fm *Form) submit() (c tea.Cmd) {
	fm.IsActive = false
	c = fm.SubmitCallback(fm.Answers)
	return
}

func (fm *Form) cancel() (c tea.Cmd) {
	fm.IsActive = false
	c = fm.CancelCallback()
	return
}

func (fm *Form) depend() {
	for _, fmc := range fm.Components {
		fmc.depend()
	}
}

// func (fm *Form) value() string { return "" }

func (fm *Form) validity() (vy string) {
	vy = ""
	for _, fmc := range fm.Components {
		vy = vy + fmc.validity()
	}
	fm.Validity = vy
	return
}

func (fm *Form) isDependMatch(depend *models.FormDepend) (is bool) {
	if depend == nil {
		is = true
		return
	}
	var r *regexp.Regexp
	var err error
	if r, err = regexp.Compile(depend.Pattern); err != nil {
		panic(err)
	}
	is = r.MatchString(fm.Answers[depend.Key])
	return
}

func (fm *Form) shown() (ks []string) {
	ks = []string{}
	for _, fmc := range fm.Components {
		ks = append(ks, fmc.shown()...)
	}
	return
}

func (fm *Form) resize() {}
