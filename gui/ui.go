//go:generate fyne bundle -o bundled.go Calendar.png
//go:generate fyne bundle -o bundled.go -append day-night.png

package gui

import (
	db "accountant/db"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type allign fyne.TextAlign

const (
	lead   allign = 0
	center        = 1
	trail         = 2
)

type accounter struct {
	dataBase *db.Database
	win      fyne.Window

	// canvas text
	periodLabel, period *canvas.Text
	// float32
	allIncomes, allSpends, total float32
	// entries
	IncomeEntry, spendEntry, dateIncomEntry, dateSpendEntry, commentIncomEntry, commentSpendEntry,
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

func (acc *accounter) makeLabel(text string, al allign) *widget.Label {
	label := widget.NewLabel(text)
	switch al {
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

func (acc *accounter) makeEntryWithData(bind binding.String, ent *widget.Entry, placeholder string) *widget.Entry {
	ent.Bind(bind)
	ent.SetPlaceHolder(placeholder)
	return ent
}

func (acc *accounter) MakeButton(btn *widget.Button, label string, f func()) *widget.Button {
	btn = widget.NewButton(label, f)
	return btn
}

func makeThemeSettings(a fyne.App) *fyne.Container {
	light := widget.NewButtonWithIcon("", resourceLightPng, func() {
		a.Settings().SetTheme(theme.LightTheme())
	})
	dark := widget.NewButtonWithIcon("", resourceDarkPng, func() {
		a.Settings().SetTheme(theme.DarkTheme())
	})
	return container.NewVBox(light, dark, instructionBtn())
}

func (acc accounter) LoadUI(app fyne.App) {
	app.Settings().SetTheme(theme.LightTheme())
	acc.dataBase.Name = "database"
	acc.dataBase.OpenAndCreateLocalDb()
	acc.win = app.NewWindow("Accounter")
	clearBtn := widget.NewButton("Clear", func() {
		acc.totalResults.RemoveAll()
	})
	acc.win.SetContent(
		container.NewBorder(nil, clearBtn, nil, makeThemeSettings(app),
			container.NewVBox(
				container.NewVBox(
					acc.makeAddBlock(),
					acc.makeReportBlock(),
				),
				acc.totalResults,
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
		IncomeEntry:               &widget.Entry{},
		spendEntry:                &widget.Entry{},
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
		totalResults:              container.NewVBox(),
	}
}
