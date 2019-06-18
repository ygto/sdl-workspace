package sprite

import "github.com/veandco/go-sdl2/sdl"
import "../../../src"
import "../position"

type Sprite struct {
	active bool
	shape  *sdl.Rect
	size   struct {
		W int32
		H int32
	}
}

func NewSprite(w int32, h int32) *Sprite {
	attr := new(Sprite)
	attr.size.W = w
	attr.size.H = h
	return attr
}
func (s *Sprite) GetName() string {
	return "sprite"
}
func (attr *Sprite) Init(obj src.GameObjectInterface) {

	if obj.GetAttr("position") == nil {
		pos := position.NewPosition(0, 0)
		obj.AddAttr(pos)
	}

	surface := src.GetSurface()

	x, y := obj.GetAttr("position").(*position.Position).GetPosition()

	surface.FillRect(&sdl.Rect{int32(x), int32(y), attr.size.W, attr.size.H}, 0xffff0000)
}

func (attr *Sprite) Update(obj src.GameObjectInterface) {

	x, y := obj.GetAttr("position").(*position.Position).GetPosition()

	surface := src.GetSurface()
	surface.FillRect(&sdl.Rect{int32(x), int32(y), attr.size.W, attr.size.H}, 0xffff0000)
}
