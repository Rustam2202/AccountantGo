package gui

import (
	"accounter/db"
	"accounter/utils"
	"errors"

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
	monthOfMonthlyReportLabel := widget.NewLabel("Report for month:")
	monthOfMonthlyReportLabel.Alignment = fyne.TextAlignTrailing
	yearOfMonthlyReportLabel := widget.NewLabel("Report for year:")
	yearOfMonthlyReportLabel.Alignment = fyne.TextAlignTrailing

	dateFromBind := binding.BindString(nil)
	dateToBind := binding.BindString(nil)
	dateFromEntry := widget.NewEntryWithData(dateFromBind)
	dateFromEntry.SetPlaceHolder(utils.Format1)
	dateToEntry := widget.NewEntryWithData(dateToBind)
	dateToEntry.SetPlaceHolder(utils.Format1)
	monthOfMonthlyReportEntry := widget.NewSelect(months, func(s string) {})
	yearOfMonthlyReportEntry := widget.NewSelectEntry(years())
	yearOfAnnualReportEntry := widget.NewSelectEntry(years())

	fromBtn := CalendarBtn(dateFromBind, win)
	toBtn := CalendarBtn(dateToBind, win)

	showPeriodBtn := widget.NewButton("Show period", func() {

		if dateFromEntry.Text == "" {
			dialog.ShowError(errors.New("Need enter period or month"), win)
			return
		}
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
		cont.RemoveAll()
		cont.Add(table)
		cont.Show()
	})
	showMonthBtn := widget.NewButton("Show month", func() {
		if monthOfMonthlyReportEntry.Selected == "" || yearOfMonthlyReportEntry.Text == "" {
			dialog.ShowError(errors.New("To Show need enter month and year"), win)
			return
		}
		month := time.Month(monthOfMonthlyReportEntry.SelectedIndex() + 1)
		year, err := strconv.Atoi(yearOfMonthlyReportEntry.Text)
		if err != nil {
			dialog.ShowError(errors.New("Year incorrect"), win)
			return
		}

		dateFrom := time.Date(year, month, 1, 0, 0, 0, 0, &time.Location{})
		dateTo := time.Date(year, month, 31, 0, 0, 0, 0, &time.Location{})

		//	dateFrom, err := utils.CheckDate(fmt.Sprintf("01.%d.%d", month, year))
		//	dateTo, err := utils.CheckDate(dateToEntry.Text)
		//	if err != nil {
		//		dialog.ShowError(err, win)
		//	return
		//	}

		table, err := MakeTable(dateFrom, dateTo, dataBase)
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		cont.RemoveAll()
		cont.Add(table)
		cont.Show()
	})
	showYearBtn := widget.NewButton("Show year", func() {
		if yearOfAnnualReportEntry.Text == "" {
			dialog.ShowError(errors.New("To Show need enter year"), win)
			return
		}
		year, err := strconv.Atoi(yearOfAnnualReportEntry.Text)
		if err != nil {
			dialog.ShowError(errors.New("Year incorrect"), win)
			return
		}

		dateFrom := time.Date(year, 1, 1, 0, 0, 0, 0, &time.Location{})
		dateTo := time.Date(year, 12, 31, 0, 0, 0, 0, &time.Location{})

		table, err := MakeTable(dateFrom, dateTo, dataBase)
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		cont.RemoveAll()
		cont.Add(table)
		cont.Show()

	})
	showAllBtn := widget.NewButton("Show all", func() {
		dateFrom := time.Date(2020, 1, 1, 0, 0, 0, 0, &time.Location{})
		dateTo := time.Now()

		table, err := MakeTable(dateFrom, dateTo, dataBase)
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		cont.RemoveAll()
		cont.Add(table)
		cont.Show()
	})

	return container.NewVBox(
		labelPeriod,
		container.NewHBox(
			container.NewGridWithColumns(4,
				fromLabel, dateFromEntry, fromBtn, monthOfMonthlyReportLabel,
				toLabel, dateToEntry, toBtn, yearOfMonthlyReportLabel,
			),
			container.NewGridWithRows(2,
				monthOfMonthlyReportEntry, yearOfAnnualReportEntry,
				yearOfMonthlyReportEntry, empty,
			),
		),
		container.NewGridWithColumns(4, showAllBtn, showPeriodBtn, showMonthBtn, showYearBtn),
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
