package event

import (
	"github.com/veandco/go-sdl2/sdl"
)

const KEY_RIGHT = uint32(79)
const KEY_LEFT = uint32(80)
const KEY_DOWN = uint32(81)
const KEY_UP = uint32(82)

var keys = make(map[uint32]key)

type key struct {
	Repeat bool
}

func All() map[uint32]key {
	return keys

}

func FlushKeyboard() {
	keys = make(map[uint32]key)
}

func FillKeyboard(e *sdl.KeyboardEvent) {
	code := uint32(e.Keysym.Scancode)
	state := e.State
	repeat := e.Repeat == 1
	if state == 1 {
		keys[code] = key{
			Repeat: repeat,
		}
	} else {
		delete(keys, code)
	}

	//fmt.Println("t", e.Type, "s", e.State, "r", e.Repeat, "m", e.Keysym.Mod, "sc", e.Keysym.Scancode, "sym", e.Keysym.Sym)
}

const NON_PRESSED = uint8(0)
const PRESSED = uint8(1)
const REPEATED = uint8(2)

func KeyDown(code uint32) uint8 {
	k, ok := keys[code]
	if !ok {
		return 0
	}
	if k.Repeat {
		return 2
	}
	return 1
}
