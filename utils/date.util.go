package utils

import "time"

func DateUtilYYYYMMDD(date string) (*time.Time, error) {
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		return nil, err
	}
	return &parsedTime, nil
}
