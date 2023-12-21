package tui

import (
	"fmt"
	"sf/controllers"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FramesOrchestrate struct {
	Body        *Body
	ID          string
	Cancel      *tuisupport.Button
	Confirm     *tuisupport.Button
	IsStreaming bool
	Streamer    *tuisupport.Streamer
	Width       int
	Height      int
}

func newFramesOrchestrate(b *Body, id string) (fo *FramesOrchestrate) {
	fo = &FramesOrchestrate{Body: b, ID: id}
	return
}

func (fo *FramesOrchestrate) Init() (c tea.Cmd) {
	fo.setCancel()
	fo.setConfirm()
	fo.setStreamer()
	c = fo.Streamer.Init()
	return
}

func (fo *FramesOrchestrate) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = fo
	cs := []tea.Cmd{}
	_, c = fo.Confirm.Update(msg)
	cs = append(cs, c)
	_, c = fo.Cancel.Update(msg)
	cs = append(cs, c)
	_, c = fo.Streamer.Update(msg)
	cs = append(cs, c)
	c = tea.Batch(cs...)
	return
}

func (fo *FramesOrchestrate) View() (v string) {
	if fo.IsStreaming {
		v = fo.Streamer.View()
		return
	}
	style := lipgloss.NewStyle().Padding(1)
	idStyle := lipgloss.NewStyle().Bold(true)
	msg := style.Render(fmt.Sprintf("Orchestrate frame %s?", idStyle.Render(fo.ID)))
	v = lipgloss.JoinVertical(lipgloss.Left,
		msg,
		lipgloss.JoinHorizontal(lipgloss.Top,
			fo.Cancel.View(),
			fo.Confirm.View(),
		),
	)
	return
}

func (fo *FramesOrchestrate) setSize(w int, h int) {
	fo.Width = w
	fo.Height = h
	fo.setStreamerSize()
}

func (fo *FramesOrchestrate) setStreamerSize() {
	fo.Streamer.SetSize(fo.Width, fo.Height-2)
}

func (fo *FramesOrchestrate) setConfirm() {
	fo.Confirm = tuisupport.NewButton(
		fmt.Sprintf("frames-%s-delete-confirm", fo.ID),
		"Confirm",
		fo.confirm,
		15,
	)
}

func (fo *FramesOrchestrate) setCancel() {
	fo.Cancel = tuisupport.NewButton(
		fmt.Sprintf("frames-%s-delete-cancel", fo.ID),
		"Cancel",
		fo.cancel,
		9,
	)
}

func (fo *FramesOrchestrate) setStreamer() {
	fo.Streamer = tuisupport.NewStreamer()
}

func (fo *FramesOrchestrate) confirm() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = fo.Body.App.call(
		controllers.FrameOrchestrationsCreate,
		&controllers.FrameOrchestrationsCreateParams{
			Frame: fo.ID,
		},
		tuisupport.Open(".."),
	)
	if result != nil {
		fo.Cancel.Enabled(false)
		fo.Confirm.Enabled(false)
		fo.Streamer.SetStream(result.Stream)
		fo.IsStreaming = true
		c = fo.Streamer.Run()
	}
	return
}

func (fo *FramesOrchestrate) cancel() (c tea.Cmd) {
	c = tuisupport.Open("..")
	return
}

func (fo *FramesOrchestrate) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{
		fo.Cancel,
		fo.Confirm,
	}
	fc = append(fc, fo.Streamer.FocusChain()...)
	return
}

func (fo *FramesOrchestrate) isFocus() (is bool) {
	is = fo.Cancel.IsFocus || fo.Confirm.IsFocus
	return
}
