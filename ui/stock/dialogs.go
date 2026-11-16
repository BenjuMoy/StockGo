package stock

import (
	"errors"
	"fmt"
	"stockgo/ui/common"
	config "stockgo/ui/config"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (v *View) showAddDialog() {
	nameEntry := widget.NewEntry()

	quantityEntry := widget.NewEntry()
	quantityEntry.Validator = validation.NewRegexp(
		`^\d+$`,
		"Numbers only",
	)

	costEntry := widget.NewEntry()
	costEntry.Validator = validation.NewRegexp(
		`^\d+$`,
		"Numbers only",
	)

	d := dialog.NewForm(
		"Añadir producto",
		"Agregar",
		"Cancelar",
		[]*widget.FormItem{
			widget.NewFormItem("Nombre", nameEntry),
			widget.NewFormItem("Cantidad", quantityEntry),
			widget.NewFormItem("Costo", costEntry),
		},
		func(ok bool) {
			if !ok {
				return
			}

			quantity, err := strconv.Atoi(quantityEntry.Text)
			if err != nil {
				dialog.ShowError(errors.New("Cantidad invalida"), v.window)
				return
			}

			cost, err := strconv.Atoi(costEntry.Text)
			if err != nil {
				dialog.ShowError(errors.New("Precio invalida"), v.window)
				return
			}

			if err := v.service.AddStock(nameEntry.Text, quantity, cost); err != nil {
				dialog.ShowError(err, v.window)
				return
			}

			common.Success(v.window, common.StockAdded)

			v.refreshAll()
		},
		v.window,
	)

	d.Resize(fyne.NewSize(config.DialogWidth, config.DialogHeight))
	d.Show()
	v.window.Canvas().Focus(nameEntry)
}

func (v *View) showEditDialog() {
	if v.selected < 0 {
		return
	}

	//stock := v.stocks[v.selected]
	stock := v.filtered[v.selected]

	stockID := widget.NewLabel(strconv.Itoa(stock.ID()))
	name := widget.NewEntry()
	qty := widget.NewEntry()
	cost := widget.NewEntry()

	name.SetText(stock.Name())
	qty.SetText(strconv.Itoa(stock.Quantity()))
	cost.SetText(strconv.Itoa(stock.Cost()))

	d := dialog.NewForm(
		"Editar Producto",
		"Agregar",
		"Cancelar",
		[]*widget.FormItem{
			widget.NewFormItem("ID", stockID),
			widget.NewFormItem("Nombre", name),
			widget.NewFormItem("Cantidad", qty),
			widget.NewFormItem("Costo", cost),
		},
		func(ok bool) {
			if !ok {
				return
			}

			quantity, err := strconv.Atoi(qty.Text)
			if err != nil {
				dialog.ShowError(
					errors.New("Cantidad invalida"),
					v.window,
				)
				return
			}

			cost, err := strconv.Atoi(cost.Text)
			if err != nil {
				dialog.ShowError(
					errors.New("Costo invalido"),
					v.window,
				)
				return
			}

			err = v.service.Update(stock.ID(), name.Text, quantity, cost)
			if err != nil {
				dialog.ShowError(err, v.window)
				return
			}

			common.Success(v.window, common.StockUpdated)

			v.refreshAll()
		},
		v.window,
	)

	d.Resize(fyne.NewSize(config.DialogWidth, config.DialogHeight))

	d.Show()

	v.window.Canvas().Focus(name)
}

func (v *View) showDeleteDialog() {
	if v.selected < 0 {
		return
	}

	//stock := v.stocks[v.selected]
	stock := v.filtered[v.selected]

	dialog.ShowConfirm(
		"Eliminar producto",
		fmt.Sprintf("Borrar '%s'?", stock.Name()),
		func(ok bool) {
			if !ok {
				return
			}

			if err := v.service.Delete(stock.ID()); err != nil {
				dialog.ShowError(err, v.window)
				return
			}

			v.refreshAll()
			v.table.UnselectAll()

			common.Success(v.window, common.StockDeleted)
		},
		v.window,
	)
}
