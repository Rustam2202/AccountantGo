package db

import "testing"

func TestCreateDB(t *testing.T) {
	var DataBase Database
	DataBase.CreateDataBase("createtest")
}

func Test1(t *testing.T) {
	var DataBase Database
	DataBase.CreateDataBase("test1")
	DataBase.AddIncome(999.99, "1998-01-23 12:45:56")
	DataBase.AddIncome(570.75, "2000-02-28 00:01:59")
	DataBase.AddIncome(1500, "2005-04-25 00:00:00")
	DataBase.AddSpend(2000.02, "2010-12-31 23:59:59")
	DataBase.AddSpend(3200, "2023-01-26 15:16:00")
	DataBase.CalculateRecords("Jan 20 2023", "Jan 26 2023")
}
