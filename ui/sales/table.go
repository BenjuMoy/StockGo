package sales

import (
	"fmt"
	"stockgo/ui/common"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (v *View) createTable() *widget.Table {
	headers := []string{
		"ID",
		"Producto",
		"Precio",
		"Fecha",
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(v.sales) + 1, len(headers)
		},

		func() fyne.CanvasObject {
			label := widget.NewLabelWithStyle(
				"",
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: true},
			)
			label.Alignment = fyne.TextAlignCenter
			return label
		},

		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)

			if id.Row == 0 {
				label.SetText(headers[id.Col])
				return
			}

			sale := v.sales[id.Row-1]

			switch id.Col {
			case 0:
				label.SetText(fmt.Sprint(sale.ID()))

			case 1:
				label.SetText(sale.ProductName())

			case 2:
				label.SetText(common.FormatMoney(sale.SoldPrice()))

			case 3:
				//label.SetText(sale.CreatedAt().Format("2006-01-02"))
				label.SetText(sale.CreatedAt().Format("02 Jan 2006  -  15:04"))
			}
		},
	)

	table.SetColumnWidth(0, 70)  // ID
	table.SetColumnWidth(1, 220) // Stock
	table.SetColumnWidth(2, 120) // Sold Price
	table.SetColumnWidth(3, 180) // Date

	return table
}
