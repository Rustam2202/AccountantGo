package gui

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

var months = []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sep", "Oct", "Nov,", "Dec"}

func (acc *accounter) makeSelect(s *widget.Select) *widget.Select {
	s = widget.NewSelect(months, func(s string) {})
	return s
}

func (acc *accounter) makeSelectWithEntry(s *widget.SelectEntry) *widget.SelectEntry {
	s = widget.NewSelectEntry(years())
	return s
}

func (acc *accounter) makeReportBlock() *fyne.Container {
	acc.totalResults.Hide()
	return container.NewVBox(
		acc.makeLabel("Enter period to show report", 1), // header
		container.NewHBox(
			container.NewGridWithColumns(4,
				acc.makeLabel("Period begin:", 2),
				acc.makeEntryWithData(acc.dateFromBind, acc.dateFromEntry),
				CalendarBtn(acc.dateFromBind, acc.win),
				acc.makeLabel("Month and year:", 2),
				acc.makeLabel("Period end:", 2),
				acc.makeEntryWithData(acc.dateToBind, acc.dateToEntry),
				CalendarBtn(acc.dateToBind, acc.win),
				acc.makeLabel("Year:", 2),
			),
			container.NewGridWithRows(2,
				acc.makeSelect(acc.monthOfMonthlyReportEntry), acc.makeSelectWithEntry(acc.yearOfAnnualReportEntry),
				acc.makeSelectWithEntry(acc.yearOfMonthlyReportEntry),
			),
		),
		container.NewGridWithColumns(4,
			acc.MakeButton(acc.showAllBtn, "Show all", acc.showAll),
			acc.showPeriodBtn, acc.showMonthBtn, acc.showYearBtn),
		acc.totalResults,
	)
}

func (acc *accounter) showAll() {
	acc.dateFromEntry.Text = "01.01.2019" // time.Date(2020, 1, 1, 0, 0, 0, 0, &time.Location{})
	acc.dateToEntry.Text = "02.02.2023"   //time.Now()

	table := acc.MakeTable()
	//	if err != nil {
	//		dialog.ShowError(err, acc.win)
	return
	//	}
	//	table.Hide()
	acc.totalResults.RemoveAll()
	acc.periodLabel.Text = "All period"
	// acc.periodLabel.SetText("All period")
	acc.makeTotal()
	acc.totalResults.Add(table)
	acc.totalResults.Show()
}

/*
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
	periodLabel.Alignment = fyne.TextAlignTrailing
	allIncomesLabel := widget.NewLabel("All incomes:")
	allIncomesLabel.Alignment = fyne.TextAlignTrailing
	allSpendsLabel := widget.NewLabel("All spends:")
	allSpendsLabel.Alignment = fyne.TextAlignTrailing
	totalValueLabel := widget.NewLabel("Total:")
	totalValueLabel.Alignment = fyne.TextAlignTrailing
	period := widget.NewLabel("")
	period.Alignment = fyne.TextAlignLeading
	allIncomes := canvas.Text{}
	allIncomes.Alignment = fyne.TextAlignLeading
	allIncomes.Color = color.NRGBA{60, 179, 113, 255}
	allIncomes.TextSize = 15
	allSpends := canvas.Text{}
	allSpends.Alignment = fyne.TextAlignLeading
	allSpends.Color = color.NRGBA{255, 99, 71, 255}
	allSpends.TextSize = 15
	total := canvas.Text{}
	total.Alignment = fyne.TextAlignLeading
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

	//fromBtn := CalendarBtn(dateFromBind, win)
	//toBtn := CalendarBtn(dateToBind, win)

	totalResults := container.NewGridWithColumns(8,
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
			container.NewGridWithColumns(4), //			fromLabel, dateFromEntry, fromBtn, monthOfMonthlyReportLabel,
			//				toLabel, dateToEntry, toBtn, yearOfMonthlyReportLabel,

			container.NewGridWithRows(2,
				monthOfMonthlyReportEntry, yearOfAnnualReportEntry,
				yearOfMonthlyReportEntry, empty,
			),
		),
		container.NewGridWithColumns(4, showAllBtn, showPeriodBtn, showMonthBtn, showYearBtn),
		totalResults,
	)
}
*/

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
