package timex

import (
	"strings"

	"gopkg.in/yaml.v3"
)

// MarshalYAML returns the YAML encoding of the Duration.
// The returned byte slice represents the duration as a string in the format "15:04:05.999999999".
func (self Duration) MarshalYAML() (any, error) {
	return self.string(), nil
}

// UnmarshalYAML parses a YAML-encoded string to a Duration object.
// The input byte slice should represent a string in the format "15:04:05.999999999".
// It returns an error if the input cannot be parsed into a valid Duration.
func (self *Duration) UnmarshalYAML(value *yaml.Node) error {
	d, err := ParseDuration(strings.Trim(value.Value, "\""))
	if err != nil {
		return err
	}
	*self = d
	return nil
}
