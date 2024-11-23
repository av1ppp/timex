package timex_test

import (
	"encoding/json"
	"testing"

	"github.com/av1ppp/timex"
)

var (
	durationJSONValue  = timex.Hour + timex.Minute*2 + timex.Second*3 + timex.Millisecond*250
	durationJSONString = `"01:02:03.25"`
)

func TestDurationMarshalJSON(t *testing.T) {
	got, err := json.Marshal(durationJSONValue)
	if err != nil {
		t.Fatal("failed to marshal:", err)
	}

	if string(got) != durationJSONString {
		t.Fatalf("want: %q got: %q", durationJSONString, string(got))
	}
}

func TestDurationUnmarshalJSON(t *testing.T) {
	var d timex.Duration

	err := json.Unmarshal([]byte(durationJSONString), &d)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if d != durationJSONValue {
		t.Fatalf("want: %v got: %v", durationJSONValue, d)
	}
}
