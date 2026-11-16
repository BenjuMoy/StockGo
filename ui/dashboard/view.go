package dashboard

import (
	"stockgo/internal/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type View struct {
	window fyne.Window

	reports *service.ReportService

	totalProducts  *widget.Label
	inventoryValue *widget.Label

	salesToday   *widget.Label
	revenueToday *widget.Label
	revenueMonth *widget.Label
	profitMonth  *widget.Label

	lowStockList *widget.List
	recentSales  *widget.List

	lowStockItems []string
	recentItems   []string
}

func New(window fyne.Window, reports *service.ReportService,
) *View {
	return &View{
		window: window,

		reports: reports,

		totalProducts:  widget.NewLabel("-"),
		inventoryValue: widget.NewLabel("-"),

		salesToday:   widget.NewLabel("-"),
		revenueToday: widget.NewLabel("-"),
		revenueMonth: widget.NewLabel("-"),
		profitMonth:  widget.NewLabel("-"),
	}
}
