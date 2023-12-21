package tuiform

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/slices"
)

type Selects struct {
	Field   *Field
	Options []*Option
}

func NewSelects(f *Field) (ss *Selects) {
	ss = &Selects{Field: f}
	return
}

func (ss *Selects) setOptions() {
	items := []*Option{}
	for _, model := range ss.Field.Model.Options {
		items = append(items, NewOption(ss.Field, model))
	}
	ss.Options = items
}

func (ss *Selects) setSelections() {
	ans := strings.Split(ss.Field.answer(), "\n")
	for _, opt := range ss.Options {
		if slices.Contains(ans, opt.Model.Value) {
			opt.IsSelected = true
		}
	}
}

func (ss *Selects) Init() (c tea.Cmd) {
	ss.setOptions()
	ss.setSelections()
	return
}

func (ss *Selects) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = ss
	switch msg := msg.(type) {
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			for _, opt := range ss.Options {
				if ss.Field.IsFocus && opt.inBounds(msg) {
					opt.IsHover = true
				} else {
					opt.IsHover = false
				}
			}
		case tea.MouseLeft:
			if ss.Field.IsHover {
				for _, opt := range ss.Options {
					if opt.IsHover {
						opt.IsHover = false
						if opt.IsSelected {
							opt.IsSelected = false
						} else {
							opt.IsSelected = true
						}
					}
				}
				ss.setAnswer()
			}
		}
	case tea.KeyMsg:
		if ss.Field.IsFocus {
			switch msg.Type {
			case tea.KeyEnter:
				c = ss.Field.enter()
				return
			case tea.KeySpace:
				for _, opt := range ss.Options {
					if opt.IsFocus {
						if opt.IsSelected {
							opt.IsSelected = false
						} else {
							opt.IsSelected = true
						}
					}
					opt.IsHover = false
				}
				ss.setAnswer()
			case tea.KeyDown:
				ss.down()
			case tea.KeyUp:
				ss.up()
			}
		}
	}
	return
}

func (ss *Selects) View() (v string) {
	vs := []string{}
	if ss.Field.IsFocus {
		for _, opt := range ss.selectedOptions() {
			vs = append(vs, opt.View())
		}
		vs = append(vs, strings.Repeat("─", ss.Field.Width-2))
		for _, opt := range ss.unselectedOptions() {
			vs = append(vs, opt.View())
		}
	} else {
		for _, opt := range ss.Options {
			if opt.IsSelected {
				vs = append(vs, opt.Model.Label)
			}
		}
	}
	v = lipgloss.JoinVertical(lipgloss.Left, vs...)
	return
}

func (ss *Selects) selectedOptions() (selected []*Option) {
	selected = []*Option{}
	for _, opt := range ss.Options {
		if opt.IsSelected {
			selected = append(selected, opt)
		}
	}
	return
}

func (ss *Selects) unselectedOptions() (unselected []*Option) {
	unselected = []*Option{}
	for _, opt := range ss.Options {
		if !opt.IsSelected {
			unselected = append(unselected, opt)
		}
	}
	return
}

func (ss *Selects) allOptions() (all []*Option) {
	all = []*Option{}
	all = append(all, ss.selectedOptions()...)
	all = append(all, ss.unselectedOptions()...)
	return
}

func (ss *Selects) setAnswer() {
	ss.Field.setAnswer(ss.value())
}

func (ss *Selects) focus(aspect string) (c tea.Cmd) {
	options := ss.allOptions()
	if aspect == "next" {
		options[0].IsFocus = true
	} else {
		options[len(options)-1].IsFocus = true
	}
	return
}

func (ss *Selects) blur() {
	for _, opt := range ss.allOptions() {
		opt.IsFocus = false
	}
}

func (ss *Selects) value() (v string) {
	values := []string{}
	for _, opt := range ss.selectedOptions() {
		values = append(values, opt.Model.Value)
	}
	v = strings.Join(values, "\n")
	return
}

func (ss *Selects) up() {
	options := ss.allOptions()
	index := 0
	for i, opt := range options {
		if opt.IsFocus {
			index = i
		}
		opt.IsFocus = false
	}
	index = index - 1
	if index == -1 {
		index = 0
	}
	opt := options[index]
	opt.IsFocus = true
}

func (ss *Selects) down() {
	options := ss.allOptions()
	index := 0
	for i, opt := range options {
		if opt.IsFocus {
			index = i
		}
		opt.IsFocus = false
	}
	index = index + 1
	if index == len(options) {
		index = index - 1
	}
	opt := options[index]
	opt.IsFocus = true
}
