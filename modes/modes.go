// Copyright (c) 2022, Daniel M. Jaén
// All rights reserved.

package modes

type Mode_t int

const(
	Vanilla Mode_t = iota
	// TODO: Add support for WebAssembly
	// WASM
)

func (m Mode_t) String() string {
	switch m {
	case 0:
		return "<Mode: Vanilla>"
	default:
		return "<Mode: Unknown>"
	}
}