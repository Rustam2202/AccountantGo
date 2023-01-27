package gui

import (
	// db "accounter/db"
	"accounter/db"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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
		inc, _ := strconv.ParseFloat(incomeEntry.Text, 32)
		dataBase.AddIncome(float32(inc), dateIncomEntry.Text)
	})
	subBtn := widget.NewButton("Sub", func() {
		spn, _ := strconv.ParseFloat(spendEntry.Text, 32)
		dataBase.AddSpend(float32(spn), dateSpendEntry.Text)
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
