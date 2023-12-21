package tui

import (
	"sf/tui/tuisupport"

	"github.com/76creates/stickers"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Body struct {
	App       *App
	Component Componenter
	Viewport  viewport.Model
	Footer    *stickers.FlexBox
	Error     *Error
	Width     int
	Height    int
	Scrollers []*tuisupport.Scroller
}

func newBody(app *App) (b *Body) {
	b = &Body{App: app}
	return
}

func (b *Body) Init() (c tea.Cmd) {
	c = tea.Batch(
		b.setViewport(),
		b.setScrollers(),
		b.setComponent(),
	)
	return
}

func (b *Body) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = b
	b.Viewport, c = b.Viewport.Update(msg)
	cs := []tea.Cmd{c}
	if b.Error == nil {
		_, c = b.Component.Update(msg)
		cs = append(cs, c)
	} else {
		_, c = b.Error.Update(msg)
		cs = append(cs, c)
	}
	for _, s := range b.Scrollers {
		_, c = s.Update(msg)
		cs = append(cs, c)
	}
	c = tea.Batch(cs...)
	return
}

func (b *Body) View() (v string) {
	if b.Error == nil {
		v = b.Component.View()
	} else {
		v = b.Error.View()
	}
	b.Viewport.SetContent(v)
	b.setFooter()
	v = lipgloss.JoinVertical(lipgloss.Left,
		b.Viewport.View(),
		b.Footer.Render(),
	)
	return
}

func (b *Body) setSize(w int, h int) {
	b.Width = w
	b.Height = h
	b.setViewportSize()
	if b.Error == nil {
		b.Component.setSize(w, h)
	}
}

func (b *Body) setViewport() (c tea.Cmd) {
	b.Viewport = viewport.New(b.Width, b.Height)
	b.Viewport.KeyMap = viewport.KeyMap{
		PageDown: key.NewBinding(key.WithKeys("ctrl+pgdown")),
		PageUp:   key.NewBinding(key.WithKeys("ctrl+pgup")),
		Up:       key.NewBinding(key.WithKeys("ctrl+up")),
		Down:     key.NewBinding(key.WithKeys("ctrl+down")),
	}
	c = b.Viewport.Init()
	return
}

func (b *Body) setViewportSize() {
	b.Viewport.Width = b.Width
	b.Viewport.Height = b.Height - 1
}

func (b *Body) setComponent() (c tea.Cmd) {
	if is, _ := b.App.MatchRoute("/edit"); is {
		b.Component = newWorkspacesEdit(b)
	} else if is, _ := b.App.MatchRoute("/inspect"); is {
		b.Component = newWorkspacesInspect(b)
	} else if is, _ := b.App.MatchRoute("/frames"); is {
		b.Component = newFramesIndex(b)
	} else if is, _ := b.App.MatchRoute("/shapes"); is {
		b.Component = newShapesIndex(b)
	} else if is, _ := b.App.MatchRoute("/frames/new"); is {
		b.Component = newFramesNew(b)
	} else if is, _ := b.App.MatchRoute("/shapes/new"); is {
		b.Component = newShapesNew(b)
	} else if is, params := b.App.MatchRoute("/frames/@([^/.]+)"); is {
		b.Component = newFramesShow(b, params[0])
	} else if is, params := b.App.MatchRoute("/frames/@([^/.]+)/label"); is {
		b.Component = newFramesEdit(b, params[0])
	} else if is, params := b.App.MatchRoute("/frames/@([^/.]+)/delete"); is {
		b.Component = newFramesDelete(b, params[0])
	} else if is, params := b.App.MatchRoute("/frames/@([^/.]+)/configure-frame"); is {
		b.Component = newFramesConfigure(b, params[0])
	} else if is, params := b.App.MatchRoute("/frames/@([^/.]+)/orchestrate"); is {
		b.Component = newFramesOrchestrate(b, params[0])
	} else if is, params := b.App.MatchRoute("/shapes/@([^/.]+).([^/.]+)"); is {
		b.Component = newShapesShow(b, params[0], params[1])
	} else if is, params := b.App.MatchRoute("/shapes/@([^/.]+).([^/.]+)/label"); is {
		b.Component = newShapesEdit(b, params[0], params[1])
	} else if is, params := b.App.MatchRoute("/shapes/@([^/.]+).([^/.]+)/delete"); is {
		b.Component = newShapesDelete(b, params[0], params[1])
	} else if is, params := b.App.MatchRoute("/shapes/@([^/.]+).([^/.]+)/configure-shape"); is {
		b.Component = newShapesConfigureShape(b, params[0], params[1])
	} else if is, params := b.App.MatchRoute("/shapes/@([^/.]+).([^/.]+)/configure-frame"); is {
		b.Component = newShapesConfigureFrame(b, params[0], params[1])
	} else {
		b.Component = newNotFound(b)
	}
	c = b.Component.Init()
	return
}

func (b *Body) focusChain() (fc []tuisupport.Focuser) {
	if b.Error == nil {
		fc = append(fc, b.Component.focusChain()...)
	} else {
		fc = b.Error.focusChain()
	}
	fc = append(fc,
		b.Scrollers[0],
		b.Scrollers[1],
		b.Scrollers[2],
		b.Scrollers[3],
	)
	return
}

func (b *Body) setScrollers() (c tea.Cmd) {
	b.Scrollers = []*tuisupport.Scroller{
		tuisupport.NewScroller("▲", tea.KeyCtrlUp),
		tuisupport.NewScroller("▼", tea.KeyCtrlDown),
		tuisupport.NewScroller("◀", tea.KeyCtrlLeft),
		tuisupport.NewScroller("▶", tea.KeyCtrlRight),
	}
	c = tea.Batch(
		b.Scrollers[0].Init(),
		b.Scrollers[1].Init(),
		b.Scrollers[2].Init(),
		b.Scrollers[3].Init(),
	)
	return
}

func (b *Body) setFooter() {
	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	help := helpStyle.Render("Scrolling ctrl+pgdn ctrl+pgup ctrl+↓ ctrl+↑ ctrl+← ctrl+→")
	scrollersStyle := lipgloss.NewStyle().Width(10).Align(lipgloss.Right)
	b.Footer = stickers.NewFlexBox(b.Width, 1)
	b.Footer.AddRows([]*stickers.FlexBoxRow{
		b.Footer.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(9, 1).SetContent(help),
				stickers.NewFlexBoxCell(1, 1).SetMinWidth(10).SetStyle(scrollersStyle).SetContent(
					lipgloss.JoinHorizontal(lipgloss.Top,
						b.Scrollers[0].View(),
						b.Scrollers[1].View(),
						b.Scrollers[2].View(),
						b.Scrollers[3].View(),
					),
				),
			},
		),
	})
}
