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
	color  uint32
}

func NewSprite(color uint32) *Sprite {
	attr := new(Sprite)
	attr.color = color
	return attr
}
func (s *Sprite) GetName() string {
	return "sprite"
}
func (s *Sprite) SetColor(c uint32) {
	s.color = c
}
func (attr *Sprite) Init(obj src.GameObjectInterface) {
	if obj.GetAttr("body") == nil {
		pos := body.NewBody(0, 0, 0, 0)
		obj.AddAttr(pos)
	}
}

func (attr *Sprite) Update(obj src.GameObjectInterface) {
	surface := src.GetSurface()
	b := obj.GetAttr("body").(*body.Body)
	surface.FillRect(&sdl.Rect{int32(b.X), int32(b.Y), int32(b.W), int32(b.H)}, attr.color)
}
