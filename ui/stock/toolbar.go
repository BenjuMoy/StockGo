package stock

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (v *View) createToolbar() *fyne.Container {
	v.addBtn = widget.NewButton(
		"+ Añadir",
		v.showAddDialog,
	)

	v.editBtn = widget.NewButton(
		"✏ Editar",
		v.showEditDialog,
	)

	v.deleteBtn = widget.NewButton(
		"🗑 Eliminar",
		v.showDeleteDialog,
	)

	v.search = widget.NewEntry()

	v.search.SetPlaceHolder("Buscar...")

	v.search.OnChanged = func(text string) {
		v.applyFilter(text)

		v.table.Refresh()
	}

	return container.NewBorder(
		nil,
		nil,
		container.NewHBox(
			v.addBtn,
			v.editBtn,
			v.deleteBtn,
		),
		nil,
		v.search,
	)
}
