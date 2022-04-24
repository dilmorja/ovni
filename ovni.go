// Copyright (c) 2022, Daniel M. Jaén
// All rights reserved.
package ovni

import(
	"github.com/dilmorja/ovni/modes"
)

type Alien struct {
	Mode modes.Mode_t
}

func New(mode modes.Mode_t) *Alien {
	var this = new(Alien)
	if mode.String() == "<Mode: Unknown>" {
		this.Mode = modes.WASM
		return this
	}
	this.Mode = mode
	return this
}