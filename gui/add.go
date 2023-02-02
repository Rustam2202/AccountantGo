package gui

import (
	"accounter/db"
	"accounter/utils"
	"errors"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (acc *accounter) makeEntryWithData(bind binding.String, ent *widget.Entry) *widget.Entry {
	ent.Bind(bind)
	return ent
}

func (acc *accounter) makeAddBlock() *fyne.Container {
	return container.NewVBox(
		acc.makeLabel("Enter income and/or spend to add in database", 1), // header
		container.NewHBox(
			container.NewGridWithColumns(4,
				acc.makeEntry(acc.incomeEntry, "Enter Income:"),
				acc.makeEntryWithData(acc.dateIncomeBind, acc.dateIncomEntry),
				CalendarBtn(acc.dateIncomeBind, acc.win),
				acc.makeEntry(acc.commentIncomEntry, "Enter comment"),
				acc.makeEntry(acc.spendEntry, "Enter Spend:"),
				acc.makeEntryWithData(acc.dateSpendBind, acc.dateSpendEntry),
				CalendarBtn(acc.dateSpendBind, acc.win),
				acc.makeEntry(acc.commentSpendEntry, "Enter comment"),
			),
			acc.MakeButton(acc.AddBtn, "Add", acc.AddBtnEvent),
		),
	)
}

func (acc *accounter) AddBtnEvent() {
	if !(acc.incomeEntry.Text != "" || acc.spendEntry.Text != "") {
		dialog.ShowError(errors.New("Income or Spend field must contain a value"), acc.win)
		return
	}

	var income, spend float32
	var dateInc, dateSpn time.Time
	var errInc, errSpn error
	income, dateInc, errInc = utils.CheckEntry(acc.incomeEntry.Text, acc.dateIncomEntry.Text)
	if errInc != nil {
		dialog.ShowError(errInc, acc.win)
		return
	}
	spend, dateSpn, errSpn = utils.CheckEntry(acc.spendEntry.Text, acc.dateSpendEntry.Text)
	if errSpn != nil {
		dialog.ShowError(errSpn, acc.win)
		return
	}
	if income != 0 && spend != 0 && utils.DatesCompare(dateInc, dateSpn) {
		acc.dataBase.AddIncomeAndSpend(income, spend, dateInc, acc.commentIncomEntry.Text, acc.commentSpendEntry.Text)
	} else { // income and spend can be both no zero but with different dates
		if income != 0 {
			acc.dataBase.AddIncome(income, dateInc, acc.commentIncomEntry.Text)
		}
		if spend != 0 {
			acc.dataBase.AddSpend(spend, dateSpn, acc.commentSpendEntry.Text)
		}
	}

	// need to fix notifications (drivers or something)
	fyne.CurrentApp().SendNotification(fyne.NewNotification("Add success", "Income added"))

	// clearing entry fields
	acc.incomeEntry.Text = ""
	acc.dateIncomEntry.Text = ""
	acc.spendEntry.Text = ""
	acc.dateSpendEntry.Text = ""
	acc.incomeEntry.Refresh()
	acc.dateIncomEntry.Refresh()
	acc.spendEntry.Refresh()
	acc.dateSpendEntry.Refresh()
}

func AddOperation(dataBase *db.Database, win fyne.Window) *fyne.Container {
	enterOperLabel := widget.NewLabel("Enter income and/or spend to add in database")
	enterOperLabel.Alignment = fyne.TextAlignCenter

	dateIncomeBind := binding.BindString(nil)
	dateSpendBind := binding.BindString(nil)

	incomeEntry := widget.NewEntry()
	incomeEntry.SetPlaceHolder("Enter Income")
	spendEntry := widget.NewEntry()
	spendEntry.SetPlaceHolder("Enter Spend")
	dateIncomEntry := widget.NewEntryWithData(dateIncomeBind)
	dateIncomEntry.SetPlaceHolder(time.Now().Format("02.01.2006"))
	dateSpendEntry := widget.NewEntryWithData(dateSpendBind)
	dateSpendEntry.SetPlaceHolder(time.Now().Format("02.01.2006"))
	commentIncomEntry := widget.NewEntry()
	commentIncomEntry.SetPlaceHolder("Enter comment")
	commentSpendEntry := widget.NewEntry()
	commentSpendEntry.SetPlaceHolder("Enter comment")

	addBtn := widget.NewButton("Add record", func() {
		if !(incomeEntry.Text != "" || spendEntry.Text != "") {
			dialog.ShowError(errors.New("Income or Spend field must contain a value"), win)
			return
		}

		var income, spend float32
		var dateInc, dateSpn time.Time
		var errInc, errSpn error
		income, dateInc, errInc = utils.CheckEntry(incomeEntry.Text, dateIncomEntry.Text)
		if errInc != nil {
			dialog.ShowError(errInc, win)
			return
		}
		spend, dateSpn, errSpn = utils.CheckEntry(spendEntry.Text, dateSpendEntry.Text)
		if errSpn != nil {
			dialog.ShowError(errSpn, win)
			return
		}
		if income != 0 && spend != 0 && utils.DatesCompare(dateInc, dateSpn) {
			dataBase.AddIncomeAndSpend(income, spend, dateInc, commentIncomEntry.Text, commentSpendEntry.Text)
		} else { // income and spend can be both no zero but with different dates
			if income != 0 {
				dataBase.AddIncome(income, dateInc, commentIncomEntry.Text)
			}
			if spend != 0 {
				dataBase.AddSpend(spend, dateSpn, commentSpendEntry.Text)
			}
		}

		// need to fix notifications (drivers or something)
		fyne.CurrentApp().SendNotification(fyne.NewNotification("Add success", "Income added"))

		// clearing entry fields
		incomeEntry.Text = ""
		dateIncomEntry.Text = ""
		spendEntry.Text = ""
		dateSpendEntry.Text = ""
		incomeEntry.Refresh()
		dateIncomEntry.Refresh()
		spendEntry.Refresh()
		dateSpendEntry.Refresh()
	})

	//	calendarBtn1 := CalendarBtn(dateIncomeBind, win)
	//	calendarBtn2 := CalendarBtn(dateSpendBind, win)

	return container.NewVBox(
		enterOperLabel,
		container.NewHBox(
			container.NewGridWithColumns(4), //				incomeEntry, dateIncomEntry, calendarBtn1, commentIncomEntry,
			//				spendEntry, dateSpendEntry, calendarBtn2, commentSpendEntry,

			addBtn),
	)
}
