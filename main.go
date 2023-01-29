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
	w.Resize(fyne.NewSize(400, 600))

	ContWithTable := container.NewWithoutLayout()

	w.SetContent(
		container.NewGridWithColumns(
			1,
			container.NewVBox(
				gui.AddOperation(&DataBase, w),
				gui.PeriodDates(ContWithTable, &DataBase, w),
			),
			ContWithTable,
		),
	)
	w.ShowAndRun()
}
