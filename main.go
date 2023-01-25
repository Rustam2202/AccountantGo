package main

import (
	acc "accounter/app"
	gui "accounter/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Accounter Manager")
	w.Resize(fyne.NewSize(400, 800))

	//	c:=gui.CalendarBtn(w)
	//cc:=gui.Calendar()

	w.SetContent(
		container.NewVBox(
			widget.NewLabel(""),
			gui.CalendarBtn(w),
			//&cc,
		),
	)
	w.ShowAndRun()
	acc.Calc()
	//gui.Dummy()
}
