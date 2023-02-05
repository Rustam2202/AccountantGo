package gui

import (
	"accounter/utils"
	"errors"

	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

func (acc *accounter) makeAddBlock() *fyne.Container {
	return container.NewVBox(
		acc.makeLabel("Enter income and/or spend to add in database", 1), // header
		container.NewHBox(
			container.NewGridWithColumns(4,
				acc.makeEntry(acc.IncomeEntry, "Enter Income:"),
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
	if !(acc.IncomeEntry.Text != "" || acc.spendEntry.Text != "") {
		dialog.ShowError(errors.New(" Income or Spend field must contain a value"), acc.win)
		return
	}

	var income, spend float32
	var dateInc, dateSpn time.Time
	var errInc, errSpn error
	income, dateInc, errInc = utils.CheckEntry(acc.IncomeEntry.Text, acc.dateIncomEntry.Text)
	if errInc != nil {
		dialog.ShowError(errInc, acc.win)
		return
	}
	spend, dateSpn, errSpn = utils.CheckEntry(acc.spendEntry.Text, acc.dateSpendEntry.Text)
	if errSpn != nil {
		dialog.ShowError(errSpn, acc.win)
		return
	}
	if income != 0 && spend != 0 && utils.IsEqualDates(dateInc, dateSpn) {
		acc.dataBase.AddIncomeAndSpend(income, spend, dateInc, acc.commentIncomEntry.Text, acc.commentSpendEntry.Text)
		// income and spend can be both no zero but with different dates
	} else {
		if income != 0 {
			acc.dataBase.AddIncome(income, dateInc, acc.commentIncomEntry.Text)
		}
		if spend != 0 {
			acc.dataBase.AddSpend(spend, dateSpn, acc.commentSpendEntry.Text)
		}
	}

	// need to fix notifications (drivers or something)
	// fyne.CurrentApp().SendNotification(fyne.NewNotification("Add success", "Income added"))

	acc.clearAddEntries()
}

func (acc *accounter) clearAddEntries() {
	acc.IncomeEntry.Text = ""
	acc.dateIncomEntry.Text = ""
	acc.spendEntry.Text = ""
	acc.dateSpendEntry.Text = ""
	acc.IncomeEntry.Refresh()
	acc.dateIncomEntry.Refresh()
	acc.spendEntry.Refresh()
	acc.dateSpendEntry.Refresh()
}
