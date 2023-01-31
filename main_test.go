package main

import (
	db "accounter/db"
	gui "accounter/gui"
	"math/rand"
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)



func TestShow(t *testing.T) {
	var DataBase db.Database
	DataBase.Name = "test_show"
	err := DataBase.OpenDataBase(DataBase.Name)
	if err != nil {
		DataBase.CreateDataBase("tutelka")
	}

	a := app.New()
	w := a.NewWindow("Accounter Manager")
	w.Resize(fyne.NewSize(300, 600))

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


