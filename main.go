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
	tableHeader.Resize(fyne.Size{600, 50})
	tableHeader.SetColumnWidth(0, 100)
	tableHeader.SetColumnWidth(1, 100)
	tableHeader.SetColumnWidth(2, 100)
	tableHeader.SetColumnWidth(3, 100)
	tableHeader.SetColumnWidth(4, 100)
	tableHeader.SetRowHeight(0, 50)

	table1 := widget.NewTable(
		func() (int, int) { return 1, 5 },
		func() fyne.CanvasObject { return widget.NewLabel("cell") },
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			
		},
	)
	table1.Resize(fyne.Size{600, 50})
	table1.Move(fyne.NewPos(0, 60))
	table1.SetColumnWidth(0, 100)
	table1.SetColumnWidth(1, 100)
	table1.SetColumnWidth(2, 100)
	table1.SetColumnWidth(3, 100)
	table1.SetColumnWidth(4, 100)
	table1.SetRowHeight(0, 50)

	//h:=fyne.Size{1000,1000}
	//table=table.CreateRenderer().MinSize().Height
	ContWithTable := container.NewMax()
//	ContWithTable.Hide()
	//contWithTable.Show()

	w.SetContent(
		container.NewGridWithColumns(
			 1,
			
			container.NewVBox(
				gui.AddOperation(&DataBase, w),
				gui.PeriodDates(ContWithTable, &DataBase, w),
			),
			

			container.NewWithoutLayout(
				tableHeader,
				table1,
			),
		),
	)
	w.ShowAndRun()
}
