package gui

import (
	"accounter/utils"
	"errors"
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"

	"fyne.io/fyne/v2/widget"
)

var months = []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sep", "Oct", "Nov,", "Dec"}

func (acc *accounter) makeReportBlock() *fyne.Container {
	acc.monthOfMonthlyReportEntry = widget.NewSelect(months, func(s string) {})
	acc.yearOfMonthlyReportEntry = widget.NewSelectEntry(years())
	acc.yearOfAnnualReportEntry = widget.NewSelectEntry(years())
	return container.NewVBox(
		acc.makeLabel("Enter period to show report", 1), // header
		container.NewHBox(
			container.NewGridWithColumns(4,
				acc.makeLabel("Period begin:", allign(trail)),
				acc.makeEntryWithData(acc.dateFromBind, acc.dateFromEntry),
				CalendarBtn(acc.dateFromBind, acc.win),
				acc.makeLabel("Month and year:", allign(trail)),
				acc.makeLabel("Period end:", allign(trail)),
				acc.makeEntryWithData(acc.dateToBind, acc.dateToEntry),
				CalendarBtn(acc.dateToBind, acc.win),
				acc.makeLabel("Year:", allign(trail)),
			),
			container.NewGridWithRows(2,
				acc.monthOfMonthlyReportEntry,
				acc.yearOfAnnualReportEntry,
				acc.yearOfMonthlyReportEntry,
			),
		),
		container.NewGridWithColumns(4,
			acc.MakeButton(acc.showAllBtn, "Show all", acc.showAll),
			acc.MakeButton(acc.showAllBtn, "Show period", acc.showPeriod),
			acc.MakeButton(acc.showAllBtn, "Show monthly", acc.showMonth),
			acc.MakeButton(acc.showAllBtn, "Show annual", acc.showYear),
		),
	)
}

func (acc *accounter) showAll() {
	//acc.dateFromEntry.Text = "01.01.2019" // time.Date(2020, 1, 1, 0, 0, 0, 0, &time.Location{})
	//acc.dateToEntry.Text = "02.02.2023"   // time.Now()
	dateFrom := time.Date(2000, 1, 1, 0, 0, 0, 0, &time.Location{})
	dateTo := time.Now()
	acc.totalResults.RemoveAll()
	table := acc.MakeTable(dateFrom, dateTo)
	acc.period.Text = "All period"
	acc.totalResults.Add(acc.makeTotal())
	acc.totalResults.Add(table)
	acc.totalResults.Show()
}

func (acc *accounter) showMonth() {
	if acc.monthOfMonthlyReportEntry.Selected == "" || acc.yearOfMonthlyReportEntry.Text == "" {
		dialog.ShowError(errors.New(" To Show need enter month and year"), acc.win)
		return
	}
	month := time.Month(acc.monthOfMonthlyReportEntry.SelectedIndex() + 1)
	year, err := strconv.Atoi(acc.yearOfMonthlyReportEntry.Text)
	if err != nil {
		dialog.ShowError(errors.New(" Year incorrect"), acc.win)
		return
	}

	dateFrom := time.Date(year, month, 1, 0, 0, 0, 0, &time.Location{})
	dateTo := time.Date(year, month, 31, 0, 0, 0, 0, &time.Location{})

	acc.totalResults.RemoveAll()
	table := acc.MakeTable(dateFrom, dateTo)
	acc.period.Text = fmt.Sprintf("%s of %d", month.String(), year)
	acc.totalResults.Add(acc.makeTotal())
	acc.totalResults.Add(table)
	acc.totalResults.Show()
}

func (acc *accounter) showYear() {
	if acc.yearOfAnnualReportEntry.Text == "" {
		dialog.ShowError(errors.New(" To Show need enter year"), acc.win)
		return
	}
	year, err := strconv.Atoi(acc.yearOfAnnualReportEntry.Text)
	if err != nil {
		dialog.ShowError(errors.New(" Year incorrect"), acc.win)
		return
	}

	dateFrom := time.Date(year, 1, 1, 0, 0, 0, 0, &time.Location{})
	dateTo := time.Date(year, 12, 31, 0, 0, 0, 0, &time.Location{})

	acc.totalResults.RemoveAll()
	table := acc.MakeTable(dateFrom, dateTo)
	acc.period.Text = fmt.Sprintf("%d year", year)
	acc.totalResults.Add(acc.makeTotal())
	acc.totalResults.Add(table)
	acc.totalResults.Show()
}

func (acc *accounter) showPeriod() {
	if acc.dateFromEntry.Text == "" {
		dialog.ShowError(errors.New(" Need enter period"), acc.win)
		return
	}
	dateFrom, err1 := utils.CheckDateFormat(acc.dateFromEntry.Text)
	if err1 != nil {
		dialog.ShowError(err1, acc.win)
		return
	}
	dateTo, err2 := utils.CheckDateFormat(acc.dateToEntry.Text)
	if err2 != nil {
		dialog.ShowError(err2, acc.win)
		return
	}

	acc.totalResults.RemoveAll()
	table := acc.MakeTable(dateFrom, dateTo)
	acc.period.Text = fmt.Sprintf("From %s to %s", dateFrom.Format("02.01.2006"), dateTo.Format("02.01.2006"))
	acc.totalResults.Add(acc.makeTotal())
	acc.totalResults.Add(table)
	acc.totalResults.Show()
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
