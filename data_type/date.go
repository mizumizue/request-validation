package data_type

import "time"

type Date struct {
	time.Time
}

func (t *Date) UnmarshalParam(src string) error {
	location, _ := time.LoadLocation("Asia/Tokyo")
	ts, err := time.ParseInLocation("2006-01-02", src, location)
	*t = Date{ts}
	return err
}
