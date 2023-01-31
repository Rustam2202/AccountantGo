package gui

import (
	"accounter/db"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	idWidth    = 40
	width      = 150
	height     = 35
	tableWidth = 850
)

func MakeTable(dateFrom time.Time, dateTo time.Time, dataBase *db.Database) (fyne.CanvasObject, error) {

	data, err := dataBase.CalculateRecords(dateFrom, dateTo)
	if err != nil {
		return nil, err
	}

	tableWithHead := container.NewWithoutLayout()
	tableWithHead.Add(tableHeader())

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

	table.Resize(fyne.Size{Width: tableWidth, Height: 400})
	table.Move(fyne.NewPos(0, 60))
	table.SetColumnWidth(0, idWidth)
	table.SetColumnWidth(1, width)
	table.SetColumnWidth(2, width)
	table.SetColumnWidth(3, width)
	table.SetColumnWidth(4, width)
	for i := 0; i < len(data); i++ {
		if data[i][2] != "" && data[i][3] != "" {
			table.SetRowHeight(i, height*2)
		}
	}
	tableWithHead.Add(table)
	return tableWithHead, nil
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
	tableHeader.Resize(fyne.Size{Width: tableWidth, Height: height})
	tableHeader.SetColumnWidth(0, idWidth)
	tableHeader.SetColumnWidth(1, width)
	tableHeader.SetColumnWidth(2, width)
	tableHeader.SetColumnWidth(3, width)
	tableHeader.SetColumnWidth(4, width)
	tableHeader.SetRowHeight(0, height)
	return tableHeader
}
