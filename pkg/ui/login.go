package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LoginScreen(window fyne.Window, onLogin func()) fyne.CanvasObject {
	userEntry := widget.NewEntry()
	userEntry.SetPlaceHoleder("Usuario")
}
