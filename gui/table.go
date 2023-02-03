package gui

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TableWithTotal struct {
	Table      fyne.CanvasObject
	AllIncomes float32
	AllSpends  float32
}

const (
	idWidth    = 40
	width      = 150
	height     = 35
	tableWidth = 850
	textHeight = 15
	lauout     = "02.01.2006"
)

var Lime = color.NRGBA{0, 255, 0, 255}
var DarkRed = color.NRGBA{139, 0, 0, 255}

func (acc *accounter) makeCanvasText(value float32, color color.Color) *canvas.Text {
	text := canvas.NewText(fmt.Sprintf("%0.2f", value), color)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = textHeight
	return text
}

func (acc *accounter) makeTotal() *fyne.Container {
	//acc.totalResults.Add(
	return container.NewGridWithColumns(8,
		acc.periodLabel, acc.period, widget.NewLabel("All incomes"),
		acc.makeCanvasText(acc.allIncomes, Lime),
		widget.NewLabel("All spends"),
		acc.makeCanvasText(acc.allSpends, DarkRed),
		widget.NewLabel("Total"), acc.makeCanvasText(acc.total, Lime),
	)
	//)
}

func (acc *accounter) MakeTable( /*dateFrom time.Time, dateTo time.Time, dataBase *db.Database*/ ) *fyne.Container {
	dateFrom, err := time.Parse(lauout, acc.dateFromEntry.Text)
	dateTo, err := time.Parse(lauout, acc.dateToEntry.Text)
	data, err := acc.dataBase.CalculateRecords(dateFrom, dateTo)
	if err != nil {
		return nil
		//return TableWithTotal{}, err
	}

	//	result := TableWithTotal{}
	tableWithHead := container.NewWithoutLayout()

	table := widget.NewTable(
		func() (int, int) {
			return len(data.Data), 5
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
				label.SetText(data.Data[i.Row][0])
			case 1:
				label.SetText(data.Data[i.Row][1])
			case 2:
				label.SetText(data.Data[i.Row][2])
			case 3:
				label.SetText(data.Data[i.Row][3])
			case 4:
				label.SetText(data.Data[i.Row][4])
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
	for i := 0; i < len(data.Data); i++ {
		if data.Data[i][2] != "" && data.Data[i][3] != "" {
			table.SetRowHeight(i, height*2)
		}
	}

	acc.allIncomes = data.AllIncomes
	acc.allSpends = data.AllSpends
	acc.total = data.AllIncomes - data.AllSpends

	//acc.allIncomes.Text = fmt.Sprintf("%0.2f", data.AllIncomes)
	//acc.allSpends.Text = fmt.Sprintf("%0.2f", data.AllSpends)
	//acc.total.Text = fmt.Sprintf("%0.2f", data.AllIncomes-data.AllSpends)
	//	result.Table = tableWithHead
	//	result.AllIncomes = data.AllIncomes
	//	result.AllSpends = data.AllSpends
	//	return result, nil
	//	tableWithHead.Add(acc.showResults())
	tableWithHead.Add(tableHeader())
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
	tableHeader.Resize(fyne.Size{Width: tableWidth, Height: height})
	tableHeader.SetColumnWidth(0, idWidth)
	tableHeader.SetColumnWidth(1, width)
	tableHeader.SetColumnWidth(2, width)
	tableHeader.SetColumnWidth(3, width)
	tableHeader.SetColumnWidth(4, width)
	tableHeader.SetRowHeight(0, height)
	return tableHeader
}
