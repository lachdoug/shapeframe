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
	LeftLinks  []*tuisupport.NavLink
	RightLinks []*tuisupport.NavLink
	Width      int
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

func (tn *TopNav) setSize(w int) {
	tn.Width = w - 2
}

func (tn *TopNav) setLinks() {
	tn.LeftLinks = []*tuisupport.NavLink{
		tuisupport.NewNavLink(tn.App, "frames", "Frames", "/frames"),
		tuisupport.NewNavLink(tn.App, "shapes", "Shapes", "/shapes"),
	}
	if is, _ := tn.App.MatchRoute("/"); is {
		tn.RightLinks = []*tuisupport.NavLink{
			tuisupport.NewNavLink(tn.App, "workspaces-inspect", "Inspect", "/inspect"),
		}
	} else if is, _ := tn.App.MatchRoute("/frames"); is {
		tn.RightLinks = []*tuisupport.NavLink{
			tuisupport.NewNavLink(tn.App, "frames-new", "New", "/frames/new"),
		}
	} else if is, _ := tn.App.MatchRoute("/shapes"); is {
		tn.RightLinks = []*tuisupport.NavLink{
			tuisupport.NewNavLink(tn.App, "shapes-new", "New", "/shapes/new"),
		}
	} else if is, _ := tn.App.MatchRoute("/frames/@[^/]+"); is {
		tn.RightLinks = []*tuisupport.NavLink{
			tuisupport.NewNavLink(tn.App, "frames-orchestrate", "Orchestrate", "orchestrate"),
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

func (tn *TopNav) allLinks() (lks []*tuisupport.NavLink) {
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
