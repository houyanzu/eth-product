package mytime

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type DateTime struct {
	data   time.Time
	isDate bool
}

func NewFromTime(t time.Time) DateTime {
	return DateTime{t, false}
}

func NewFromUnix(timestamp int64) DateTime {
	return DateTime{time.Unix(timestamp, 0), false}
}

func NewFromNow() DateTime {
	return DateTime{time.Now(), false}
}

func ParseInLocation(t string) (DateTime, error) {
	if len(t) == 19 {
		ext, err := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
		if err != nil {
			return DateTime{}, err
		}
		return DateTime{ext, false}, nil
	} else {
		ext, err := time.ParseInLocation("2006-01-02", t, time.Local)
		if err != nil {
			return DateTime{}, err
		}
		return DateTime{ext, true}, nil
	}
}

func (date DateTime) MarshalJSON() ([]byte, error) {
	//date = date.Add(-8 * time.Hour)
	if date.isDate {
		var stamp = fmt.Sprintf("\"%s\"", date.data.Format("2006-01-02"))
		return []byte(stamp), nil
	} else {
		var stamp = fmt.Sprintf("\"%s\"", date.data.Format("2006-01-02 15:04:05"))
		return []byte(stamp), nil
	}
}

func (date *DateTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	b = bytes.Trim(b, "\"")
	if len(b) == 19 {
		ext, err := time.ParseInLocation("2006-01-02 15:04:05", string(b), time.Local)
		if err != nil {
			return err
		}
		*date = DateTime{ext, false}
		return nil
	} else {
		ext, err := time.ParseInLocation("2006-01-02", string(b), time.Local)
		if err != nil {
			return err
		}
		*date = DateTime{ext, true}
		return nil
	}

}

func (date *DateTime) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = DateTime{nullTime.Time, false}
	return
}

func (date DateTime) Value() (driver.Value, error) {
	y, m, d := date.data.Date()
	h, i, s := date.data.Clock()
	return time.Date(y, m, d, h, i, s, 0, date.data.Location()), nil
}

func (date DateTime) Time() time.Time {
	return date.data
}

func (date DateTime) Format() string {
	layout := "2006-01-02 15:04:05"
	if date.isDate {
		layout = "2006-01-02"
	}
	return date.data.Format(layout)
}

func (date DateTime) Before(date2 DateTime) bool {
	return date.data.Before(date2.data)
}

func (date DateTime) Add(duration time.Duration) DateTime {
	date.data = date.data.Add(duration)
	return date
}

func (date DateTime) Unix() int64 {
	return date.data.Unix()
}

func (date DateTime) UnixNano() int64 {
	return date.data.UnixNano()
}

func (date DateTime) AddDate(years, months, days int) DateTime {
	return NewFromTime(date.data.AddDate(years, months, days))
}

func (date DateTime) SetIsDate(isDate bool) DateTime {
	date.isDate = isDate
	return date
}

func (date DateTime) IsDate() bool {
	return date.isDate
}

func (date DateTime) DiffDays(date2 DateTime) int {
	d := date.data.Sub(date2.data)
	return int(d.Hours() / 24)
}
