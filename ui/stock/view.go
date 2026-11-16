package stock

import (
	"stockgo/internal/domain"
	"stockgo/internal/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Add selected int here.

type View struct {
	window fyne.Window

	service *service.StockService

	stocks []domain.Product

	filtered []domain.Product

	search *widget.Entry

	table *widget.Table

	selected int

	addBtn    *widget.Button
	editBtn   *widget.Button
	deleteBtn *widget.Button

	refreshAll func()
}

func New(
	window fyne.Window,
	service *service.StockService,
	refreshAll func(),
) *View {
	v := &View{
		window:  window,
		service: service,

		refreshAll: refreshAll,

		selected: -1,
	}

	return v
}
