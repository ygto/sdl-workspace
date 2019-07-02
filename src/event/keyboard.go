package event

import (
	"github.com/veandco/go-sdl2/sdl"
)

const KEY_RIGHT = uint32(79)
const KEY_LEFT = uint32(80)
const KEY_DOWN = uint32(81)
const KEY_UP = uint32(82)
const KEY_ENTER = uint32(40)
const KEY_ESC = uint32(41)
const KEY_SPACE = uint32(44)

var keys = make(map[uint32]key)
var deleteCodes = make(map[uint32]bool)

type key struct {
	Repeat  bool
	Release bool
}

func All() map[uint32]key {
	return keys

}

func FlushKeyboard() {
	for k := range deleteCodes {
		delete(keys, k)
		delete(deleteCodes, k)
	}
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
		deleteCodes[code] = true

		keys[code] = key{
			Release: true,
		}

	}

	//fmt.Println("t", e.Type, "s", e.State, "r", e.Repeat, "m", e.Keysym.Mod, "sc", e.Keysym.Scancode, "sym", e.Keysym.Sym)
}

const NON_PRESSED = uint8(0)
const PRESSED = uint8(1)
const REPEATED = uint8(2)
const RELEASED = uint8(3)

func KeyDown(code uint32) uint8 {

	k, ok := keys[code]
	if !ok {
		return NON_PRESSED
	}
	if k.Repeat {
		return REPEATED
	}
	if k.Release {
		return RELEASED
	}
	return PRESSED
}
