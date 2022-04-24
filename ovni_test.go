// Copyright (c) 2022, Daniel M. Jaén
// All rights reserved.

package ovni

import(
	"testing"

	"github.com/dilmorja/ovni/modes"
)

func Test_New(t *testing.T) {
	expected := []interface{}{
		"<Mode: WASM>",
	}

	x := New(modes.WASM)

	got := []interface{}{
		x.Mode.String(),
	}

	for i, v := range got {
		if v != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], v)
		}
	}
}