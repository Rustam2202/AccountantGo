package utils

import (
	"errors"
	"testing"
	"time"
)

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
		if out, _ := CheckDate(test.input); !DatesCompare(out, test.expected) {
			t.Errorf("Output %q not equal to expected %q", out, test.expected)
		}
	}

	for _, test := range wrongFormatDatesTests {
		if _, err := CheckDate(test.input); err == nil {
			t.Errorf("Expected error %q", err)
		}
	}
}

func TestEntry(t *testing.T) {
	for _, test := range entries {
		sum, date, err := CheckEntry(test.inputSum, test.inputDate)
		if sum != test.expectedSum {
			t.Errorf("Expexted: %f, got: %f", test.expectedSum, sum)
		}
		if !DatesCompare(date, test.expectedDate) {
			t.Errorf("Expected: %q, got: %q", test.expectedDate, date)
		}
		/*
		if !errors.Is(err,test.err) {
			t.Errorf("Expected: %q, got: %q", test.err, err)
		}
		*/
	}
}


