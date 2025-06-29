package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RegisterDayScreen(window fyne.Window, onDaySelected func(string)) fyne.CanvasObject {
	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("YYYY-MM-DD")

	exercises := []string{"Press Banca", "Sentadilla", "Peso muerto"}
	exerciseSelect := widget.NewSelect(exercises, nil)

	nextBtn := widget.NewButton("Siguiente", func() {
		if dateEntry.Text != "" && exerciseSelect.Selected != "" {
			onDaySelected(dateEntry.Text + "|" + exerciseSelect.Selected)
		}
	})

	return container.NewVBox(
		widget.NewLabel("Registrar Nuevo Día"),
		dateEntry,
		widget.NewLabel("Ejercicio:"),
		exerciseSelect,
		nextBtn,
	)
}
