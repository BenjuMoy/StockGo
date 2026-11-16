package sales

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (v *View) createToolbar() *fyne.Container {

	add := widget.NewButton(
		"Nueva venta",
		v.showCreateDialog,
	)

	return container.NewHBox(
		add,
	)
}
