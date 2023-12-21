package tuiform

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Radios struct {
	Field   *Field
	Options []*Option
}

func NewRadios(f *Field) (rs *Radios) {
	rs = &Radios{Field: f}
	return
}

func (rs *Radios) setOptions() {
	options := []*Option{}
	for _, model := range rs.Field.Model.Options {
		options = append(options, NewOption(rs.Field, model))
	}
	rs.Options = options
}

func (rs *Radios) setSelection() {
	an := rs.Field.answer()
	isSelected := false
	for _, opt := range rs.Options {
		if an == opt.Model.Value {
			isSelected = true
			opt.IsFocus = true
			opt.IsSelected = true
		}
	}
	if !isSelected {
		opt := rs.Options[0]
		opt.IsSelected = true
		opt.IsFocus = true
	}
}

func (rs *Radios) Init() (c tea.Cmd) {
	rs.setOptions()
	rs.setSelection()
	return
}

func (rs *Radios) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = rs
	switch msg := msg.(type) {
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			for _, opt := range rs.Options {
				if rs.Field.IsFocus && opt.inBounds(msg) {
					opt.IsHover = true
				} else {
					opt.IsHover = false
				}
			}
		case tea.MouseLeft:
			if rs.Field.IsHover {
				var hover *Option
				for _, opt := range rs.Options {
					if opt.IsHover {
						hover = opt
					}
				}
				if hover != nil {
					for _, opt := range rs.Options {
						if opt == hover {
							opt.IsFocus = true
							opt.IsSelected = true
						} else {
							opt.IsFocus = false
							opt.IsSelected = false
						}
					}
				}
				rs.setAnswer()
			}
		}
	case tea.KeyMsg:
		if rs.Field.IsFocus {
			switch msg.Type {
			case tea.KeyEnter:
				for _, opt := range rs.Options {
					if opt.IsFocus {
						opt.IsSelected = true
					} else {
						opt.IsSelected = false
					}
				}
				rs.setAnswer()
				c = rs.Field.enter()
			case tea.KeyDown:
				rs.down()
				rs.setAnswer()
			case tea.KeyUp:
				rs.up()
				rs.setAnswer()
			}
		}
	}
	return
}

func (rs *Radios) View() (v string) {
	vs := []string{}
	for _, opt := range rs.Options {
		vs = append(vs, opt.View())
	}
	v = lipgloss.JoinVertical(lipgloss.Left, vs...)
	return
}

func (rs *Radios) setAnswer() {
	rs.Field.setAnswer(rs.value())
}

func (rs *Radios) focus(aspect string) (c tea.Cmd) {
	for _, opt := range rs.Options {
		if opt.IsSelected {
			opt.IsFocus = true
		} else {
			opt.IsFocus = false
		}
	}
	return
}

func (rs *Radios) blur() {
	for _, opt := range rs.Options {
		opt.IsFocus = false
	}
}

func (rs *Radios) value() (v string) {
	for _, opt := range rs.Options {
		if opt.IsSelected {
			v = opt.Model.Value
			return
		}
	}
	return
}

func (rs *Radios) up() {
	index := 0
	for i, opt := range rs.Options {
		if opt.IsSelected {
			index = i
		}
		opt.IsFocus = false
		opt.IsSelected = false
	}
	index = index - 1
	if index == -1 {
		index = 0
	}
	opt := rs.Options[index]
	opt.IsFocus = true
	opt.IsSelected = true
}

func (rs *Radios) down() {
	index := 0
	for i, opt := range rs.Options {
		if opt.IsSelected {
			index = i
		}
		opt.IsFocus = false
		opt.IsSelected = false
	}
	index = index + 1
	if index == len(rs.Options) {
		index = index - 1
	}
	opt := rs.Options[index]
	opt.IsFocus = true
	opt.IsSelected = true
}
