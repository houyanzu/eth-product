package mytime

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type DateTime time.Time

func NewFromTime(t time.Time) DateTime {
	return DateTime(t)
}

func NewFromNow() DateTime {
	return DateTime(time.Now())
}

func ParseInLocation(t string) (DateTime, error) {
	ext, err := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	if err != nil {
		return DateTime{}, err
	}
	return DateTime(ext), nil
}

func (date DateTime) MarshalJSON() ([]byte, error) {
	date = date.Add(-8 * time.Hour)
	var stamp = fmt.Sprintf("\"%s\"", time.Time(date).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (date *DateTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	b = bytes.Trim(b, "\"")
	ext, err := time.ParseInLocation("2006-01-02 15:04:05", string(b), time.Local)
	if err != nil {
		return err
	}
	*date = DateTime(ext)
	return nil
}

func (date *DateTime) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = DateTime(nullTime.Time)
	return
}

func (date DateTime) Value() (driver.Value, error) {
	y, m, d := time.Time(date).Date()
	h, i, s := time.Time(date).Clock()
	return time.Date(y, m, d, h, i, s, 0, time.Time(date).Location()), nil
}

func (date DateTime) Time() time.Time {
	return time.Time(date)
}

func (date DateTime) Format() string {
	return time.Time(date).Format("2006-01-02 15:04:05")
}

func (date DateTime) Before(date2 DateTime) bool {
	return time.Time(date).Before(time.Time(date2))
}

func (date DateTime) Add(duration time.Duration) DateTime {
	return DateTime(time.Time(date).Add(duration))
}
