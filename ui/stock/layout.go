package stock

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (v *View) Content() fyne.CanvasObject {
	title := widget.NewLabelWithStyle(
		"Inventario",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	toolbar := v.createToolbar()

	v.table = v.createTable()

	content := container.NewBorder(
		container.NewVBox(
			title,
			toolbar,
		),
		nil,
		nil,
		nil,
		v.table,
	)
	content.Resize(fyne.NewSize(500, 400))

	return content
}
