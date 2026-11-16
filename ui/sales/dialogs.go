package sales

import (
	"stockgo/internal/domain"
	"stockgo/ui/common"
	config "stockgo/ui/config"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (v *View) showCreateDialog() {
	stocks, err := v.stockService.GetAll()
	if err != nil {
		dialog.ShowError(err, v.window)
		return
	}

	if len(stocks) == 0 {
		dialog.ShowInformation(
			"No Stock",
			"No hay productos disponibles",
			v.window,
		)
		return
	}

	selected := stocks[0]

	product := widget.NewButton(
		selected.Name(),
		nil,
	)

	product.OnTapped = func() {
		common.SelectStock(
			v.window,
			stocks,

			func(s domain.Product) {
				selected = s
				product.SetText(s.Name())
			},
		)
	}

	priceEntry := widget.NewEntry()
	priceEntry.Validator = validation.NewRegexp(
		`^\d+$`,
		"Numbers only",
	)
	priceEntry.SetPlaceHolder("Precio")

	d := dialog.NewForm(
		"Añadir venta",
		"Agregar",
		"Cancelar",
		[]*widget.FormItem{
			widget.NewFormItem("Producto", product),
			widget.NewFormItem("Precio", priceEntry),
		},
		func(ok bool) {
			if !ok {
				return
			}

			price, err := strconv.Atoi(priceEntry.Text)
			if err != nil {
				dialog.ShowError(err, v.window)
				return
			}

			stockID := selected.ID()

			if _, err := v.salesService.CreateSale(stockID, price); err != nil {
				dialog.ShowError(err, v.window)
				return
			}

			v.refreshAll()

			common.Success(v.window, common.SaleCreated)

		},
		v.window,
	)

	d.Resize(fyne.NewSize(config.DialogWidth, config.DialogHeight))
	d.Show()
}
