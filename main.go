package main

import (
	//acc "accounter/app"
	"accounter/db"
	gui "accounter/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	db.CreateDataBase()
	a := app.New()
	w := a.NewWindow("Accounter Manager")
	w.Resize(fyne.NewSize(600, 600))

	//var record acc.Record

	w.SetContent(
		container.NewVBox(
			gui.Operation(w),
			gui.PeriodDates(w),
		),
	)

	w.ShowAndRun()
	//gui.Dummy()
}
