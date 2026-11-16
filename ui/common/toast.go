package common

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const margin = 16

func placeBottomRight(
	win fyne.Window,
	p *widget.PopUp,
	width,
	height float32,
) {
	size := win.Canvas().Size()

	p.Move(fyne.NewPos(
		size.Width-width-margin,
		size.Height-height-margin,
	))
}

func Success(win fyne.Window, msg string) {
	bg := canvas.NewRectangle(theme.Color(theme.ColorNameSuccess))

	label := widget.NewLabel(msg)

	content := container.NewStack(
		bg,
		container.NewPadded(label),
	)

	popup := widget.NewPopUp(
		content,
		win.Canvas(),
	)

	popup.Resize(fyne.NewSize(260, 48))

	placeBottomRight(win, popup, 260, 48)

	popup.Show()

	go func() {
		time.Sleep(2 * time.Second)

		fyne.Do(func() {
			popup.Hide()
		})
	}()
}
