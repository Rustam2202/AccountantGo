package gui

import (
	// db "accounter/db"
	"accounter/db"

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

	dateIncomeBind := binding.BindString(nil)
	dateSpendBind := binding.BindString(nil)

	incomeEntry := widget.NewEntry()
	spendEntry := widget.NewEntry()
	dateIncomEntry := widget.NewEntryWithData(dateIncomeBind)
	dateIncomEntry.SetPlaceHolder("01/01/2001")
	dateSpendEntry := widget.NewEntryWithData(dateSpendBind)
	dateSpendEntry.SetPlaceHolder("01/01/2001")

	addBtn := widget.NewButton("Add", func() {
		dataBase.AddIncome(incomeEntry.Text, dateIncomEntry.Text)
	})
	subBtn := widget.NewButton("Sub", func() {
		dataBase.AddSpend(spendEntry.Text, dateSpendEntry.Text)
	})

	calendarBtn1 := CalendarBtn(dateIncomeBind, win)
	calendarBtn2 := CalendarBtn(dateSpendBind, win)

	c := container.NewVBox(
		enterOperLabel,
		container.NewGridWithColumns(5,
			emptyLabel, sumLabel, dateLabel, emptyLabel, emptyLabel,
			incomeLabel, incomeEntry, dateIncomEntry, calendarBtn1, addBtn,
			spendLabel, spendEntry, dateSpendEntry, calendarBtn2, subBtn,
		),
	)

	return c

}
