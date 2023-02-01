package gui

import (
	"accounter/db"
	"accounter/utils"
	"errors"
	"fmt"
	"image/color"

	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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
	periodLabel := widget.NewLabel("Period:")
	periodLabel.Alignment = fyne.TextAlignCenter
	allIncomesLabel := widget.NewLabel("All incomes:")
	allIncomesLabel.Alignment = fyne.TextAlignCenter
	allSpendsLabel := widget.NewLabel("All spends:")
	allSpendsLabel.Alignment = fyne.TextAlignCenter
	totalValueLabel := widget.NewLabel("Total:")
	totalValueLabel.Alignment = fyne.TextAlignCenter
	period := widget.NewLabel("")
	period.Alignment = fyne.TextAlignCenter

	allIncomes := canvas.Text{}
	allIncomes.Alignment = fyne.TextAlignCenter
	allIncomes.Color = color.NRGBA{60, 179, 113, 255}
	allIncomes.TextSize = 15

	allSpends := canvas.Text{}
	allSpends.Alignment = fyne.TextAlignCenter
	allSpends.Color = color.NRGBA{255, 99, 71, 255}
	allSpends.TextSize = 15

	total := canvas.Text{}
	total.Alignment = fyne.TextAlignCenter
	total.TextSize = 15

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

	totalResults := container.NewHBox(
		periodLabel, period, allIncomesLabel, &allIncomes,
		allSpendsLabel, &allSpends, totalValueLabel, &total,
	)
	totalResults.Hide()

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

		period.SetText(fmt.Sprintf("%s ... %s", dateFrom.String(), dateTo.String()))
		allIncomes.Text = fmt.Sprintf("%0.2f", table.AllIncomes)
		allSpends.Text = fmt.Sprintf("%0.2f", table.AllSpends)
		total.Text = fmt.Sprintf("%0.2f", table.AllIncomes-table.AllSpends)
		totalResults.Show()

		cont.RemoveAll()
		cont.Add(table.Table)
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

		period.SetText(fmt.Sprintf("%s of %d", month.String(), year))
		allIncomes.Text = fmt.Sprintf("%0.2f", table.AllIncomes)
		allSpends.Text = fmt.Sprintf("%0.2f", table.AllSpends)
		total.Text = fmt.Sprintf("%0.2f", table.AllIncomes-table.AllSpends)
		totalResults.Show()

		cont.RemoveAll()
		cont.Add(table.Table)
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

		period.SetText(yearOfAnnualReportEntry.Text)
		allIncomes.Text = fmt.Sprintf("%0.2f", table.AllIncomes)
		allSpends.Text = fmt.Sprintf("%0.2f", table.AllSpends)
		total.Text = fmt.Sprintf("%0.2f", table.AllIncomes-table.AllSpends)
		totalResults.Show()

		cont.RemoveAll()
		cont.Add(table.Table)
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

		period.SetText("All period")
		allIncomes.Text = fmt.Sprintf("%0.2f", table.AllIncomes)
		allSpends.Text = fmt.Sprintf("%0.2f", table.AllSpends)
		total.Text = fmt.Sprintf("%0.2f", table.AllIncomes-table.AllSpends)
		totalResults.Show()
		cont.RemoveAll()
		cont.Add(table.Table)
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
		totalResults,
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
