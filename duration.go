package timex

import (
	"time"
)

var (
	durationLayout = "15:04:05.999999999"
	durationOrigin = time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
)

type Duration time.Duration

const (
	Nanosecond  = Duration(time.Nanosecond)
	Microsecond = Duration(time.Microsecond)
	Millisecond = Duration(time.Millisecond)
	Second      = Duration(time.Second)
	Minute      = Duration(time.Minute)
	Hour        = Duration(time.Hour)
	Day         = Hour * 24
	Week        = Day * 7
	Fortnight   = Week * 2
	Month       = Day * 30   // Approximation
	Year        = Day * 365  // Approximation
	Decade      = Year * 10  // Approximation
	Century     = Year * 100 // Approximation
	// Millennium  = Year * 1000 // Approximation. WARNING: overflows int64
)

// ParseDuration parses a string representation of a duration and returns
// the corresponding Duration.
//
// The string must be in the format "15:04:05.999999999".
func ParseDuration(s string) (Duration, error) {
	t, err := time.Parse(durationLayout, s)
	if err != nil {
		return -1, err
	}
	return Duration(t.Sub(durationOrigin)), nil
}

// String returns the string representation of the Duration.
func (self Duration) String() string {
	return self.string()
}

// Days returns the duration as a floating point number of days.
func (d Duration) Days() float64 {
	hour := d / Hour
	nsec := d % Hour
	return float64(hour) + float64(nsec)*(1e-9/60/60/24)
}

// Weeks returns the duration as a floating point number of days.
func (d Duration) Weeks() float64 {
	hour := d / Hour
	nsec := d % Hour
	return float64(hour) + float64(nsec)*(1e-9/60/60/24/7)
}

func (self Duration) string() string {
	t := durationOrigin.Add(time.Duration(self))
	return t.Format(durationLayout)
}
