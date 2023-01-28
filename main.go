package main

import (
	db "accounter/db"
	gui "accounter/gui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var DataBase db.Database
var CalendByte []byte

func main() {
	DataBase.CreateDataBase("tutelka")

	a := app.New()
	w := a.NewWindow("Accounter Manager")
	w.Resize(fyne.NewSize(800, 600))

	w.SetContent(
		container.NewVBox(
			gui.AddOperation(&DataBase, w),
			gui.PeriodDates(&DataBase, w),

		//	gui.MakeTable(&DataBase),
		),
	)

	w.ShowAndRun()
}
