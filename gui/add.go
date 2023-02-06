package gui

import (
	"accountant/utils"
	"errors"

	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

func (acc *accounter) makeAddBlock() *fyne.Container {
	today := time.Now().Format("02.01.2006")
	return container.NewVBox(
		acc.makeLabel("Enter income and/or spend to add in database", 1), // header
		container.NewHBox(
			container.NewGridWithColumns(2,
				acc.makeEntry(acc.IncomeEntry, "Enter Income"),
				acc.makeEntryWithData(acc.dateIncomeBind, acc.dateIncomEntry, today),
				acc.makeEntry(acc.spendEntry, "Enter Spend"),
				acc.makeEntryWithData(acc.dateSpendBind, acc.dateSpendEntry, today),
			),
			container.NewGridWithColumns(1,
				CalendarBtn(acc.dateIncomeBind, acc.win),
				CalendarBtn(acc.dateSpendBind, acc.win),
			),
			container.NewGridWithColumns(1,
				acc.makeEntry(acc.commentIncomEntry, "Enter comment\t\t\t\t"), // '\t' make more width for entry
				acc.makeEntry(acc.commentSpendEntry, "Enter comment\t\t\t\t"),
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
	acc.commentIncomEntry.Text = ""
	acc.commentSpendEntry.Text = ""
	acc.IncomeEntry.Refresh()
	acc.dateIncomEntry.Refresh()
	acc.spendEntry.Refresh()
	acc.dateSpendEntry.Refresh()
	acc.commentIncomEntry.Refresh()
	acc.commentSpendEntry.Refresh()
}
