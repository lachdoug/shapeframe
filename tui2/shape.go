package tui2

import (
	"fmt"
	"sf/controllers"
	"sf/models"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

type Shape struct {
	App        *App
	Components []ViewModeller
	Workspace  string
	Frame      string
	Shape      string
	Errors     []string
	Reader     *models.ShapeReader
}

func NewShape(app *App) (vm *Shape) {
	vm = &Shape{
		App:       app,
		Workspace: app.Workspace,
		Frame:     app.Frame,
		Shape:     app.Shape,
	}
	return
}

func (vm *Shape) Init() (c tea.Cmd) {
	var err error
	var result *controllers.Result
	params := &controllers.Params{
		Payload: &controllers.ShapesReadParams{
			Workspace: vm.Workspace,
			Frame:     vm.Frame,
			Shape:     vm.Shape,
		},
	}
	if result, err = controllers.ShapesRead(params); err != nil {
		vm.Errors = append(vm.Errors, err.Error())
	}
	vm.Errors = append(vm.Errors, "OH NO")
	vm.Errors = append(vm.Errors, "MORE OH NO")
	vm.Reader = result.Payload.(*models.ShapeReader)
	cs := []tea.Cmd{}
	for _, cmt := range vm.Components {
		cs = append(cs, cmt.Init())
	}
	c = tea.Batch(cs...)
	return
}

func (vm *Shape) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = vm
	// switch msg := msg.(type) {
	// case tea.MouseMsg:
	// 	switch msg.Type {
	// 	case tea.MouseLeft:
	// 		if zone.Get("exit-shape").InBounds(msg) {
	// 			vm.App.back()
	// 		}
	// 	}
	// }
	return
}

func (vm *Shape) View() (v string) {
	vb := NewViewBuilder()
	vb.Writeln(viewErrors(vm.Errors))
	if vm.App.Action == "show" {
		vm.show(vb)
	}
	if vm.App.Action == "edit" {
		vm.edit(vb)
	}
	v = vb.String()
	return
}

func (vm *Shape) show(vb *ViewBuilder) {
	vb.Writeln(vm.Reader.Name)
	vb.Writeln(vm.Reader.About)
	vb.Writeln(zone.Mark("back", buttonStyle.Render("Done")))
	vb.Writeln(zone.Mark("edit", buttonStyle.Render("Edit")))
}

func (vm *Shape) edit(vb *ViewBuilder) {
	// if w, err = models.NewWorkspace(uc, vm.App.Workspace,
	// 	"Frames",
	// ); err != nil {
	// 	return
	// }
	// if f, err = models.ResolveFrame(uc, w, frame,
	// 	"Shapes",
	// ); err != nil {
	// 	return
	// }
	// if s, err = models.ResolveShape(uc, f, shape,
	// 	"Configuration",
	// ); err != nil {
	// 	return
	// }

	vb.Writeln(fmt.Sprintf("Edit %v", vm.Reader.Name))
	// fm := cliform.NewForm(
	// 	"Shape configuration settings",
	// 	s.Shaper.Shape,
	// 	s.ShapeConfiguration.Settings,
	// )

	// vb.Writeln()
	vb.Writeln(newButton(vm, "#back-button", "Done", vm.App.back).View())
	// vb.Writeln(zone.Mark("back", buttonStyle.Render("Done")))
}

func (vm *Shape) app() (app *App) {
	app = vm.App
	return
}

func (vm *Shape) registerComponent(cmt ViewModeller) {
	vm.Components = append(vm.Components, cmt)
}

type ViewModeller interface {
	registerComponent(ViewModeller)
	// registerClickable(string)
	app() *App
	View() string
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
}

func newButton(parent ViewModeller, id string, text string, onclick func()) (vm ViewModeller) {
	vm = &Button{Parent: parent, ID: id, Text: text, Onclick: onclick}
	parent.registerComponent(vm)
	return
}

type Button struct {
	Parent     ViewModeller
	Components []ViewModeller
	ID         string
	Text       string
	Onclick    func()
	Hover      bool
}

func (btn *Button) app() (app *App) {
	app = btn.Parent.app()
	return
}

func (btn *Button) registerComponent(cmt ViewModeller) {
	btn.Components = append(btn.Components, cmt)
}

func (btn *Button) Init() (c tea.Cmd) {
	btn.app().registerClickable(btn.ID)
	return
}

func (btn *Button) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = btn
	switch msg := msg.(type) {
	case tea.MouseMsg:
		btn.Hover = true
		switch msg.Type {
		case tea.MouseLeft:
			btn.Onclick()
		}
	}
	return
}

func (btn *Button) View() (v string) {
	v = zone.Mark(btn.ID, buttonStyle.Render("Done"))
	return
}
