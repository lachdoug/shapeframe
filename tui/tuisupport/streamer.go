package tuisupport

import (
	"fmt"
	"sf/app/streams"
	"sf/utils"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Streamer struct {
	Stream       *streams.Stream
	Output       string
	Viewport     viewport.Model
	Done         *Button
	IsFocus      bool
	IsUserScroll bool
	Width        int
	Height       int
}

type StreamerTickMsg struct {
	ID string
}

func NewStreamer() (str *Streamer) {
	str = &Streamer{}
	return
}

func (str *Streamer) Init() (c tea.Cmd) {
	str.setViewport()
	str.setDone()
	return
}

func (str *Streamer) SetStream(st *streams.Stream) {
	str.Stream = st
}

func (str *Streamer) Run() (c tea.Cmd) {
	go str.Stream.Read(str.write)
	c = tea.Batch(
		str.tick(),
		str.takeFocus(),
	)
	return
}

func (str *Streamer) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = str
	cs := []tea.Cmd{}
	switch msg := msg.(type) {
	case StreamerTickMsg:
		if !str.Stream.Complete && msg.ID == str.Stream.Identifier {
			c = str.tick()
			cs = append(cs, c)
		}
	case tea.KeyMsg:
		if str.IsFocus {
			switch msg.Type {
			case tea.KeyEnd:
				str.IsUserScroll = false
			default:
				str.IsUserScroll = true
			}
			str.Viewport, c = str.Viewport.Update(msg)
			cs = append(cs, c)
		}
	case tea.MouseMsg:
		if str.IsFocus {
			str.Viewport, c = str.Viewport.Update(msg)
			cs = append(cs, c)
		}
	}
	_, c = str.Done.Update(msg)
	cs = append(cs, c)
	c = tea.Batch(cs...)
	return
}

func (str *Streamer) View() (v string) {
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Width(str.Width - 2)

	if str.IsFocus {
		style.
			BorderForeground(lipgloss.Color("15"))
	} else {
		style.
			Foreground(lipgloss.Color("7")).
			BorderForeground(lipgloss.Color("8"))
	}

	lines := []string{}
	for _, l := range strings.Split(str.Output, "\n") {
		lines = append(lines, utils.StringChunks(l, str.Width-2)...)
	}

	str.Viewport.SetContent(strings.Join(lines, "\n"))
	if !str.IsUserScroll {
		str.Viewport.GotoBottom()
	}
	v = style.Render(str.Viewport.View())
	if str.Stream.Complete {
		str.Done.Enabled(true)
		v = lipgloss.JoinVertical(lipgloss.Left,
			v,
			str.Done.View(),
		)
	}
	return
}

func (str *Streamer) tick() (c tea.Cmd) {
	c = tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return StreamerTickMsg{
			ID: str.Stream.Identifier,
		}
	})
	return
}

func (str *Streamer) takeFocus() (c tea.Cmd) {
	c = TakeFocusCommand(str)
	return
}

func (str *Streamer) SetSize(w int, h int) {
	str.Width = w
	str.Height = h
	if str.Width < 10 {
		str.Width = 30
	}
	if str.Height < 10 {
		str.Height = 30
	}
	str.setViewportSize()
}

func (str *Streamer) setViewport() {
	str.Viewport = viewport.New(str.Width, str.Height)
}

func (str *Streamer) setDone() {
	str.Done = NewButton(
		"streamer-done",
		"Done",
		str.done,
		15,
	)
	str.Done.Enabled(false)
}

func (str *Streamer) done() (c tea.Cmd) {
	c = Open("..")
	return
}

func (str *Streamer) setViewportSize() {
	str.Viewport.Width = str.Width - 2
	str.Viewport.Height = str.Height - 4
}

func (str *Streamer) write(a ...any) {
	str.Output = str.Output + fmt.Sprint(a...)
}

func (str *Streamer) FocusChain() (fc []Focuser) {
	fc = []Focuser{
		str,
		str.Done,
	}
	return
}

func (str *Streamer) Focus(aspect string) (c tea.Cmd) {
	str.IsFocus = true
	return
}

func (str *Streamer) Blur() {
	str.IsFocus = false
}
