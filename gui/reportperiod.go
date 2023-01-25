package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func PeriodDates(win fyne.Window) *fyne.Container {

	empty := widget.NewLabel("")
	label := widget.NewLabel("Enter period or month for report")
	fromLabel := widget.NewLabel("From")
	toLabel := widget.NewLabel("To")
	fromBtn := CalendarBtn(win)
	toBtn := CalendarBtn(win)
	fromEntry := widget.NewEntry()
	toEntry := widget.NewEntry()
	monthLabel := widget.NewLabel("Month:")
	monthLabel.Alignment = fyne.TextAlignTrailing
	yearLabel := widget.NewLabel("Year:")
	yearLabel.Alignment = fyne.TextAlignTrailing
	monthEntry := widget.NewSelect([]string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sep", "Oct", "Nov,", "Dec"}, func(s string) {})
	yearEntry := widget.NewSelectEntry([]string{"2020", "2021", "2022", "2023"})

	c := container.NewVBox(
		label,
		container.NewGridWithColumns(4,
			fromLabel, fromEntry, fromBtn, empty,
			toLabel, toEntry, toBtn, empty,
			monthLabel, monthEntry, yearLabel, yearEntry,
		),
	)
	return c
}
