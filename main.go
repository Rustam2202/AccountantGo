package main

import (
	acc "accounter/app"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Accounter Manager")

	enterOperText := widget.NewLabel("Enter operation:")
	enterOperText.Alignment = fyne.TextAlignCenter
	//content := container.New(layout.NewHBoxLayout(), enterOperText)

	emptyLabel := widget.NewLabel("")
	incomeLabel := widget.NewLabel("Income")
	spendLabel := widget.NewLabel("Spend")
	sumLabel := widget.NewLabel("Sum")
	dateLabel := widget.NewLabel("Date")

	incomeEntry := widget.NewEntry()
	spendEntry := widget.NewEntry()
	dateIncomEntry := widget.NewEntry()
	dateSpendEntry := widget.NewEntry()

	addBtn := widget.NewButton("Add", func() {
		// income := incomeEntry.Text
	})
	subBtn := widget.NewButton("Sub", func() {
	})

	//calendar := calendar()
	calendarBtn1 := calendarBtn()
	calendarBtn2 := calendarBtn()

	w.SetContent(
		container.NewVBox(
			enterOperText,
			container.NewGridWithColumns(5,
				emptyLabel, sumLabel, dateLabel, emptyLabel, emptyLabel,
				incomeLabel, incomeEntry, dateIncomEntry, &calendarBtn1, addBtn,
				spendLabel, spendEntry, dateSpendEntry, &calendarBtn2, subBtn,
			),
		),
	)
	w.ShowAndRun()
	acc.Calc()
}

type date struct {
	instruction *widget.Label
	dateChosen  *widget.Label
}

func (d *date) onSelected(t time.Time) {
	d.instruction.SetText("Date Selected:")
	d.dateChosen.SetText(t.Format("Mon 2 Jan 2006"))
}

func calendar() fyne.Container {
	i := widget.NewLabel("Choose a date")
	i.Alignment = fyne.TextAlignCenter
	l := widget.NewLabel("")
	l.Alignment = fyne.TextAlignCenter
	d := &date{instruction: i, dateChosen: l}
	startingDate := time.Now()
	calendar := xwidget.NewCalendar(startingDate, d.onSelected)
	c := container.NewVBox(i, l, calendar)
	return *c
}

func calendarBtn() fyne.Container {
	c := container.NewVBox(
		widget.NewButtonWithIcon("", theme.GridIcon(), func() {
			w := fyne.CurrentApp().NewWindow("calendar")
			cal := calendar()
			w.SetContent(&cal)
			w.Show()
		}),
	)

	return *c
}
