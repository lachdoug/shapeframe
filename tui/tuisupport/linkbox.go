package tuisupport

import (
	"math"
	"strings"

	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LinkBox struct {
	ID       string
	MainText string
	LinkText string
	Path     string
	Link     *Link
	FlexBox  *stickers.FlexBox
	Width    int
	Height   int
}

func NewLinkBox(mainText string, id string, linkText string, path string) (lb *LinkBox) {
	lb = &LinkBox{MainText: mainText, ID: id, LinkText: linkText, Path: path}
	return
}

func (lb *LinkBox) Init() (c tea.Cmd) {
	lb.setLink()
	return
}

func (lb *LinkBox) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = lb
	_, c = lb.Link.Update(msg)
	return
}

func (lb *LinkBox) View() (v string) {
	lb.setFlexBox()
	style := lipgloss.NewStyle().Width(lb.Width)
	v = style.Render(lb.FlexBox.Render())
	return
}

func (lb *LinkBox) setFlexBox() {
	lb.FlexBox = stickers.NewFlexBox(lb.Width, lb.Height)
	lb.FlexBox.AddRows([]*stickers.FlexBoxRow{
		lb.FlexBox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(11, 1).SetContent(lb.MainText),
				stickers.NewFlexBoxCell(1, 1).SetContent(lb.renderLink()),
			},
		),
	})
}

func (lb *LinkBox) renderLink() (v string) {
	style := lipgloss.NewStyle().Width(lb.Width / 12).Align(lipgloss.Right)
	v = style.Render(lb.Link.View())
	return
}

func (lb *LinkBox) SetSize(w int) {
	lb.Width = w
	lb.Height = lb.totalLinesRequired()
}

func (lb *LinkBox) totalLinesRequired() (n int) {
	for _, line := range strings.Split(lb.MainText, "\n") {
		n = n + lb.linesRequired(line)
	}
	return
}

func (lb *LinkBox) linesRequired(line string) (n int) {
	textWidth := lipgloss.Width(line)
	colWidth := lb.Width * 11 / 12
	n = textWidth / colWidth
	if math.Mod(float64(textWidth), float64(colWidth)) > 0 {
		n = n + 1
	}
	return
}

func (lb *LinkBox) setLink() {
	lb.Link = NewLink(lb.ID, lb.LinkText, lb.Path)
}

func (lb *LinkBox) Focus(_ string) (c tea.Cmd) {
	lb.Link.IsFocus = true
	return
}

func (lb *LinkBox) Blur() {
	lb.Link.IsFocus = false
}

func (lb *LinkBox) IsFocus() (is bool) {
	is = lb.Link.IsFocus
	return
}
