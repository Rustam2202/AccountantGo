package db

import (
	"accounter/utils"
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

const floatCompare = 10e-2

type addInput struct {
	sum     float32
	date    time.Time
	addType int
}

var test1Inputs = []addInput{
	{sum: 0.001, date: time.Date(2023, time.January, 25, 0, 0, 0, 0, &time.Location{})},     // 0
	{sum: 100.02, date: time.Date(2023, time.January, 26, 13, 59, 40, 0, &time.Location{})}, // 1
	{sum: -2500.5, date: time.Date(2023, time.January, 27, 0, 0, 0, 0, &time.Location{})},   // 2
	{sum: 2500.4, date: time.Date(2023, time.January, 27, 0, 0, 0, 0, &time.Location{})},    // 3
	{sum: 10000, date: time.Date(2023, time.January, 28, 0, 0, 0, 0, &time.Location{})},     // 4
	{sum: 1.1, date: time.Date(2023, time.January, 29, 0, 0, 1, 0, &time.Location{})},       // 5
	{sum: -500, date: time.Date(2023, time.January, 29, 0, 0, 0, 0, &time.Location{})},      // 6
	{sum: -0.03, date: time.Date(2023, time.January, 30, 0, 0, 0, 0, &time.Location{})},     // 7
}

func TestCreateAndDropSQLite(t *testing.T) {
	var DataBase Database
	DataBase.Name = "db"
	err := DataBase.OpenAndCreateLocalDb()
	if err != nil {
		t.Error(err)
	}
	err = DataBase.DropTable()
	if err != nil {
		t.Error(err)
	}
}

func Test1(t *testing.T) {
	var DataBase Database
	DataBase.Name = "test1"
	DataBase.OpenAndCreateLocalDb()
	DataBase.AddSpend(test1Inputs[0].sum, test1Inputs[0].date, "text")
	DataBase.AddIncome(test1Inputs[1].sum, test1Inputs[1].date, "loooooooooooooooooong text")
	DataBase.AddIncome(test1Inputs[2].sum, test1Inputs[2].date, "")
	DataBase.AddIncome(test1Inputs[3].sum, test1Inputs[3].date, "")
	DataBase.AddSpend(test1Inputs[4].sum, test1Inputs[4].date, "")
	DataBase.AddIncome(test1Inputs[5].sum, test1Inputs[5].date, "")
	DataBase.AddSpend(test1Inputs[6].sum, test1Inputs[6].date, "")
	DataBase.AddIncome(test1Inputs[7].sum, test1Inputs[7].date, "")
	result, err := DataBase.CalculateRecords(
		time.Date(2023, time.January, 26, 0, 0, 0, 0, &time.Location{}),    // 1...
		time.Date(2023, time.January, 29, 23, 59, 59, 0, &time.Location{}), // ...6
	)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(result.Data); i++ {
		date, _ := time.Parse("02.01.2006", result.Data[i][1])
		income, _ := strconv.ParseFloat(result.Data[i][2], 32)
		spend, _ := strconv.ParseFloat(result.Data[i][3], 32)
		if !(math.Abs(float64(test1Inputs[i+1].sum)-income-spend) <= 10e-3) {
			t.Errorf("Expected: %f, got: +%0.2f -%0.2f", test1Inputs[i].sum, income, spend)
		}
		if !utils.DatesCompare(test1Inputs[i+1].date, date) {
			t.Errorf("Expected: %q, got: %q", test1Inputs[i].date, date)
		}
	}
	DataBase.DropTable()
}

// Test2 checks SQLite local DB-file
func Test2(t *testing.T) {
	var localDb Database
	localDb.Name = "test2"

	err := localDb.OpenAndCreateLocalDb()
	if err != nil {
		t.Error(err)
	}

	inputs := makeRandomData()

	someComment := "Some comment for all"
	for _, d := range inputs {
		if d.addType == 0 {
			err := localDb.AddIncome(d.sum, d.date, someComment)
			if err != nil {
				t.Error(err)
			}
		} else if d.addType == 1 {
			err := localDb.AddSpend(d.sum, d.date, someComment)
			if err != nil {
				t.Error(err)
			}
		} else {
			err := localDb.AddIncomeAndSpend(d.sum, d.sum/2, d.date, someComment, someComment)
			if err != nil {
				t.Error(err)
			}
		}
	}

	outAll, err := localDb.CalculateRecords(time.Time{}, time.Now())
	if err != nil {
		t.Error(err)
	}
	for _, item := range outAll.Data {
		income, _ := strconv.ParseFloat(item[2], 32)
		spend, _ := strconv.ParseFloat(item[3], 32)
		id, _ := strconv.Atoi(item[0])
		id--
		if income != 0 && spend != 0 {
			sub := math.Abs((income + spend) - float64(inputs[id].sum+inputs[id].sum/2))
			if sub > floatCompare {
				t.Errorf("Expected: %0.3f, got: %0.3f", inputs[id].sum+inputs[id].sum/2, income+spend)
			}
		} else {
			sub := math.Abs((income + spend) - float64(inputs[id].sum))
			if sub > floatCompare {
				t.Errorf("Expected: %0.3f, got: %0.3f", inputs[id].sum, income+spend)
			}
		}
	}
	err = localDb.DropTable()
	if err != nil {
		t.Errorf("Data base wasn't drop")
	}
}

func makeRandomData() []addInput {
	var data [20]addInput
	var koef float64 = 200000.0 // determinate range from -100k to +100k with rand.Float32()
	minDate := time.Date(2020, 1, 1, 0, 0, 0, 0, &time.Location{}).Unix()
	maxDate := time.Now().Unix()
	delta := maxDate - minDate
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(data); i++ {
		data[i].sum = float32(math.Abs((rand.Float64() - 0.5) * koef))
		data[i].date = time.Unix(rand.Int63n(delta)+minDate, 0)
		data[i].addType = rand.Intn(3)
	}
	return data[:]
}
