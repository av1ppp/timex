package timex_test

import (
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/av1ppp/timex"
)

var (
	durationYAMLValue  = timex.Hour + timex.Minute*2 + timex.Second*3 + timex.Millisecond*250
	durationYAMLString = "\"01:02:03.25\"\n"
)

func TestDurationMarshalYAML(t *testing.T) {
	got, err := yaml.Marshal(durationYAMLValue)
	if err != nil {
		t.Fatal("failed to marshal:", err)
	}

	if string(got) != durationYAMLString {
		t.Fatalf("want: %q got: %q", durationYAMLString, string(got))
	}
}

func TestDurationUnmarshalYAML(t *testing.T) {
	var d timex.Duration

	err := yaml.Unmarshal([]byte(durationYAMLString), &d)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if d != durationYAMLValue {
		t.Fatalf("want: %v got: %v", durationYAMLValue, d)
	}
}
