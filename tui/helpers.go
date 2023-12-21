package tui

import (
	"fmt"
	"sf/models"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func configurationInfo(title string, configuration *models.Configuration) (v string) {
	if len(configuration.Info) == 0 {
		v = fmt.Sprintf("%s: <none>", title)
		return
	} else if len(configuration.Settings) == 0 {
		v = fmt.Sprintf("%s: <unset>", title)
		return
	}
	vs := []string{}
	for _, setting := range configuration.Info {
		vs = append(vs, configurationSettingInfo(setting))
	}
	v = lipgloss.JoinVertical(lipgloss.Left,
		title,
		lipgloss.JoinVertical(lipgloss.Left, vs...),
	)
	return
}

func configurationSettingInfo(setting map[string]any) (v string) {
	indentStyle := lipgloss.NewStyle().PaddingLeft(2)
	ctype := setting["Type"].(string)
	label := setting["Label"].(string)
	value := setting["Value"].(string)
	options := setting["Options"].(map[string]string)
	vs := []string{}
	if strings.Contains(value, "\n") {
		vs = append(vs, fmt.Sprintf("%s:", label))
		vals := strings.Split(value, "\n")
		for _, val := range vals {
			var display string
			if ctype == "select" || ctype == "selects" {
				display = options[val]
			} else {
				display = val
			}
			vs = append(vs, indentStyle.Render(display))
		}
	} else {
		var display string
		if ctype == "select" || ctype == "selects" {
			display = options[value]
		} else {
			display = value
		}
		vs = append(vs, fmt.Sprintf("%s: %s", label, display))
	}
	v = indentStyle.Render(lipgloss.JoinVertical(lipgloss.Left, vs...))
	return
}
