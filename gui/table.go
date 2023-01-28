package gui

import (
	"accounter/db"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func MakeTable(dataBase *db.Database) fyne.CanvasObject {
	data := dataBase.CalculateRecords("", "")
	table := widget.NewTable(
		func() (int, int) {
			return len(data), 5
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			switch i.Col {
			case 0:
				label.SetText(data[i.Row][0])
			case 1:
				label.SetText(data[i.Row][1])
			case 2:
				label.SetText(data[i.Row][2])
			case 3:
				label.SetText(data[i.Row][3])
			case 4:
				label.SetText(data[i.Row][4])
			default:
			}
		})

	table.SetColumnWidth(0, 30)
	table.SetColumnWidth(1, 100)
	table.SetColumnWidth(2, 60)
	table.SetColumnWidth(3, 60)
	table.SetColumnWidth(4, 100)

	return table
}
