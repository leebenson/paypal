package paypal

import (
	"errors"
	"fmt"
	"time"
)

const (
	ISO8601 = "2006-01-02 MST"
)

type Date struct {
	time.Time
}

// String returns the formatted date
func (d Date) String() string {
	return d.Time.Format(ISO8601)
}

// MarshalJSON implements the json.Marshaler interface.
func (d Date) MarshalJSON() ([]byte, error) {
	if y := d.Year(); y < 0 || y >= 10000 {
		// ISO8601 is clear that years are 4 digits exactly.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(d.Format(`"` + ISO8601 + `"`)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Date) UnmarshalJSON(data []byte) (err error) {
	fmt.Println(string(data))
	d.Time, err = time.Parse(`"`+ISO8601+`"`, string(data))
	return
}
