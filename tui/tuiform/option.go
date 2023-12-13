package tuiform

import (
	"fmt"
	"sf/models"

	"github.com/charmbracelet/lipgloss"
)

type Option struct {
	Value string
	Label string
}

func NewOption(fmpo *models.FormOption) (li *Option) {
	li = &Option{
		Value: fmpo.Value,
		Label: fmpo.Label,
	}
	return
}

func (li *Option) View(isFocussed bool, prefix string) (v string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("7"))
	if isFocussed {
		style.Foreground(lipgloss.AdaptiveColor{Light: "0", Dark: "15"})
	}

	textstyle := style.Copy()
	if isFocussed {
		textstyle.Bold(true)
	}

	display := li.Label
	if display == "" {
		display = li.Value
	}
	v = textstyle.Render(display)
	if prefix != "" {
		v = fmt.Sprintf("%s  %s", style.Render(prefix), v)
	}
	return
}
