package db

import (
	"accounter/utils"
	"math"
	"strconv"
	"testing"
	"time"
)

type addInput struct {
	sum  float32
	date time.Time
	err  error
}

var test1Inputs = []addInput{
	{sum: 0.001, date: time.Date(2023, time.January, 25, 0, 0, 0, 0, &time.Location{}), err: nil},     // 0
	{sum: 100.02, date: time.Date(2023, time.January, 26, 13, 59, 40, 0, &time.Location{}), err: nil}, // 1
	{sum: -2500.5, date: time.Date(2023, time.January, 27, 0, 0, 0, 0, &time.Location{}), err: nil},   // 2
	{sum: 2500.4, date: time.Date(2023, time.January, 27, 0, 0, 0, 0, &time.Location{}), err: nil},    // 3
	{sum: 10000, date: time.Date(2023, time.January, 28, 0, 0, 0, 0, &time.Location{}), err: nil},     // 4
	{sum: 1.1, date: time.Date(2023, time.January, 29, 0, 0, 1, 0, &time.Location{}), err: nil},       // 5
	{sum: -500, date: time.Date(2023, time.January, 29, 0, 0, 0, 0, &time.Location{}), err: nil},      // 6
	{sum: -0.03, date: time.Date(2023, time.January, 30, 0, 0, 0, 0, &time.Location{}), err: nil},     // 7
}

func TestCreateDB(t *testing.T) {
	var DataBase Database
	err := DataBase.CreateDataBase("create_db_test")
	if err != nil {
		t.Error(err)
	}
}

func Test1(t *testing.T) {
	var DataBase Database
	DataBase.CreateDataBase("test_1")
	DataBase.AddSpend(test1Inputs[0].sum, test1Inputs[0].date)
	DataBase.AddIncome(test1Inputs[1].sum, test1Inputs[1].date)
	DataBase.AddIncome(test1Inputs[2].sum, test1Inputs[2].date)
	DataBase.AddIncome(test1Inputs[3].sum, test1Inputs[3].date)
	DataBase.AddSpend(test1Inputs[4].sum, test1Inputs[4].date)
	DataBase.AddIncome(test1Inputs[5].sum, test1Inputs[5].date)
	DataBase.AddSpend(test1Inputs[6].sum, test1Inputs[6].date)
	DataBase.AddIncome(test1Inputs[7].sum, test1Inputs[7].date)
	result, err := DataBase.CalculateRecords(
		time.Date(2023, time.January, 26, 0, 0, 0, 0, &time.Location{}),    // 1...
		time.Date(2023, time.January, 29, 23, 59, 59, 0, &time.Location{}), // ...6
	)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(result); i++ {
		date, _ := time.Parse("02.01.2006", result[i][1])
		income, _ := strconv.ParseFloat(result[i][2], 32)
		spend, _ := strconv.ParseFloat(result[i][3], 32)
		if !(math.Abs(float64(test1Inputs[i+1].sum)-income-spend) <= 10e-3) {
			t.Errorf("Expected: %f, got: +%0.f -%f", test1Inputs[i].sum, income, spend)
		}
		if !utils.DatesCompare(test1Inputs[i+1].date, date) {
			t.Errorf("Expected: %q, got: %q", test1Inputs[i].date, date)
		}
	}
}
