package sales

import (
	"stockgo/ui/common"
	"strconv"

	"fyne.io/fyne/v2/dialog"
)

func (v *View) Refresh() {
	sales, err := v.salesService.GetAll()
	if err != nil {
		dialog.ShowError(err, v.window)
		return
	}
	v.sales = sales

	revenueToday, err := v.reportService.GetRevenueToday()
	if err != nil {
		dialog.ShowError(err, v.window)
		return
	}

	salesToday, err := v.reportService.CountSalesToday()
	if err != nil {
		dialog.ShowError(err, v.window)
		return
	}

	if v.salesToday != nil {
		v.salesToday.SetText(strconv.Itoa(salesToday))
	}

	if v.revenueToday != nil {
		v.revenueToday.SetText(common.FormatMoney(revenueToday))
	}

	if v.table != nil {
		v.table.Refresh()
	}
}
