package gui

import (
	// db "accounter/db"
	"accounter/db"
	"errors"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func AddOperation(dataBase *db.Database, win fyne.Window) *fyne.Container {
	enterOperLabel := widget.NewLabel("Enter operation:")
	enterOperLabel.Alignment = fyne.TextAlignCenter

	emptyLabel := widget.NewLabel("")
	incomeLabel := widget.NewLabel("Income")
	spendLabel := widget.NewLabel("Spend")
	sumLabel := widget.NewLabel("Sum")
	dateLabel := widget.NewLabel("Date")
	commentLabel := widget.NewLabel("Comment")

	dateIncomeBind := binding.BindString(nil)
	dateSpendBind := binding.BindString(nil)

	incomeEntry := widget.NewEntry()
	spendEntry := widget.NewEntry()
	dateIncomEntry := widget.NewEntryWithData(dateIncomeBind)
	dateIncomEntry.SetPlaceHolder("01/01/2001")
	dateSpendEntry := widget.NewEntryWithData(dateSpendBind)
	dateSpendEntry.SetPlaceHolder("01/01/2001")
	commentIncomEntry := widget.NewEntry()
	commentSpendEntry := widget.NewEntry()

	addBtn := widget.NewButton("Add", func() {
		dataBase.AddIncome(checkEntry(incomeEntry.Text, dateIncomEntry.Text, win))
		// Need uiniq ID or toml for Notification
		fyne.CurrentApp().SendNotification(fyne.NewNotification("Add success", "Income added"))
		// clear entry fields
		incomeEntry.Text = ""
		dateIncomEntry.Text = ""
		incomeEntry.Refresh()
		dateIncomEntry.Refresh()
	})
	subBtn := widget.NewButton("Sub", func() {
		dataBase.AddSpend(checkEntry(spendEntry.Text, dateSpendEntry.Text, win))
		// Need uiniq ID or toml for Notification
		fyne.CurrentApp().SendNotification(fyne.NewNotification("Sub success", "Spend added"))
		// clear entry fields
		spendEntry.Text = ""
		dateSpendEntry.Text = ""
		spendEntry.Refresh()
		dateSpendEntry.Refresh()
	})

	calendarBtn1 := CalendarBtn(dateIncomeBind, win)
	calendarBtn2 := CalendarBtn(dateSpendBind, win)

	c := container.NewVBox(
		enterOperLabel,
		container.NewGridWithColumns(6,
			emptyLabel, sumLabel, dateLabel, emptyLabel, commentLabel, emptyLabel,
			incomeLabel, incomeEntry, dateIncomEntry, calendarBtn1, commentIncomEntry, addBtn,
			spendLabel, spendEntry, dateSpendEntry, calendarBtn2, commentSpendEntry, subBtn,
		),
	)
	return c
}

func checkEntry(income string, date string, win fyne.Window) (float32, string) {
	var inc float64
	var dat string
	var err error
	if income != "" {
		inc, err = strconv.ParseFloat(income, 32)
		if err != nil {
			dialog.ShowError(errors.New("Income format error"), win)
		}
		// add regexpr for different date format input (31.01.2001 31/01/2001 31-01-2001 31-01-01)
		if date == "" {
			d := time.Now()
			dat = d.Format("02.01.2006")
		} else {
			d, err2 := time.Parse("02.01.2006", date)
			if err2 != nil {
				dialog.ShowError(errors.New("Date format error"), win)
			}
			dat = d.Format("02.01.2006")
		}
	} else {
		dialog.ShowError(errors.New("Income must contain a value"), win)
	}
	return float32(inc), dat
}
