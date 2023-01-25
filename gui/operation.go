package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var dateBind = binding.BindString(nil)

func Operation(win fyne.Window) *fyne.Container {
	enterOperLabel := widget.NewLabel("Enter operation:")
	enterOperLabel.Alignment = fyne.TextAlignCenter

	emptyLabel := widget.NewLabel("")
	incomeLabel := widget.NewLabel("Income")
	spendLabel := widget.NewLabel("Spend")
	sumLabel := widget.NewLabel("Sum")
	dateLabel := widget.NewLabel("Date")

	incomeEntry := widget.NewEntry()
	spendEntry := widget.NewEntry()
	dateIncomEntry := widget.NewEntryWithData(dateBind)
	dateIncomEntry.SetPlaceHolder("01/01/2001")
	dateSpendEntry := widget.NewEntry()

	addBtn := widget.NewButton("Add", func() {
		// income := incomeEntry.Text
	})
	subBtn := widget.NewButton("Sub", func() {
	})

	//calendar := calendar()
	calendarBtn1 := CalendarBtn(win)
	//calendarBtn2 := CalendarBtn(&date_str, win)
	calendarBtn1.Resize(fyne.NewSize(5,5))
	

	c := container.NewVBox(
		enterOperLabel,
		container.NewGridWithColumns(5,
			emptyLabel, sumLabel, dateLabel, emptyLabel, emptyLabel,
			incomeLabel, incomeEntry, dateIncomEntry, calendarBtn1, addBtn,
			spendLabel, spendEntry, dateSpendEntry, calendarBtn1, subBtn,
		),
	)

	return c

}
