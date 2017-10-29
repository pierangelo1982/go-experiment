package utils

import (
	"fmt"
	"time"
)

func ParseData(date string) time.Time {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		x := "1900-01-01"
		tm, err := time.Parse("2006-01-02", x)
		if err != nil {
			fmt.Println(err)
		}
		return tm
	}
	return t
}
