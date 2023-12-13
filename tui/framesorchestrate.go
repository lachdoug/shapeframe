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
	WID         string
	ID          string
	Cancel      *tuisupport.Button
	Confirm     *tuisupport.Button
	IsStreaming bool
	Streamer    *Streamer
	Width       int
	Height      int
}

func newFramesOrchestrate(b *Body, wid string, id string) (fo *FramesOrchestrate) {
	fo = &FramesOrchestrate{Body: b, WID: wid, ID: id}
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
	id := fmt.Sprintf("%s.%s", fo.WID, fo.ID)
	msg := style.Render(fmt.Sprintf("Orchestrate frame %s?", idStyle.Render(id)))
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
	fo.Streamer.setSize(fo.Width, fo.Height-5)
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
	fo.Streamer = newStreamer()
}

func (fo *FramesOrchestrate) confirm() (c tea.Cmd) {
	result := fo.Body.call(
		controllers.FrameOrchestrationsCreate,
		&controllers.FrameOrchestrationsCreateParams{
			Frame: fo.ID,
		},
		"..",
	)
	if result != nil {
		fo.Cancel.Enabled(false)
		fo.Confirm.Enabled(false)
		fo.Streamer.setStream(result.Stream)
		fo.IsStreaming = true
		c = fo.Streamer.run()
	}
	return
}

func (fo *FramesOrchestrate) cancel() (c tea.Cmd) {
	c = Open("..")
	return
}

func (fo *FramesOrchestrate) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{
		fo.Cancel,
		fo.Confirm,
	}
	fc = append(fc, fo.Streamer.focusChain()...)
	return
}

// func (fo *FramesOrchestrate) Focus(aspect string) (c tea.Cmd) {
// 	return
// }

// func (fo *FramesOrchestrate) Blur() {}
