package main

import (
	acc "accounter/app"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Accounter Manager")

	enterOperText := widget.NewLabel("Enter operation:")
	//content := container.New(layout.NewHBoxLayout(), enterOperText)

	emptyLabel := widget.NewLabel("")
	incomeLabel := widget.NewLabel("Income")
	spendLabel := widget.NewLabel("Spend")
	sumLabel := widget.NewLabel("Sum")
	dateLabel := widget.NewLabel("Date")

	incomeEntry := widget.NewEntry()
	spendEntry := widget.NewEntry()

	addBtn := widget.NewButton("Add", func() {
		// income := incomeEntry.Text
	})
	subBtn:=widget.NewButton("Sub",func() {
	})

	w.SetContent(
		container.NewVBox(
			enterOperText,
			container.NewGridWithColumns(4,
				emptyLabel, sumLabel, dateLabel, emptyLabel,
				incomeLabel, incomeEntry, emptyLabel, addBtn,
				spendLabel, spendEntry, emptyLabel, subBtn,
			),
		),
	)
	w.ShowAndRun()
	acc.Calc()
}
