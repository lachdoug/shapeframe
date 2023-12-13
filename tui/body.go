package tui

import (
	"sf/app/errors"
	"sf/controllers"
	"sf/tui/tuisupport"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Body struct {
	App       *App
	Component Componenter
	Viewport  viewport.Model
	Error     *Error
	Width     int
	Height    int
}

func newBody(app *App) (b *Body) {
	b = &Body{App: app}
	return
}

func (b *Body) Init() (c tea.Cmd) {
	b.setViewport()
	b.setComponent()
	c = b.Component.Init()
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
	v = b.Viewport.View()
	return
}

func (b *Body) setSize(w int, h int) {
	b.Width = w
	b.Height = h
	b.setViewportSize()
	b.Component.setSize(w, h)
}

func (b *Body) setViewport() {
	b.Viewport = viewport.New(b.Width, b.Height)
}

func (b *Body) setViewportSize() {
	b.Viewport.Width = b.Width
	b.Viewport.Height = b.Height
}

func (b *Body) setComponent() {
	if is, _ := b.App.matchRoute("/workspaces"); is {
		b.Component = newWorkspacesIndex(b)
	} else if is, _ := b.App.matchRoute("/frames"); is {
		b.Component = newFramesIndex(b)
	} else if is, _ := b.App.matchRoute("/shapes"); is {
		b.Component = newShapesIndex(b)
	} else if is, _ := b.App.matchRoute("/workspaces/new"); is {
		b.Component = newWorkspacesNew(b)
	} else if is, _ := b.App.matchRoute("/frames/new"); is {
		b.Component = newFramesNew(b)
	} else if is, params := b.App.matchRoute("/workspaces/@([^/]+)"); is {
		b.Component = newWorkspacesShow(b, params[0])
	} else if is, params := b.App.matchRoute("/workspaces/@([^/]+)/delete"); is {
		b.Component = newWorkspacesDelete(b, params[0])
	} else if is, params := b.App.matchRoute("/workspaces/@([^/]+)/inspect"); is {
		b.Component = newWorkspacesInspect(b, params[0])
	} else if is, params := b.App.matchRoute("/frames/@([^/.]+).([^/.]+)"); is {
		b.Component = newFramesShow(b, params[0], params[1])
	} else if is, params := b.App.matchRoute("/frames/@([^/.]+).([^/.]+)/delete"); is {
		b.Component = newFramesDelete(b, params[0], params[1])
	} else if is, params := b.App.matchRoute("/frames/@([^/.]+).([^/.]+)/orchestrate"); is {
		b.Component = newFramesOrchestrate(b, params[0], params[1])
	} else if is, params := b.App.matchRoute("/shapes/@([^/.]+).([^/.]+).([^/.]+)"); is {
		b.Component = newShapesShow(b, params[0], params[1], params[2])
	} else if is, params := b.App.matchRoute("/shapes/@([^/.]+).([^/.]+).([^/.]+)/delete"); is {
		b.Component = newShapesDelete(b, params[0], params[1], params[2])
	} else {
		b.Component = newNotFound(b)
	}
}

func (b *Body) focusChain() (fc []tuisupport.Focuser) {
	fc = b.Component.focusChain()
	return
}

func (b *Body) call(
	controller func(*controllers.Params) (*controllers.Result, error),
	params any,
	returnPath string,
) (result *controllers.Result) {
	var err error

	if result, err = controller(&controllers.Params{Payload: params}); err == nil && result.Validation.IsInvalid() {
		err = errors.ValidationError(result.Validation.Maps())
	}

	if err != nil {
		result = nil
		b.Error = newError(
			err,
			b.errorReturnCallback(returnPath),
		)
		b.Error.Init()
	}

	return
}

func (b *Body) errorReturnCallback(returnPath string) (cb func() tea.Cmd) {
	cb = func() (c tea.Cmd) {
		b.Error = nil
		c = Open(returnPath)
		return
	}
	return
}
