package ui

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func WorkoutMatrixScreen(window fyne.Window, data string, onSave func([][]string)) fyne.CanvasObject {
	parts := strings.Split(data, "|")
	date, exercise := parts[0], parts[1]

	matrix := [][]string{}
	var matrixUI *widget.Table
	refreshMatrixUI := func() {
		matrixUI = widget.NewTable(
			func() (int, int) { return len(matrix), 3 },
			func() fyne.CanvasObject { return widget.NewEntry() },
			func(tci widget.TableCellID, co fyne.CanvasObject) {
				entry := co.(*widget.Entry)
				entry.SetText(matrix[tci.Row][tci.Col])
				entry.OnChanged = func(text string) {
					matrix[tci.Row][tci.Col] = text
				}
			},
		)
		matrixUI.SetColumnWidth(0, 100)
		matrixUI.SetColumnWidth(1, 100)
		matrixUI.SetColumnWidth(2, 100)
	}
	addSeriesBtn := widget.NewButton("Añadir Serie", func() {
		matrix = append(matrix, []string{"", "", ""})
		refreshMatrixUI()
	})

	saveBtn := widget.NewButton("Guardar entrenamioento", func() {
		onSave(matrix)
	})
	return container.NewVBox(
		widget.NewLabel(fmt.Sprintf("Fecha: %s - Ejercicio: %s", date, exercise)),
		addSeriesBtn,
		matrixUI,
		saveBtn)
}
