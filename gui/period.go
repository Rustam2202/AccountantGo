package gui

import (
	"accounter/db"
	"accounter/utils"

	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"

	"fyne.io/fyne/v2/widget"
)

var months = []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sep", "Oct", "Nov,", "Dec"}

func PeriodDates(cont *fyne.Container, dataBase *db.Database, win fyne.Window) *fyne.Container {

	empty := widget.NewLabel("")
	labelPeriod := widget.NewLabel("Enter period or month/year to show report")
	labelPeriod.Alignment = fyne.TextAlignCenter
	labelMonth := widget.NewLabel(" or month for report:")
	labelMonth.Alignment = fyne.TextAlignTrailing
	fromLabel := widget.NewLabel("Period begin:")
	fromLabel.Alignment = fyne.TextAlignTrailing
	toLabel := widget.NewLabel("Period end:")
	toLabel.Alignment = fyne.TextAlignTrailing
	monthLabel := widget.NewLabel("Report for month:")
	monthLabel.Alignment = fyne.TextAlignTrailing
	yearLabel := widget.NewLabel("Report for year:")
	yearLabel.Alignment = fyne.TextAlignTrailing

	dateFromBind := binding.BindString(nil)
	dateToBind := binding.BindString(nil)
	dateFromEntry := widget.NewEntryWithData(dateFromBind)
	dateFromEntry.SetPlaceHolder(utils.Format1)
	dateToEntry := widget.NewEntryWithData(dateToBind)
	dateToEntry.SetPlaceHolder(utils.Format1)
	monthEntry := widget.NewSelect(months, func(s string) {})
	yearOfMonth := widget.NewSelectEntry(years())
	yearEntry := widget.NewSelectEntry(years())

	fromBtn := CalendarBtn(dateFromBind, win)
	toBtn := CalendarBtn(dateToBind, win)

	confirmBtn := widget.NewButton("Show", func() {
		dateFrom, err := utils.CheckDate(dateFromEntry.Text)
		dateTo, err := utils.CheckDate(dateToEntry.Text)

		if err != nil {
			dialog.ShowError(err, win)
			return
		}

		table, err := MakeTable(dateFrom, dateTo, dataBase)
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		cont.Add(table)
		cont.Show()
	})

	return container.NewVBox(
		labelPeriod,
		container.NewHBox(
			container.NewGridWithColumns(4,
				fromLabel, dateFromEntry, fromBtn, monthLabel,
				toLabel, dateToEntry, toBtn, yearLabel,
			),
			container.NewGridWithRows(2,
				monthEntry, yearOfMonth,
				yearEntry, empty,
			),
		),
		confirmBtn,
	)
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
