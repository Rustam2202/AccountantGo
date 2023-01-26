package main

import (
	"accounter/db"
	gui "accounter/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var DataBase db.Database

func main() {
	DataBase.CreateDataBase()

	a := app.New()
	w := a.NewWindow("Accounter Manager")
	w.Resize(fyne.NewSize(600, 600))

	w.SetContent(
		container.NewVBox(
			gui.AddOperation(&DataBase, w),
			gui.PeriodDates(&DataBase, w),
		),
	)

	w.ShowAndRun()
}
