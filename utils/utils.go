package utils

import (
	"fmt"
	"time"
)

func ConvertDateStringToTime(dateString string) (time.Time, error) {
	layout := "02/01/2006"
	date, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %v", err)
	}
	return date, nil
}
