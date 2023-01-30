package gui

import (
	"accounter/db"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	no         = 40
	width      = 150
	height     = 35
	tableWidth = 600
)

func MakeTable(dateFrom time.Time, dateTo time.Time, dataBase *db.Database) fyne.CanvasObject {
	tableWithHead := container.NewWithoutLayout()
	tableWithHead.Add(tableHeader())

	data := dataBase.CalculateRecords(dateFrom, dateFrom)
	table := widget.NewTable(
		func() (int, int) {
			return len(data), 5
		},
		func() fyne.CanvasObject {
			lable := widget.NewLabel("")
			lable.Alignment = fyne.TextAlignCenter
			return lable
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

	table.Resize(fyne.Size{tableWidth, 300})
	table.Move(fyne.NewPos(0, 60))
	table.SetColumnWidth(0, no)
	table.SetColumnWidth(1, width)
	table.SetColumnWidth(2, width)
	table.SetColumnWidth(3, width)
	table.SetColumnWidth(4, width)
	table.SetRowHeight(0, height)
	tableWithHead.Add(table)
	return tableWithHead
}

func tableHeader() *widget.Table {
	tableHeader := widget.NewTable(
		func() (int, int) { return 1, 5 },
		func() fyne.CanvasObject {
			lable := widget.NewLabel("")
			lable.Alignment = fyne.TextAlignCenter
			return lable
		},
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
	tableHeader.Resize(fyne.Size{tableWidth, height})
	tableHeader.SetColumnWidth(0, no)
	tableHeader.SetColumnWidth(1, width)
	tableHeader.SetColumnWidth(2, width)
	tableHeader.SetColumnWidth(3, width)
	tableHeader.SetColumnWidth(4, width)
	tableHeader.SetRowHeight(0, height)
	return tableHeader
}
