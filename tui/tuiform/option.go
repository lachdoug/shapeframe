package tuiform

import (
	"fmt"
	"sf/models"
	"sf/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Option struct {
	Field      *Field
	Model      *models.FormOption
	IsSelected bool
	IsFocus    bool
	IsHover    bool
}

func NewOption(field *Field, model *models.FormOption) (li *Option) {
	li = &Option{
		Field: field,
		Model: model,
	}
	return
}

func (opt *Option) ID() (id string) {
	id = fmt.Sprintf("%s-option-%s", opt.Field.ID, opt.Model.Value)
	return
}

func (opt *Option) View() (v string) {
	prefix := " "
	style := lipgloss.NewStyle()

	if opt.IsHover {
		style = style.Background(lipgloss.Color("0"))
	}
	if opt.Field.IsFocus {
		if opt.IsFocus {
			prefix = "⮞"
		}
	} else {
		if !opt.IsSelected {
			style = style.Foreground(lipgloss.Color("8"))
		}
	}

	if opt.Field.Model.As == "checks" {
		if opt.IsSelected {
			prefix = fmt.Sprintf("%s %s", prefix, "🗹")
		} else {
			prefix = fmt.Sprintf("%s %s", prefix, "☐")
		}
	} else if opt.Field.Model.As == "radios" {
		if opt.IsSelected {
			prefix = fmt.Sprintf("%s %s", prefix, "◉")
		} else {
			prefix = fmt.Sprintf("%s %s", prefix, "○")
		}
	}

	line := utils.FixedLengthString(
		fmt.Sprintf("%s %s", prefix, opt.Model.Label),
		opt.Field.Width-2,
	)

	v = zone.Mark(opt.ID(), style.Render(line))
	return
}

func (opt *Option) inBounds(msg tea.MouseMsg) (is bool) {
	is = zone.Get(opt.ID()).InBounds(msg)
	return
}
