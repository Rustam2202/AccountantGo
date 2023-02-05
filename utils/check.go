package utils

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

// allowed manual input date formats (dd.mm.yyyy)
var dateInputFormats = []string{"02.01.06", "02,01,06", "02/01/06", "02-01-06", "02.01.2006", "02,01,2006", "02/01/2006", "02-01-2006"}
var YearsAgoMin = 5

func CheckEntry(sumStr string, dateStr string) (float32, time.Time, error) {
	if sumStr == "" {
		return 0, time.Time{}, nil
	}
	sum, err := strconv.ParseFloat(sumStr, 32)
	if err != nil {
		return 0, time.Time{}, errors.New(" Sum format error")
	}

	var date time.Time
	if dateStr == "" {
		date = time.Now() // if no manual or calendar input, then set today
	} else {
		temp, err2 := CheckDateFormat(dateStr)
		if err2 != nil {
			return 0, date, err2
		} else {
			date = temp
		}
	}
	return float32(math.Abs(sum)), date, nil
}

func CheckDateFormat(date string) (time.Time, error) {
	if date == "" {
		return time.Now(), nil
	}

	for _, format := range dateInputFormats {
		if t, err := time.Parse(format, date); err == nil {
			if err := checkDateRange(t); err != nil {
				return time.Time{}, err
			}
			return t, nil
		}
	}
	return time.Time{}, errors.New(" Date format error")
}

func checkDateRange(date time.Time) error {
	// minimum of date range if 1 year and 6 month ago from today
	if date.Before(time.Now().AddDate(-YearsAgoMin, 0, 0)) {
		return fmt.Errorf(" Date older than %d years ago", YearsAgoMin)
	}
	// maximum range can't be in future
	if date.After(time.Now()) {
		return fmt.Errorf(" Date greater then today")
	}
	return nil
}

func IsEqualDates(first, second time.Time) bool {
	if first.After(second) || first.Before(second) {
		return false
	}
	// if aboth false then them equal => true
	return !(first.After(second) && first.Before(second))
}
