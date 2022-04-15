package modes

import (
	"testing"
)

func Test_Mode_t_String(t *testing.T) {
	var (
		mode     Mode_t = 100
		expected        = "<Mode: Unknown>"
	)
	got := mode.String()

	if expected != got {
		t.Errorf("\nExpected: %s\nGot: %s\n", expected, got)
	}
}
