package gui

import (
	"accounter/db"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func PeriodDates(dataBase *db.Database, win fyne.Window) *fyne.Container {

	//	empty := widget.NewLabel("")
	label := widget.NewLabel("Enter period or month for report")
	label.Alignment = fyne.TextAlignCenter
	fromLabel := widget.NewLabel("From")
	toLabel := widget.NewLabel("To")

	dateFromBind := binding.BindString(nil)
	dateToBind := binding.BindString(nil)
	dateFromEntry := widget.NewEntryWithData(dateFromBind)
	dateFromEntry.SetPlaceHolder("01/01/2001")
	dateToEntry := widget.NewEntryWithData(dateToBind)
	dateToEntry.SetPlaceHolder("01/01/2001")

	fromBtn := CalendarBtn(dateFromBind, win)
	toBtn := CalendarBtn(dateToBind, win)

	//fromEntry := widget.NewEntry()
	//toEntry := widget.NewEntry()
	//	monthLabel := widget.NewLabel("Month:")
	//monthLabel.Alignment = fyne.TextAlignTrailing
	//	yearLabel := widget.NewLabel("Year:")
	//yearLabel.Alignment = fyne.TextAlignTrailing
	//	monthEntry := widget.NewSelect([]string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sep", "Oct", "Nov,", "Dec"}, func(s string) {})
	//	yearEntry := widget.NewSelectEntry([]string{"2020", "2021", "2022", "2023"})

	confirmBtn := widget.NewButton("Show", func() {
		dataBase.ShowRecords(dateFromEntry.Text, dateToEntry.Text)
	})

	labelCont := container.NewHBox(label)
	//labelCont.MinSize().Width = fyne.Container.Layout.MinSize(50)
	fromToCont := container.NewVBox(fromLabel, toLabel)
	//fromEntry.Resize(fyne.NewSize(100,100))
	entryCont := container.NewVBox(dateFromEntry, dateToEntry)
	//entryCont.Resize(fyne.NewSize( 500,500))
	btnCont := container.NewVBox(fromBtn, toBtn)

	c := container.NewVBox(
		labelCont,
		container.NewHBox(
			fromToCont,
			entryCont,
			btnCont,
		),
		confirmBtn,
	)
	return c
}
