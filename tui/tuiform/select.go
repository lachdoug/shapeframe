package tuiform

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Select struct {
	Field   *Field
	Options []*Option
}

func NewSelect(f *Field) (s *Select) {
	s = &Select{Field: f}
	return
}

func (s *Select) setOptions() {
	options := []*Option{}
	for _, model := range s.Field.Model.Options {
		options = append(options, NewOption(s.Field, model))
	}
	s.Options = options
}

func (s *Select) setSelection() {
	an := s.Field.answer()
	isSelected := false
	for _, opt := range s.Options {
		if an == opt.Model.Value {
			isSelected = true
			opt.IsFocus = true
			opt.IsSelected = true
		}
	}
	if !isSelected {
		opt := s.Options[0]
		opt.IsFocus = true
		opt.IsSelected = true
	}
}

func (s *Select) Init() (c tea.Cmd) {
	s.setOptions()
	s.setSelection()
	return
}

func (s *Select) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = s
	switch msg := msg.(type) {
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			for _, opt := range s.Options {
				if s.Field.IsFocus && opt.inBounds(msg) {
					opt.IsHover = true
				} else {
					opt.IsHover = false
				}
			}
		case tea.MouseLeft:
			if s.Field.IsHover {
				for _, opt := range s.Options {
					if opt.IsHover {
						opt.IsFocus = true
						opt.IsSelected = true
					} else {
						opt.IsFocus = false
						opt.IsSelected = false
					}
				}
				s.setAnswer()
			}
		}
	case tea.KeyMsg:
		if s.Field.IsFocus {
			switch msg.Type {
			case tea.KeyEnter:
				for _, opt := range s.Options {
					if opt.IsFocus {
						opt.IsSelected = true
					} else {
						opt.IsSelected = false
					}
				}
				s.setAnswer()
				c = s.Field.enter()
			case tea.KeyDown:
				s.down()
				s.setAnswer()
			case tea.KeyUp:
				s.up()
				s.setAnswer()
			}
		}
	}
	return
}

func (s *Select) View() (v string) {
	vs := []string{}
	if s.Field.IsFocus {
		for _, opt := range s.Options {
			vs = append(vs, opt.View())
		}
	} else {
		for _, opt := range s.Options {
			if opt.IsSelected {
				vs = append(vs, opt.Model.Label)
			}
		}

	}
	v = lipgloss.JoinVertical(lipgloss.Left, vs...)
	return
}

func (s *Select) setAnswer() {
	s.Field.setAnswer(s.value())
}

func (s *Select) focus(aspect string) (c tea.Cmd) {
	for _, opt := range s.Options {
		if opt.IsSelected {
			opt.IsFocus = true
		} else {
			opt.IsFocus = false
		}
	}
	return
}

func (s *Select) blur() {
	for _, opt := range s.Options {
		opt.IsFocus = false
	}
}

func (s *Select) value() (v string) {
	for _, opt := range s.Options {
		if opt.IsSelected {
			v = opt.Model.Value
			return
		}
	}
	return
}

func (s *Select) up() {
	index := 0
	for i, opt := range s.Options {
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
	opt := s.Options[index]
	opt.IsFocus = true
	opt.IsSelected = true
}

func (s *Select) down() {
	index := 0
	for i, opt := range s.Options {
		if opt.IsSelected {
			index = i
		}
		opt.IsFocus = false
		opt.IsSelected = false
	}
	index = index + 1
	if index == len(s.Options) {
		index = index - 1
	}
	opt := s.Options[index]
	opt.IsFocus = true
	opt.IsSelected = true
}
