package dashboard

import (
	"stockgo/ui/common"
	"strconv"

	"fyne.io/fyne/v2/dialog"
)

func (v *View) Refresh() {
	data, err := v.reports.GetDashboard()
	if err != nil {
		dialog.ShowError(err, v.window)
		return
	}

	//------------------------------------------------------------
	// Placeholder values
	//------------------------------------------------------------

	v.totalProducts.SetText(strconv.Itoa(data.ProductCount))

	v.inventoryValue.SetText(common.FormatMoney(data.InventoryValue))

	v.salesToday.SetText(strconv.Itoa(data.SalesToday))

	v.revenueToday.SetText(common.FormatMoney(data.RevenueToday))

	v.revenueMonth.SetText(common.FormatMoney(data.RevenueMonth))

	v.profitMonth.SetText(common.FormatMoney(data.ProfitMonth))

	//------------------------------------------------------------
	// Lists
	//------------------------------------------------------------

	v.lowStockItems = []string{
		"Milk (2)",
		"Rice (3)",
		"Sugar (1)",
	}

	v.recentItems = []string{
		"10:45  Coke      $8",
		"10:39  Bread     $4",
		"10:30  Rice      $12",
		"10:15  Water     $5",
	}

	//v.lowStockList.Refresh()
	//v.recentSales.Refresh()
}
