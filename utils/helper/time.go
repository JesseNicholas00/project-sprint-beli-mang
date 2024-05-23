package helper

import "time"

func MustParse(timeString string) time.Time {
	res, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		panic(err)
	}
	return res
}

func MustParseDateOnly(timeString string) time.Time {
	res, err := time.Parse(time.DateOnly, timeString)
	if err != nil {
		panic(err)
	}
	return res
}
