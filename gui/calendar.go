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
var dateChoosen string

type date struct {
	instruction *widget.Label
	dateChosen  *widget.Label
}

func (d *date) onSelected(t time.Time) {
	d.instruction.SetText("Date Selected:")
	d.dateChosen.SetText(t.Format("Mon 2 Jan 2006"))
}

func calendar() *fyne.Container /* *xwidget.Calendar */ {

	i.Alignment = fyne.TextAlignCenter

	l.Alignment = fyne.TextAlignCenter

	startingDate := time.Now()
	cal := xwidget.NewCalendar(startingDate, d.onSelected)
	return container.NewVBox(i, l, cal)
}

func CalendarBtn(date binding.String, win fyne.Window) *fyne.Container {
	calendFile, err := os.Open("calendar.png")
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(calendFile)

	CalendByte, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	icon := fyne.NewStaticResource("icon", CalendByte)
	c := container.NewVBox(
		//	widget.NewButtonWithIcon("", theme.GridIcon(), func() {
		widget.NewButtonWithIcon("Calendar", icon, func() {
			dialog.NewCustomConfirm(
				"Choose a date",
				"OK",
				"Cancel",
				calendar(), func(b bool) {
					date.Set(d.dateChosen.Text)
				},
				win,
			).Show()
		}),
		//dateLabel,
		//entryDate,
		//widget.NewLabel("Label"),
	)
	return c
}
