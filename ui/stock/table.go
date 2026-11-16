package stock

import (
	"stockgo/ui/common"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (v *View) createTable() *widget.Table {
	headers := []string{
		"Nombre",
		"Cantidad",
		"Costo",
		"Valor",
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(v.filtered) + 1, len(headers)
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

			stock := v.filtered[id.Row-1]

			switch id.Col {
			case 0:
				label.SetText(stock.Name())

			case 1:
				label.SetText(strconv.Itoa(stock.Quantity()))

			case 2:
				label.SetText(common.FormatMoney(stock.Cost()))

			case 3:
				label.SetText(common.FormatMoney(stock.Quantity() * stock.Cost()))
			}
		},
	)

	table.SetRowHeight(0, 36)

	for i := 1; i < len(v.stocks)+1; i++ {
		table.SetRowHeight(i, 32)
	}

	table.SetColumnWidth(0, 220)
	table.SetColumnWidth(1, 90)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 120)

	v.editBtn.Disable()
	v.deleteBtn.Disable()

	table.OnSelected = func(id widget.TableCellID) {
		if id.Row == 0 {
			v.editBtn.Disable()
			v.deleteBtn.Disable()
			return
		}

		v.selected = id.Row - 1

		v.editBtn.Enable()
		v.deleteBtn.Enable()

	}

	return table
}
