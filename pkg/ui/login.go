package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LoginScreen(window fyne.Window, onLogin func()) fyne.CanvasObject {
	userEntry := widget.NewEntry()
	userEntry.SetPlaceHolder("Usuario")
	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("Contraseña")

	loginBtn := widget.NewButton("Ingresar", func() {
		if userEntry.Text == "admin" && passEntry.Text == "admin" {
			onLogin()
		}
	})

	return container.NewVBox(
		widget.NewLabel("GymProgress Tracker"),
		userEntry,
		passEntry,
		loginBtn,
	)
}
