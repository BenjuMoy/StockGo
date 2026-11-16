package common

import (
	"fmt"
	"strings"

	"stockgo/internal/domain"
	config "stockgo/ui/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func SelectStock(
	window fyne.Window,
	stocks []domain.Product,
	onSelect func(domain.Product),
) {
	filtered := append([]domain.Product(nil), stocks...)

	search := widget.NewEntry()
	search.SetPlaceHolder("Buscar...")

	list := widget.NewList(

		func() int {
			return len(filtered)
		},

		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},

		func(id widget.ListItemID, obj fyne.CanvasObject) {

			s := filtered[id]

			obj.(*widget.Label).SetText(
				fmt.Sprintf(
					"%s (%d)",
					s.Name(),
					s.Quantity(),
				),
			)
		},
	)

	search.OnChanged = func(text string) {
		text = strings.ToLower(strings.TrimSpace(text))

		filtered = filtered[:0]

		for _, stock := range stocks {
			if text != "" &&
				!strings.Contains(
					strings.ToLower(stock.Name()),
					text,
				) {
				continue
			}
			filtered = append(filtered, stock)
		}

		list.Refresh()
	}

	d := dialog.NewCustom(
		"Seleccionar producto",
		"Close",
		container.NewBorder(
			search,
			nil,
			nil,
			nil,
			list,
		),
		window,
	)

	list.OnSelected = func(id widget.ListItemID) {
		d.Hide()
		onSelect(filtered[id])
	}

	d.Resize(fyne.NewSize(config.PickerWidth, config.PickerHeight))
	d.Show()
}
