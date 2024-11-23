package timex

import "strings"

// MarshalJSON returns the JSON encoding of the Duration.
// The returned byte slice represents the duration as a string in the format "15:04:05.999999999".
func (self Duration) MarshalJSON() ([]byte, error) {
	return []byte("\"" + self.string() + "\""), nil
}

// UnmarshalJSON parses a JSON-encoded string to a Duration object.
// The input byte slice should represent a string in the format "15:04:05.999999999".
// It returns an error if the input cannot be parsed into a valid Duration.
func (self *Duration) UnmarshalJSON(b []byte) error {
	d, err := ParseDuration(strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}
	*self = d
	return nil
}
