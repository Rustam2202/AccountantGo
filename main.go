package main

import (
	db "accounter/db"
	gui "accounter/gui"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var CalendByte []byte
var DataBase db.Database

func main() {
	DataBase.Name = "test2"
	err := DataBase.OpenAndCreateLocalDb()
	fmt.Println(err)

	//	err := DataBase.OpenDataBase(DataBase.Name)
	//if err != nil {
	//	DataBase.CreateDataBase("tutelka")
	//}

	a := app.New()
	w := a.NewWindow("Accounter Manager")
	w.Resize(fyne.NewSize(900, 850))

	ContWithTable := container.NewWithoutLayout()

	//	dbCreate := fyne.NewMenuItem("New", func() {})
	//	dbOpen := fyne.NewMenuItem("Open", func() {})
	//	dbClose := fyne.NewMenuItem("Close", func() {})
	//	menuDb := fyne.NewMenu("Data base", dbCreate, dbOpen, dbClose)
	//main_menu := fyne.NewMainMenu(menuDb)

	addDbBtn := widget.NewButtonWithIcon("New data base", theme.ContentAddIcon(), func() {})
	addDbBtn.Alignment = widget.ButtonAlignCenter
	openDbBtn := widget.NewButtonWithIcon("Open data base", theme.FolderOpenIcon(), func() {})
	str := "data base name"
	strb := binding.BindString(&str)
	currentDb := widget.NewLabelWithData(strb)
	top := container.NewHBox(addDbBtn, openDbBtn, currentDb)
	bottom := widget.NewButton("Clear", func() {
		ContWithTable.RemoveAll()
	})

	w.SetContent(
		container.NewBorder(
			top, bottom, nil, nil,
			container.NewVBox(
				//	2,
				container.NewVBox(
					gui.AddOperation(&DataBase, w),
					gui.PeriodDates(ContWithTable, &DataBase, w),
				),
				ContWithTable,
			)),
	)
	//	w.SetMainMenu(main_menu)

	w.ShowAndRun()
}
