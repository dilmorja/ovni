// Copyright (c) 2022, Daniel M. Jaén
// All rights reserved.

package modes

type Mode_t int

const(
	// WebAssembly
	WASM Mode_t = iota
	// Vanilla
)

func (m Mode_t) String() string {
	switch m {
	case 0:
		return "<Mode: WASM>"
	default:
		return "<Mode: Unknown>"
	}
}