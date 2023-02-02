package gui

import (
	db "accounter/db"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type accounter struct {
	dataBase *db.Database
	win      fyne.Window
	// entries:
	incomeEntry, spendEntry, dateIncomEntry, dateSpendEntry, commentIncomEntry, commentSpendEntry,
	dateFromEntry, dateToEntry, monthOfMonthlyReportEntry, yearOfMonthlyReportEntry,
	yearOfAnnualReportEntry *widget.Entry
	// buttons:
	calendarBtn1, calendarBtn2, AddBtn, fromBtn, toBtn,
	showPeriodBtn, showMonthBtn, showYearBtn, showAllBtn *widget.Button
}

func (acc accounter) LoadUI(app fyne.App) {
	acc.dataBase.Name = "test4"
	acc.dataBase.OpenAndCreateLocalDb()
	acc.win = app.NewWindow("Accounter")
	acc.win.SetContent(
		container.NewVBox(
			container.NewVBox(
				acc.makeAddBlock(),
			),
		),
	)
	acc.win.Resize(fyne.NewSize(900, 850))
	acc.win.Show()
}

func NewApp() *accounter {

	return &accounter{
		dataBase:                  &db.Database{},
		win:                       nil,
		incomeEntry:               widget.NewEntry(),
		spendEntry:                widget.NewEntry(),
		dateIncomEntry:            &widget.Entry{},
		dateSpendEntry:            &widget.Entry{},
		commentIncomEntry:         &widget.Entry{},
		commentSpendEntry:         &widget.Entry{},
		dateFromEntry:             &widget.Entry{},
		dateToEntry:               &widget.Entry{},
		monthOfMonthlyReportEntry: &widget.Entry{},
		yearOfMonthlyReportEntry:  &widget.Entry{},
		yearOfAnnualReportEntry:   &widget.Entry{},
		calendarBtn1:              &widget.Button{},
		calendarBtn2:              &widget.Button{},
		AddBtn:                    &widget.Button{},
		fromBtn:                   &widget.Button{},
		toBtn:                     &widget.Button{},
		showPeriodBtn:             &widget.Button{},
		showMonthBtn:              &widget.Button{},
		showYearBtn:               &widget.Button{},
		showAllBtn:                &widget.Button{},
	}
}
