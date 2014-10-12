package paypal

import (
	"errors"
	"time"
)

const (
	ISO8601Date     = "2006-01-02 MST"
	ISO8601Datetime = "2006-01-02 15:04:05 MST"
)

type Date struct {
	time.Time
}

type Datetime struct {
	time.Time
}

type DatetimeRFC3339 struct {
	time.Time
}

// String returns the formatted date
func (d Date) String() string {
	return d.Time.Format(ISO8601Date)
}

// MarshalJSON implements the json.Marshaler interface.
func (d Date) MarshalJSON() ([]byte, error) {
	if y := d.Year(); y < 0 || y >= 10000 {
		// ISO8601 is clear that years are 4 digits exactly.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(d.Format(`"` + ISO8601Date + `"`)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Date) UnmarshalJSON(data []byte) (err error) {
	d.Time, err = time.Parse(`"`+ISO8601Date+`"`, string(data))
	return
}

// String returns the formatted date
func (d Datetime) String() string {
	return d.Time.Format(ISO8601Datetime)
}

// MarshalJSON implements the json.Marshaler interface.
func (d Datetime) MarshalJSON() ([]byte, error) {
	if y := d.Year(); y < 0 || y >= 10000 {
		// ISO8601 is clear that years are 4 digits exactly.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(d.Format(`"` + ISO8601Datetime + `"`)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Datetime) UnmarshalJSON(data []byte) (err error) {
	d.Time, err = time.Parse(`"`+ISO8601Datetime+`"`, string(data))
	return
}

// String returns the formatted date
func (d DatetimeRFC3339) String() string {
	return d.Time.Format(time.RFC3339)
}

// MarshalJSON implements the json.Marshaler interface.
func (d DatetimeRFC3339) MarshalJSON() ([]byte, error) {
	if y := d.Year(); y < 0 || y >= 10000 {
		// ISO8601 is clear that years are 4 digits exactly.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(d.Format(`"` + time.RFC3339 + `"`)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *DatetimeRFC3339) UnmarshalJSON(data []byte) (err error) {
	d.Time, err = time.Parse(`"`+time.RFC3339+`"`, string(data))
	return
}
