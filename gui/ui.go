package gui

import (
	db "accounter/db"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const (
	lead   int = 0
	center     = 1
	trail      = 2
)

type accounter struct {
	dataBase *db.Database
	win      fyne.Window
	// canvas text
	periodLabel, period, allIncomes, allSpends, total *canvas.Text
	// entries
	incomeEntry, spendEntry, dateIncomEntry, dateSpendEntry, commentIncomEntry, commentSpendEntry,
	dateFromEntry, dateToEntry *widget.Entry
	// buttons
	calendarBtn1, calendarBtn2, AddBtn, fromBtn, toBtn,
	showPeriodBtn, showMonthBtn, showYearBtn, showAllBtn *widget.Button
	// selects
	monthOfMonthlyReportEntry                         *widget.Select
	yearOfMonthlyReportEntry, yearOfAnnualReportEntry *widget.SelectEntry
	// binds
	dateIncomeBind, dateSpendBind, dateFromBind, dateToBind binding.String
	// container
	totalResults *fyne.Container
}

func (acc *accounter) makeLabel(text string, alig int) *widget.Label {
	label := widget.NewLabel(text)
	switch alig {
	case 0:
		label.Alignment = fyne.TextAlignLeading
	case 1:
		label.Alignment = fyne.TextAlignCenter
	case 2:
		label.Alignment = fyne.TextAlignTrailing
	default:
		label.Alignment = fyne.TextAlignCenter
	}

	return label
}

func (acc *accounter) makeEntry(ent *widget.Entry, placeholder string) *widget.Entry {
	ent.SetPlaceHolder(placeholder)
	return ent
}

func (acc *accounter) MakeButton(btn *widget.Button, label string, f func()) *widget.Button {
	btn = widget.NewButton(label, f)
	return btn
}

func (acc accounter) LoadUI(app fyne.App) {
	acc.dataBase.Name = "test4"
	acc.dataBase.OpenAndCreateLocalDb()
	acc.win = app.NewWindow("Accounter")
	acc.totalResults.Hide()
	acc.win.SetContent(
		container.NewVBox(
			container.NewVBox(
				acc.makeAddBlock(),
				acc.makeReportBlock(),
				acc.showResults(),
				acc.MakeTable(),
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
		periodLabel:               &canvas.Text{},
		period:                    &canvas.Text{},
		allIncomes:                &canvas.Text{},
		allSpends:                 &canvas.Text{},
		total:                     &canvas.Text{},
		incomeEntry:               widget.NewEntry(),
		spendEntry:                widget.NewEntry(),
		dateIncomEntry:            &widget.Entry{},
		dateSpendEntry:            &widget.Entry{},
		commentIncomEntry:         &widget.Entry{},
		commentSpendEntry:         &widget.Entry{},
		dateFromEntry:             &widget.Entry{},
		dateToEntry:               &widget.Entry{},
		calendarBtn1:              &widget.Button{},
		calendarBtn2:              &widget.Button{},
		AddBtn:                    &widget.Button{},
		fromBtn:                   &widget.Button{},
		toBtn:                     &widget.Button{},
		showPeriodBtn:             &widget.Button{},
		showMonthBtn:              &widget.Button{},
		showYearBtn:               &widget.Button{},
		showAllBtn:                &widget.Button{},
		monthOfMonthlyReportEntry: &widget.Select{},
		yearOfMonthlyReportEntry:  &widget.SelectEntry{},
		yearOfAnnualReportEntry:   &widget.SelectEntry{},
		dateIncomeBind:            binding.BindString(nil),
		dateSpendBind:             binding.BindString(nil),
		dateFromBind:              binding.BindString(nil),
		dateToBind:                binding.BindString(nil),
		totalResults:              &fyne.Container{},
	}
}
