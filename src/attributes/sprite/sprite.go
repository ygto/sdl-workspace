package sprite

import (
	"encoding/binary"
	"github.com/veandco/go-sdl2/sdl"
)
import "../../../src"
import "../body"

const (
	RED   = 0xffff0000
	GREEN = 0xff00ff00
	BLUE  = 0xff0000ff
)

func GetColor(red byte, green byte, blue byte) uint32 {
	return binary.BigEndian.Uint32([]byte{0xff, red, green, blue})
}

type Sprite struct {
	src.Attribute
	active bool
	shape  *sdl.Rect
	body   *body.Body
	color  uint32
}

func NewSprite(color uint32, b *body.Body) *Sprite {
	attr := new(Sprite)
	attr.color = color
	attr.body = b
	return attr
}
func (s *Sprite) GetName() string {
	return "sprite"
}

func (s *Sprite) SetColor(c uint32) {
	s.color = c
}

func (attr *Sprite) Update(obj src.GameObjectInterface) {
	addDrawQueue(attr.body.Z,
		&sdl.Rect{int32(attr.body.X), int32(attr.body.Y), int32(attr.body.W), int32(attr.body.H)},
		attr.color,
	)
}
