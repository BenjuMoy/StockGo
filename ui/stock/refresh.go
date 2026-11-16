package stock

import (
	"stockgo/internal/domain"
	"strings"

	"fyne.io/fyne/v2/dialog"
)

func (v *View) Refresh() {
	stocks, err := v.service.GetAll()
	if err != nil {
		dialog.ShowError(err, v.window)
		return
	}

	v.stocks = stocks

	//if v.search == nil {
	v.applyFilter("")
	//} else {
	//	v.applyFilter(v.search.Text)
	//}

	if v.table != nil {
		v.table.Refresh()
	}

	v.selected = -1

	if v.editBtn != nil {
		v.editBtn.Disable()
	}

	if v.deleteBtn != nil {
		v.deleteBtn.Disable()
	}
}

func (v *View) applyFilter(text string) {
	name := strings.ToLower(
		strings.TrimSpace(text),
	)

	if name == "" {
		v.filtered = append([]domain.Product(nil), v.stocks...)
		return
	}

	v.filtered = v.filtered[:0]

	for _, stock := range v.stocks {
		if !strings.Contains(
			strings.ToLower(stock.Name()),
			name,
		) {
			continue
		}
		v.filtered = append(v.filtered, stock)
	}
}
