package tui

import (
	"sf/tui/tuisupport"

	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TopNav struct {
	App        *App
	FlexBox    *stickers.FlexBox
	LeftLinks  []*Link
	RightLinks []*Link
	Width      int
	// IsFocus    bool
	// FocusIndex int
}

func newTopNav(app *App) (tn *TopNav) {
	tn = &TopNav{App: app}
	return
}

func (tn *TopNav) Init() (c tea.Cmd) {
	tn.setLinks()
	for _, lk := range tn.allLinks() {
		lk.Init()
	}
	return
}

func (tn *TopNav) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = tn
	cs := []tea.Cmd{}
	for _, lk := range tn.allLinks() {
		_, c = lk.Update(msg)
		cs = append(cs, c)
	}
	c = tea.Batch(cs...)
	return
}

func (tn *TopNav) View() (v string) {
	tn.setFlexBox()
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Width(tn.Width).Height(1)

	if tn.isFocus() {
		style.
			BorderForeground(lipgloss.Color("15"))
	} else {
		style.
			Foreground(lipgloss.Color("7")).
			BorderForeground(lipgloss.Color("8"))
	}

	v = style.Render(tn.FlexBox.Render())
	return
}

func (tn *TopNav) setSize(w int, h int) {
	tn.Width = w - 2
}

func (tn *TopNav) setLinks() {
	tn.LeftLinks = []*Link{
		newLink(tn.App, "graph", "Graph", "/graph"),
		newLink(tn.App, "workspaces", "Workspaces", "/workspaces"),
		newLink(tn.App, "frames", "Frames", "/frames"),
		newLink(tn.App, "shapes", "Shapes", "/shapes"),
	}
	if is, _ := tn.App.matchRoute("/workspaces"); is {
		tn.RightLinks = []*Link{
			newLink(tn.App, "workspaces-new", "New", "new"),
		}
	} else if is, _ := tn.App.matchRoute("/frames"); is {
		tn.RightLinks = []*Link{
			newLink(tn.App, "frames-new", "New", "new"),
		}
	} else if is, _ := tn.App.matchRoute("/shapes"); is {
		tn.RightLinks = []*Link{
			newLink(tn.App, "shapes-new", "New", "new"),
		}
	} else if is, _ := tn.App.matchRoute("/workspaces/@[^/]+"); is {
		tn.RightLinks = []*Link{
			newLink(tn.App, "workspaces-inspect", "Inspect", "inspect"),
			newLink(tn.App, "workspaces-delete", "Delete", "delete"),
		}
	} else if is, _ := tn.App.matchRoute("/frames/@[^/]+"); is {
		tn.RightLinks = []*Link{
			newLink(tn.App, "frames-orchestrate", "Orchestrate", "orchestrate"),
			newLink(tn.App, "frames-delete", "Delete", "delete"),
		}
	} else if is, _ := tn.App.matchRoute("/shapes/@[^/]+"); is {
		tn.RightLinks = []*Link{
			newLink(tn.App, "shapes-delete", "Delete", "delete"),
		}
	} else {
		tn.RightLinks = nil
	}
}

func (tn *TopNav) setFlexBox() {
	tn.FlexBox = stickers.NewFlexBox(tn.Width, 1)
	tn.FlexBox.AddRows([]*stickers.FlexBoxRow{
		tn.FlexBox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(1, 1).SetContent(tn.renderLeftLinks()),
				stickers.NewFlexBoxCell(1, 1).SetContent(tn.renderRightLinks()),
			},
		),
	})
}

func (tn *TopNav) renderLeftLinks() (v string) {
	vs := []string{}
	for _, lk := range tn.LeftLinks {
		vs = append(vs, lk.View())
	}
	v = lipgloss.JoinHorizontal(lipgloss.Top, vs...)
	return
}

func (tn *TopNav) renderRightLinks() (v string) {
	style := lipgloss.NewStyle().Width(tn.Width / 2).Align(lipgloss.Right)
	vs := []string{}
	for _, lk := range tn.RightLinks {
		vs = append(vs, lk.View())
	}
	v = style.Render(lipgloss.JoinHorizontal(lipgloss.Top, vs...))
	return
}

func (tn *TopNav) allLinks() (lks []*Link) {
	lks = append(lks, tn.LeftLinks...)
	lks = append(lks, tn.RightLinks...)
	return
}

func (tn *TopNav) isFocus() (is bool) {
	for _, lk := range tn.allLinks() {
		if lk.IsFocus {
			is = true
			return
		}
	}
	return
}

func (tn *TopNav) focusChain() (fc []tuisupport.Focuser) {
	for _, lk := range tn.allLinks() {
		fc = append(fc, lk)
	}
	return
}

// func (tn *TopNav) focus(on ...string) (c tea.Cmd) {
// 	tn.IsFocus = true
// 	if slices.Contains(on, "first") {
// 		tn.FocusIndex = 0
// 	} else if slices.Contains(on, "last") {
// 		tn.FocusIndex = len(tn.LeftLinks) + len(tn.RightLinks) - 1
// 	}
// 	if tn.FocusIndex < 4 {
// 		c = tn.LeftLinks[tn.FocusIndex].focus()
// 	} else {
// 		c = tn.RightLinks[tn.FocusIndex-4].focus()
// 	}
// 	return
// }

// func (tn *TopNav) blur() {
// 	tn.IsFocus = false
// 	// for _, lk := range tn.LeftLinks {
// 	// 	lk.blur()
// 	// }
// 	// for _, lk := range tn.RightLinks {
// 	// 	lk.blur()
// 	// }
// 	if tn.FocusIndex < 4 {
// 		tn.LeftLinks[tn.FocusIndex].blur()
// 	} else {
// 		tn.RightLinks[tn.FocusIndex-3].blur()
// 	}
// }

// func (tn *TopNav) next() (c tea.Cmd) {
// 	// tn.blur()
// 	tn.FocusIndex++
// 	if tn.FocusIndex == len(tn.LeftLinks)+len(tn.RightLinks) {
// 		tn.FocusIndex = len(tn.LeftLinks) + len(tn.RightLinks) - 1
// 		c = tn.App.next()
// 	} else {
// 		c = tn.focus("next")
// 	}
// 	return
// }

// func (tn *TopNav) previous() (c tea.Cmd) {
// 	// tn.blur()
// 	tn.FocusIndex--
// 	if tn.FocusIndex == -1 {
// 		tn.FocusIndex = 0
// 		c = tn.App.previous()
// 	} else {
// 		c = tn.focus("previous")
// 	}
// 	return
// }
