package utils

import (
	"errors"
	"math"
	"strconv"
	"time"
)

// allowed manual input date formats (dd.mm.yyyy)
const (
	Format1 = "02.01.2006"
	format2 = "02/01/2006"
	format3 = "02-01-2006"
	format4 = "02.01.06"
	format5 = "02/01/06"
	format6 = "02-01-06"
	format7 = "02,01,2006"
	format8 = "02,01,06"
)

func CheckEntry(sumStr string, dateStr string) (float32, time.Time, error) {

	sum, err := strconv.ParseFloat(sumStr, 32)
	if err != nil {
		return 0, time.Time{}, errors.New("Sum format error")
	}

	var date time.Time
	if dateStr == "" {
		date = time.Now() // if no manual or calendar input then set today
	} else {
		temp, err2 := CheckDate(dateStr)
		if err2 != nil {
			return 0, date, err2
		} else {
			date = temp
		}
	}
	return float32(math.Abs(sum)), date, nil

}

func CheckDate(date string) (time.Time, error) {
	if date == "" {
		return time.Now(), nil
	}
	var t time.Time
	var err error

	// try change on switch
	if t, err = time.Parse(Format1, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(Format1, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format2, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format3, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format4, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format5, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format6, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format7, date); err == nil {
		return t, nil
	} else if t, err = time.Parse(format8, date); err == nil {
		return t, nil
	} else {
		return t, err
	}
}

func DatesCompare(first, second time.Time) bool {
	// if aboth false then them equal => true
	return !(first.After(second) && first.Before(second))
}
