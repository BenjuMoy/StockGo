package sales

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (v *View) Content() fyne.CanvasObject {
	title := widget.NewLabelWithStyle(
		"Ventas",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	toolbar := v.createToolbar()

	v.table = v.createTable()

	salesCard := widget.NewCard(
		"Ventas Hoy",
		"",
		v.salesToday,
	)

	revenueCard := widget.NewCard(
		"Ingresos Hoy",
		"",
		v.revenueToday,
	)

	summary := container.NewGridWithColumns(
		2,
		salesCard,
		revenueCard,
	)

	return container.NewBorder(
		container.NewVBox(
			title,
			toolbar,
			summary,
		),
		nil,
		nil,
		nil,
		v.table,
	)
}
