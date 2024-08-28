package models

import (
	"time"
)

// структура для работы со временем. перевод строки в rfc строку, или текущего времени в rfc.
type TimeWorker struct {
}

func (instance *TimeWorker) GetNowRfc() string {
	thisTime := time.Now()
	temp := thisTime.Format(time.RFC3339)
	return temp
}

func (instance *TimeWorker) ConvertStringToRfc(oldTime string) string {
	dateValue, err := time.Parse("2006.01.02", oldTime)
	if err == nil {
		temp := dateValue.Format(time.RFC3339)
		return temp
	}

	dateValue, err = time.Parse("2006-01-02", oldTime)
	if err == nil {
		temp := dateValue.Format(time.RFC3339)
		return temp
	}

	return oldTime
}
