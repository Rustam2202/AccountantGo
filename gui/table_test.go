package gui

import (
	"accounter/db"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"testing"
)

func TestShowTable(t *testing.T) {
	var db db.Database
	db.CreateDataBase("testtable")
	db.AddIncome(999.99, "1998-01-23 12:45:56")
	db.AddIncome(570.75, "2000-02-28 00:01:59")
	db.AddIncome(1500, "2005-04-25 00:00:00")
	db.AddSpend(2000.02, "2010-12-31 23:59:59")
	db.AddSpend(3200, "2023-01-26 15:16:00")

	a := app.New()
	w := a.NewWindow("Accounter Manager")
	w.Resize(fyne.NewSize(800, 600))

	w.SetContent(MakeTable(&db))

	w.ShowAndRun()
}
