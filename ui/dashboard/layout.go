package dashboard

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// TODO
// Add colors
// Empty states
//

func (v *View) Content() fyne.CanvasObject {
	title := widget.NewLabelWithStyle(
		"Inicio",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	//----------------------------------------------------------------
	// KPI cards
	//----------------------------------------------------------------

	row1 := container.NewGridWithColumns(
		2,
		widget.NewCard(
			"Valor Del Inventorio",
			"",
			v.inventoryValue,
		),

		widget.NewCard(
			"Productos",
			"",
			v.totalProducts,
		),
	)

	row2 := container.NewGridWithColumns(
		2,
		widget.NewCard(
			"Ventas Hoy",
			"",
			v.salesToday,
		),

		widget.NewCard(
			"Ingresos Hoy",
			"",
			v.revenueToday,
		),
	)

	row3 := container.NewGridWithColumns(
		2,
		widget.NewCard(
			"Ingresos Este Mes",
			"",
			v.revenueMonth,
		),

		widget.NewCard(
			"Ganancia Este Mes",
			"",
			v.profitMonth,
		),
	)

	//----------------------------------------------------------------
	// Low stock
	//----------------------------------------------------------------

	v.lowStockList = widget.NewList(
		func() int {
			return len(v.lowStockItems)
		},

		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},

		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(v.lowStockItems[id])
		},
	)

	//----------------------------------------------------------------
	// Recent sales
	//----------------------------------------------------------------

	v.recentSales = widget.NewList(
		func() int {
			return len(v.recentItems)
		},

		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},

		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(v.recentItems[id])
		},
	)

	lists := container.NewGridWithColumns(
		2,

		widget.NewCard(
			"Stock Bajo",
			"",
			v.lowStockList,
		),

		widget.NewCard(
			"Ventas recientes",
			"",
			v.recentSales,
		),
	)

	return container.NewBorder(
		container.NewVBox(
			title,
		),

		nil,
		nil,
		nil,

		container.NewVBox(
			row1,
			row2,
			row3,
			lists,
		),
	)
}
