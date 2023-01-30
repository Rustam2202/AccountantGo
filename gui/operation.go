package gui

import (
	// db "accounter/db"
	"accounter/db"
	"errors"
	"math"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// allowed manual input date formats (dd.mm.yyyy)
const (
	format1 = "02.01.2006"
	format2 = "02/01/2006"
	format3 = "02-01-2006"
	format4 = "02.01.06"
	format5 = "02/01/06"
	format6 = "02-01-06"
	format7 = "02,01,2006"
	format8 = "02,01,06"
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
	incomeEntry.SetPlaceHolder("Enter Income")
	spendEntry := widget.NewEntry()
	spendEntry.SetPlaceHolder("Enter Spend")
	dateIncomEntry := widget.NewEntryWithData(dateIncomeBind)
	dateIncomEntry.SetPlaceHolder(format1)
	dateSpendEntry := widget.NewEntryWithData(dateSpendBind)
	dateSpendEntry.SetPlaceHolder(format1)
	commentIncomEntry := widget.NewEntry()
	commentSpendEntry := widget.NewEntry()

	addBtn := widget.NewButton("Add record", func() {
		if !(incomeEntry.Text != "" || spendEntry.Text != "") {
			dialog.ShowError(errors.New("Income or Spend field must contain a value"), win)
			return
		}

		var income, spend float32
		var dateInc, dateSpn time.Time
		var errInc, errSpn error
		if incomeEntry.Text != "" {
			income, dateInc, errInc = checkEntry(incomeEntry.Text, dateIncomEntry.Text)
			if errInc != nil {
				dialog.ShowError(errInc, win)
				return
			}
			dataBase.AddIncome(income, dateInc)
		}
		if spendEntry.Text != "" {
			spend, dateSpn, errSpn = checkEntry(spendEntry.Text, dateSpendEntry.Text)
			if errSpn != nil {
				dialog.ShowError(errSpn, win)
				return
			}
			dataBase.AddSpend(spend, dateSpn)
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

func checkEntry(sumStr string, dateStr string) (float32, time.Time, error) {

	sum, err := strconv.ParseFloat(sumStr, 32)
	if err != nil {
		return 0, time.Time{}, errors.New("Sum format error")
	}

	var date time.Time
	if dateStr == "" {
		date = time.Now() // if no manual or calendar input then set today
	} else {
		temp, err2 := checkDate(dateStr)
		if err2 != nil {
			return 0, date, err2
		} else {
			date = temp
		}
	}
	return float32(math.Abs(sum)), date, nil

}

func checkDate(date string) (time.Time, error) {

	var t time.Time
	var err error

	// try to change on switch
	if t, err = time.Parse(format1, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format1, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format2, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format3, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format4, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format5, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format6, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format7, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format8, date); err == nil {
		return t, nil
	} else {
		return t, err
	}
}
