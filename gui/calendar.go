package gui

import (
	"bufio"
	"io/ioutil"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

var i = widget.NewLabel("Choose a date")
var l = widget.NewLabel("")
var d = &date{instruction: i, dateChosen: l}
var showDateFormat = "02 Jan 2006 Mon"

type date struct {
	instruction *widget.Label
	dateChosen  *widget.Label
}

func (d *date) onSelected(t time.Time) {
	d.instruction.SetText("Date Selected:")
	d.dateChosen.SetText(t.Format(showDateFormat))
}

func calendar() *fyne.Container {
	i.Alignment = fyne.TextAlignCenter
	l.Alignment = fyne.TextAlignCenter
	startingDate := time.Now()
	cal := xwidget.NewCalendar(startingDate, d.onSelected)
	return container.NewVBox(i, l, cal)
}

func CalendarBtn(date binding.String, win fyne.Window) *fyne.Container {
	icon := calenIcon()

	c := container.NewVBox(
		widget.NewButtonWithIcon("Calendar", icon, func() {
			dialog.NewCustomConfirm(
				"Choose a date",
				"OK",
				"Cancel",
				calendar(), func(b bool) {
					dateToEntry, _ := time.Parse(showDateFormat, d.dateChosen.Text)
					date.Set(dateToEntry.Format("02.01.2006"))
				},
				win,
			).Show()
		}),
	)
	return c
}

func calenIcon() *fyne.StaticResource {
	calendFile, err := os.Open("calendar.png")
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(calendFile)
	CalendByte, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return fyne.NewStaticResource("icon", CalendByte)
}
