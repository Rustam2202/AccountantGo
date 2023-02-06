package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func instructionBtn() *widget.Button {
	return widget.NewButtonWithIcon("", theme.QuestionIcon(), func() {
		w := fyne.CurrentApp().NewWindow("Accountant instruction")
		w.SetContent(widget.NewTextGridFromString(instruction))
		w.Show()
	})
}

var instruction = `Hi!
Welcome to your accountant application!

In directory where placed 'accountant.exe' file must be also 'database.db' file. 
Dispose them together,or a new database will be created with launching an application, 
thus your old data can be lost.
In 'database.db' stored all your records (No., Date, Income or/and Spend, Comments). 

To create some record in database just enter Income or Spend sum and push add. 
Thus it's adding your sums for today. You can add some comment to know aboute 
source of income or expense.Also you can choose or enter date, but not later then 
5 years ago and not older then today. Again, if date entryis empty, then it is implied as today. 
You can enter the in ome of the valid formats:
'21.12.2012'
'21.12.12'
'21,12,2012'
'21,12,12'
'21/12/2012'
'21/12/12'
'21-12-2012'
'21-12-12'

To view your database, you can use:
- 'Show all' button that show all records in database;
- Choose some period of time by entering 'Period begin' and 'Period end' dates and push 
	'Show period' (rules by entering dates is same as in Add);
- Select 'Month and year' then push 'Show monthly' report;
- Selct or enter 'Year' to annual report.

In the application window will be displayed suma of all income, all expenses and the 
difference for the selected period.
Below is a table with details. You can delete any row from table by entering it's 'No.' 
and push 'Delete'. 'Clear' button cleare the output results.

The application can show you some errors as incorrect sums input (must be only decimal value), 
out of date range (less the 5 years ago and in future).
`
