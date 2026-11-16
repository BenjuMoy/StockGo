package sales

import (
	"stockgo/internal/domain"
	"stockgo/internal/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type View struct {
	window fyne.Window

	salesService  *service.SaleService
	stockService  *service.StockService
	reportService *service.ReportService

	sales []domain.Sale

	table *widget.Table

	salesToday   *widget.Label
	revenueToday *widget.Label

	refreshAll func()
}

func New(window fyne.Window, saleService *service.SaleService, stockService *service.StockService, reportService *service.ReportService, refreshAll func(),
) *View {
	return &View{
		window:        window,
		salesService:  saleService,
		stockService:  stockService,
		reportService: reportService,

		refreshAll: refreshAll,

		salesToday:   widget.NewLabel("-"),
		revenueToday: widget.NewLabel("-"),
	}
}
