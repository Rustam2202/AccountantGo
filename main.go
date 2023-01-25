package main

import (
	acc "accounter/app"
	gui "accounter/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Accounter Manager")
	w.Resize(fyne.NewSize(600, 600))

	//	c:=gui.CalendarBtn(w)
	//cc:=gui.Calendar()

	w.SetContent(
		container.NewVBox(
			gui.Operation(w),
			gui.PeriodDates(w),
		),
	)

	w.ShowAndRun()
	acc.Calc()
	//gui.Dummy()
}
