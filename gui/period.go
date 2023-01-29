package gui

import (
	"accounter/db"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"

	"fyne.io/fyne/v2/widget"
)

func PeriodDates(cont *fyne.Container, dataBase *db.Database, win fyne.Window) *fyne.Container {

	empty := widget.NewLabel("")
	labelPeriod := widget.NewLabel("Enter period:")
	labelPeriod.Alignment = fyne.TextAlignCenter
	labelMonth := widget.NewLabel(" or month for report:")
	labelMonth.Alignment = fyne.TextAlignCenter
	fromLabel := widget.NewLabel("From")
	toLabel := widget.NewLabel("To")
	monthLabel := widget.NewLabel("Month")
	monthLabel.Alignment = fyne.TextAlignLeading
	yearLabel := widget.NewLabel("Year")
	yearLabel.Alignment = fyne.TextAlignLeading

	dateFromBind := binding.BindString(nil)
	dateToBind := binding.BindString(nil)
	dateFromEntry := widget.NewEntryWithData(dateFromBind)
	dateFromEntry.SetPlaceHolder("01/01/2001")
	dateToEntry := widget.NewEntryWithData(dateToBind)
	dateToEntry.SetPlaceHolder("01/01/2001")
	monthEntry := widget.NewSelect([]string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sep", "Oct", "Nov,", "Dec"}, func(s string) {})
	yearEntry := widget.NewSelectEntry([]string{"2020", "2021", "2022", "2023"})

	fromBtn := CalendarBtn(dateFromBind, win)
	toBtn := CalendarBtn(dateToBind, win)

	confirmBtn := widget.NewButton("Show", func() {
		/*
			w := fyne.CurrentApp().NewWindow("Hello")
			w.SetContent(MakeTable(dataBase))
			w.Resize(fyne.NewSize(400,400))
			w.Show()
		*/
		table:=MakeTable(dataBase)
	//	table.Move(fyne.NewPos(0,100))
		cont.AddObject(table)
	//	cont.Add(table)
		cont.Show()
		//t := MakeTable(dataBase)
		//t.Resize(fyne.NewSize(100, 100))
		//dialog.ShowCustom("title", "dissmis", MakeTable(dataBase), win)
	})
	// confirmBtn.Resize(fyne.NewSize(100,100))

	c := container.NewVBox(
		// 3,
		container.NewGridWithColumns(6,
			empty, labelPeriod, empty, empty, labelMonth, empty,
			fromLabel, dateFromEntry, fromBtn, empty, monthEntry, monthLabel,
			toLabel, dateToEntry, toBtn, empty, yearEntry, yearLabel,
		),
		confirmBtn,
	)
	// c.Add(container.NewMax(MakeTable(dataBase)))
	return c
}
