package utils

import (
	"errors"
	"testing"
	"time"
)

const dateLayout = "02.01.2006"

type dateTest struct {
	input    string
	expected time.Time
}

type entryTest struct {
	inputSum     string
	inputDate    string
	expectedSum  float32
	expectedDate time.Time
	err          error
}

var somedate = time.Date(2023, time.January, 30, 0, 0, 0, 0, &time.Location{})

var correctDatesTests = []dateTest{
	{input: "30.01.2023", expected: somedate},
	{input: "30-01-2023", expected: somedate},
	{input: "30/01/23", expected: somedate},
	{input: "30-01-23", expected: somedate},
}

var wrongFormatDatesTests = []dateTest{
	{input: "symbols"},
	{input: "00"},
	{input: "00.00.0000"},
	{input: "32.01.2023"},
	{input: "29.02.2023"},
}

var entries = []entryTest{
	{inputSum: "1000.02", inputDate: "30-01-23", expectedSum: 1000.02, expectedDate: somedate, err: nil},
	{inputSum: "-200", inputDate: "30-01-23", expectedSum: 200, expectedDate: somedate, err: nil},
	{inputSum: "not numbers", inputDate: "30-01-23", expectedSum: 0, expectedDate: time.Time{}, err: errors.New("Sum format error")},
}

func TestDate(t *testing.T) {
	for _, test := range correctDatesTests {
		if out, _ := CheckDateFormat(test.input); !IsEqualDates(out, test.expected) {
			t.Errorf("Output %q not equal to expected %q", out, test.expected)
		}
	}

	for _, test := range wrongFormatDatesTests {
		if _, err := CheckDateFormat(test.input); err == nil {
			t.Errorf("Expected error %q", err)
		}
	}
}

func TestIsEqualDates(t *testing.T) {
	date1 := time.Date(2023, 02, 02, 0, 0, 0, 0, &time.Location{})
	date2 := time.Date(2023, 02, 02, 0, 0, 0, 0, &time.Location{})
	date3 := time.Date(2023, 03, 02, 0, 0, 0, 0, &time.Location{})
	date4 := time.Date(2023, 02, 01, 0, 0, 0, 0, &time.Location{})
	if !IsEqualDates(date1, date2) {
		t.Errorf("%s and %s expected true", date1.String(), date2.String())
	}
	if IsEqualDates(date1, date3) {
		t.Errorf("%s and %s expected false", date1.String(), date3.String())
	}
	if IsEqualDates(date1, date4) {
		t.Errorf("%s and %s expected false", date1.String(), date4.String())
	}
}

func TestCheckDateRange(t *testing.T) {

	yearsAgoExactly := time.Now().AddDate(-YearsAgoMin, 0, 0)
	yearsAgoAndDay := time.Now().AddDate(-YearsAgoMin, 0, -1)
	yearsAgoMinusOne := time.Now().AddDate(-YearsAgoMin+1, 0, 0)
	todayPlusDay := time.Now().AddDate(0, 0, 1)
	today := time.Now()

	if err := checkDateRange(yearsAgoExactly); err != nil {
		t.Errorf("%s is exactly %d years ago and must be in range", yearsAgoExactly.Format(dateLayout), YearsAgoMin)
	}
	if err := checkDateRange(yearsAgoAndDay); err == nil {
		t.Errorf("%s expected as less then %d years ago and must out range", yearsAgoAndDay.Format(dateLayout), YearsAgoMin)
	}
	if err := checkDateRange(yearsAgoMinusOne); err != nil {
		t.Errorf("%s expected as great then %d years ago and must be in range", yearsAgoMinusOne.Format(dateLayout), YearsAgoMin)
	}
	if err := checkDateRange(todayPlusDay); err == nil {
		t.Errorf("%s expected as great then today ago and must out range", todayPlusDay.Format(dateLayout))
	}
	if err := checkDateRange(today); err != nil {
		t.Errorf("%s expected as today and must be in range", today.Format(dateLayout))
	}
}

func TestEntry(t *testing.T) {
	for _, test := range entries {
		sum, date, _ := CheckEntry(test.inputSum, test.inputDate)
		if sum != test.expectedSum {
			t.Errorf("Expexted: %f, got: %f", test.expectedSum, sum)
		}
		if !IsEqualDates(date, test.expectedDate) {
			t.Errorf("Expected: %q, got: %q", test.expectedDate, date)
		}
		/*
			if !errors.Is(err,test.err) {
				t.Errorf("Expected: %q, got: %q", test.err, err)
			}
		*/
	}
}
