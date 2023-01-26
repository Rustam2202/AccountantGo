package db

import "testing"

func Test1(t *testing.T) {
	var DataBase Database
	DataBase.CreateDataBase("test1")
	DataBase.AddIncome("10200", "Jan 20 2023")
	DataBase.AddIncome("570", "Jan 22 2023")
	DataBase.AddIncome("1500", "Jan 25 2023")
	DataBase.AddSpend("2000", "Jan 25 2023")
	DataBase.AddSpend("3200", "jan 26 2023")
	DataBase.ShowRecords("Jan 20 2023", "Jan 26 2023")
}
