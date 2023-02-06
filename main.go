//go:generate fyne bundle -o bundled.go Icon.png

package main

import (
	db "accounter/db"
	"accounter/gui"

	"fyne.io/fyne/v2/app"
)

var CalendByte []byte
var DataBase db.Database

func main() {
	app := app.New()

	// uncomment for packing .exe
	// app.SetIcon(resourceIconPng) 

	c := gui.NewApp()
	c.LoadUI(app)

	app.Run()
}
