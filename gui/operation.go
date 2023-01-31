package gui

import (
	// db "accounter/db"
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

func AddOperation(dataBase *db.Database, win fyne.Window) *fyne.Container {
	enterOperLabel := widget.NewLabel("Enter income and/or spend to add in database")
	enterOperLabel.Alignment = fyne.TextAlignCenter

	/*
		emptyLabel := widget.NewLabel("")
		incomeLabel := widget.NewLabel("Income")
		spendLabel := widget.NewLabel("Spend")
		sumLabel := widget.NewLabel("Sum")
		dateLabel := widget.NewLabel("Date")
		commentLabel := widget.NewLabel("Comment")
	*/

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
			dataBase.AddIncomeAndSpend(income, spend, dateInc)
		} else { // income and spend can be both no zero but with different dates
			if income != 0 {
				dataBase.AddIncome(income, dateInc)
			}
			if spend != 0 {
				dataBase.AddSpend(spend, dateSpn)
			}
		}
		/*
			if incomeEntry.Text != "" {
				income, dateInc, errInc = utils.CheckEntry(incomeEntry.Text, dateIncomEntry.Text)
				if errInc != nil {
					dialog.ShowError(errInc, win)
					return
				}
				dataBase.AddIncome(income, dateInc)
			}
			if spendEntry.Text != "" {
				spend, dateSpn, errSpn = utils.CheckEntry(spendEntry.Text, dateSpendEntry.Text)
				if errSpn != nil {
					dialog.ShowError(errSpn, win)
					return
				}
				dataBase.AddSpend(spend, dateSpn)
			}
		*/

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

	calendarBtn1 := CalendarBtn(dateIncomeBind, win)
	calendarBtn2 := CalendarBtn(dateSpendBind, win)

	return container.NewVBox(
		enterOperLabel,
		container.NewHBox(
			container.NewGridWithColumns(4,
				incomeEntry, dateIncomEntry, calendarBtn1, commentIncomEntry,
				spendEntry, dateSpendEntry, calendarBtn2, commentSpendEntry,
			),
			addBtn),
	)
}
