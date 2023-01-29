package main

import (
	db "accounter/db"
	gui "accounter/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var DataBase db.Database
var CalendByte []byte

func main() {
	DataBase.CreateDataBase("tutelka")

	a := app.New()
	w := a.NewWindow("Accounter Manager")
	w.Resize(fyne.NewSize(400, 600))

	tableHeader := widget.NewTable(
		func() (int, int) { return 1, 5 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			label := co.(*widget.Label)
			switch tci.Col {
			case 0:
				label.SetText("No")
			case 1:
				label.SetText("Date")
			case 2:
				label.SetText("Income")
			case 3:
				label.SetText("Spend")
			case 4:
				label.SetText("Comment")
			default:
			}
		},
	)
	tableHeader.SetColumnWidth(0, 100)
	tableHeader.SetColumnWidth(1, 100)
	tableHeader.SetColumnWidth(2, 100)
	tableHeader.SetColumnWidth(3, 100)
	tableHeader.SetColumnWidth(4, 100)
	tableHeader.SetRowHeight(0,100)

	table := widget.NewTable(
		func() (int, int) { return 0, 0 },
		func() fyne.CanvasObject { return widget.NewLabel("Cell") },
		func(tci widget.TableCellID, co fyne.CanvasObject) {},
	)
	table.SetColumnWidth(0, 100)
	table.SetColumnWidth(1, 100)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 100)
	table.SetColumnWidth(4, 100)

	table.Resize(fyne.NewSize(100, 100))
	//h:=fyne.Size{1000,1000}
	//table=table.CreateRenderer().MinSize().Height
	ContWithTable := container.NewMax(table)
	ContWithTable.Hide()
	//contWithTable.Show()

	w.SetContent(

		container.NewAdaptiveGrid(
			1,
			gui.AddOperation(&DataBase, w),
			gui.PeriodDates(ContWithTable, &DataBase, w),
		//	container.NewVBox(
				tableHeader,
				ContWithTable,
		//	),

		//	gui.MakeTable(&DataBase),
		),
	)

	w.ShowAndRun()
}
