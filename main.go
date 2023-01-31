package main

import (
	db "accounter/db"
	gui "accounter/gui"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var CalendByte []byte
var DataBase db.Database

func main() {

	DataBase.Name = "tutelka"
	err := DataBase.OpenDataBase(DataBase.Name)
	if err != nil {
		DataBase.CreateDataBase("tutelka")
	}

	// makeData()

	a := app.New()
	w := a.NewWindow("Accounter Manager")
	w.Resize(fyne.NewSize(900, 800))

	ContWithTable := container.NewWithoutLayout()
	bottom := widget.NewLabel("bottom")

	w.SetContent(
		container.NewVBox(
			//	2,
			container.NewVBox(
				gui.AddOperation(&DataBase, w),
				gui.PeriodDates(ContWithTable, &DataBase, w),
			),
			ContWithTable,
			bottom,
		))
	w.ShowAndRun()
}

type addInput struct {
	sum     float32
	date    time.Time
	addType int
}

func makeData() {
	var data [20]addInput
	var koef float32 = 200000.0 // determinate range from -100k to +100k with rand.Float32()
	minDate := time.Date(2020, 1, 1, 0, 0, 0, 0, &time.Location{}).Unix()
	maxDate := time.Now().Unix()
	delta := maxDate - minDate
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(data); i++ {
		data[i].sum = (rand.Float32() - 0.5) * koef
		data[i].date = time.Unix(rand.Int63n(delta)+minDate, 0)
		data[i].addType = rand.Intn(3)
	}
	someComment := "Some comment for all"
	for _, d := range data {
		if d.addType == 0 {
			DataBase.AddIncome(d.sum, d.date, someComment)
		} else if d.addType == 1 {
			DataBase.AddSpend(d.sum, d.date, someComment)
		} else {
			DataBase.AddIncomeAndSpend(d.sum, (rand.Float32()-0.5)*koef, d.date, someComment, someComment)
		}
	}
}
