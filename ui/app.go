package ui

import (
	"stockgo/internal/service"
	config "stockgo/ui/config"
	"stockgo/ui/dashboard"
	"stockgo/ui/sales"
	"stockgo/ui/stock"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// TODO:
// [x] Search in Stock
// [x] Dashboard tab
// [x] ReportService
// [x] Sales Today / Revenue Month / Profit Month
// Low Stock card
// Recent Sales list
// [x] Searchable product picker
// Sale details dialog
// switch ui to spanish

type Refresher interface {
	Refresh()
}

type App struct {
	window fyne.Window

	dashboard *dashboard.View
	stock     *stock.View
	sales     *sales.View

	refreshers []Refresher
}

func New(app fyne.App, services service.Services) *App {
	window := app.NewWindow("Gestor de Ventas")
	a := &App{
		window: window,
	}

	a.dashboard = dashboard.New(
		window,
		services.Report,
	)

	a.stock = stock.New(
		window,
		services.Stock,
		a.RefreshAll,
	)

	a.sales = sales.New(
		window,
		services.Sale,
		services.Stock,
		services.Report,
		a.RefreshAll,
	)

	a.refreshers = []Refresher{
		a.dashboard,
		a.stock,
		a.sales,
	}

	return a
}

func (a *App) Run() {
	a.RefreshAll()

	tabs := container.NewAppTabs(
		container.NewTabItem(config.DashboardTitle, a.dashboard.Content()),
		container.NewTabItem(config.StockTitle, a.stock.Content()),
		container.NewTabItem(config.SalesTitle, a.sales.Content()),
	)

	a.window.SetContent(tabs)
	a.window.Resize(fyne.NewSize(config.WindowWidth, config.WindowHeight))
	a.window.CenterOnScreen()
	a.window.ShowAndRun()
}

func (a *App) RefreshAll() {
	for _, r := range a.refreshers {
		r.Refresh()
	}
}
