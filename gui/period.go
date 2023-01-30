package gui

import (
	"accounter/db"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"

	"fyne.io/fyne/v2/widget"
)

var months = []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sep", "Oct", "Nov,", "Dec"}

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
	dateFromEntry.SetPlaceHolder(format1)
	dateToEntry := widget.NewEntryWithData(dateToBind)
	dateToEntry.SetPlaceHolder(format1)
	monthEntry := widget.NewSelect(months, func(s string) {})
	yearEntry := widget.NewSelectEntry(years())

	fromBtn := CalendarBtn(dateFromBind, win)
	toBtn := CalendarBtn(dateToBind, win)

	confirmBtn := widget.NewButton("Show", func() {

		//	table := MakeTable( dataBase)
		//	cont.AddObject(table)
		cont.Show()
	})

	c := container.NewVBox(
		container.NewGridWithColumns(6,
			empty, labelPeriod, empty, empty, labelMonth, empty,
			fromLabel, dateFromEntry, fromBtn, empty, monthEntry, monthLabel,
			toLabel, dateToEntry, toBtn, empty, yearEntry, yearLabel,
		),
		confirmBtn,
	)
	return c
}

// Make 3 (2 last and current) years for yearEntry widget
func years() []string {
	var years [3]string
	timeNow := time.Now()
	yearNow := timeNow.Year()
	for i, j := 0, len(years)-1; i < len(years); i, j = i+1, j-1 {
		years[i] = strconv.Itoa(yearNow - j)
	}
	return years[:]
}
