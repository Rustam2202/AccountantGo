package gui

import (
	

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var data_table = [][]string{
	[]string{"top left", "top right"}, 
	[]string{"bottom left", "bottom right"}}

func MakeTable() fyne.CanvasObject {
	list := widget.NewTable(
		func() (int, int) {
			return len(data_table), len(data_table[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data_table[i.Row][i.Col])
		})

/*
	t := widget.NewTable(
		func() (int, int) { return 500, 500 },
		func() fyne.CanvasObject {
			return widget.NewLabel("Cell 000, 000")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			label := co.(*widget.Label)
			switch tci.Col {
			case 0:
				label.SetText(fmt.Sprintf("%d", tci.Row+1))
			case 1:
				label.SetText("A longer cell")
			default:
				label.SetText(fmt.Sprintf("Cell %d, %d", tci.Row+1, tci.Col+1))
			}
		})
	t.Resize(fyne.NewSize(500, 500))
	t.SetColumnWidth(0, 20)
	t.SetColumnWidth(1, 102)
	t.SetRowHeight(2, 50)
	return t
	*/

	return list
}
