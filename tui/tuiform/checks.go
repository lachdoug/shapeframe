package tuiform

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/slices"
)

type Checks struct {
	Field   *Field
	Options []*Option
}

func NewChecks(f *Field) (cs *Checks) {
	cs = &Checks{Field: f}
	return
}

func (cs *Checks) setOptions() {
	items := []*Option{}
	for _, model := range cs.Field.Model.Options {
		items = append(items, NewOption(cs.Field, model))
	}
	cs.Options = items
}

func (cs *Checks) setSelections() {
	ans := strings.Split(cs.Field.answer(), "\n")
	for _, opt := range cs.Options {
		if slices.Contains(ans, opt.Model.Value) {
			opt.IsSelected = true
		}
	}
}

func (cs *Checks) Init() (c tea.Cmd) {
	cs.setOptions()
	cs.setSelections()
	return
}

func (cs *Checks) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = cs
	switch msg := msg.(type) {
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			for _, opt := range cs.Options {
				if cs.Field.IsFocus && opt.inBounds(msg) {
					opt.IsHover = true
				} else {
					opt.IsHover = false
				}
			}
		case tea.MouseLeft:
			if cs.Field.IsHover {
				for _, opt := range cs.Options {
					if opt.IsHover {
						if opt.IsSelected {
							opt.IsSelected = false
						} else {
							opt.IsSelected = true
						}
					}
				}
				cs.setAnswer()
			}
		}
	case tea.KeyMsg:
		if cs.Field.IsFocus {
			switch msg.Type {
			case tea.KeyEnter:
				c = cs.Field.enter()
			case tea.KeySpace:
				for _, opt := range cs.Options {
					if opt.IsFocus {
						if opt.IsSelected {
							opt.IsSelected = false
						} else {
							opt.IsSelected = true
						}
					}
				}
				cs.setAnswer()
			case tea.KeyDown:
				cs.down()
			case tea.KeyUp:
				cs.up()
			}
		}
	}
	return
}

func (cs *Checks) View() (v string) {
	vs := []string{}
	for _, opt := range cs.Options {
		vs = append(vs, opt.View())
	}
	v = lipgloss.JoinVertical(lipgloss.Left, vs...)
	return
}

func (cs *Checks) setAnswer() {
	cs.Field.setAnswer(cs.value())
}

func (cs *Checks) focus(aspect string) (c tea.Cmd) {
	if aspect == "next" {
		cs.Options[0].IsFocus = true
	} else {
		cs.Options[len(cs.Options)-1].IsFocus = true
	}
	return
}

func (cs *Checks) blur() {
	for _, opt := range cs.Options {
		opt.IsFocus = false
	}
}

func (cs *Checks) value() (v string) {
	values := []string{}
	for _, opt := range cs.Options {
		if opt.IsSelected {
			values = append(values, opt.Model.Value)
		}
	}
	v = strings.Join(values, "\n")
	return
}

func (cs *Checks) up() {
	index := 0
	for i, opt := range cs.Options {
		if opt.IsFocus {
			index = i
		}
		opt.IsFocus = false
	}
	index = index - 1
	if index == -1 {
		index = 0
	}
	opt := cs.Options[index]
	opt.IsFocus = true
}

func (cs *Checks) down() {
	index := 0
	for i, opt := range cs.Options {
		if opt.IsFocus {
			index = i
		}
		opt.IsFocus = false
	}
	index = index + 1
	if index == len(cs.Options) {
		index = index - 1
	}
	opt := cs.Options[index]
	opt.IsFocus = true
}
