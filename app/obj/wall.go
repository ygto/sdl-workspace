package obj

import (
	"../../src"
	"../../src/attributes/body"
	"../../src/attributes/collision"
	"../../src/attributes/sprite"
)

type Wall struct {
	*src.GameObject
}

func NewWall(b *body.Body) *Wall {
	obj := new(Wall)
	obj.GameObject = src.NewGameObject()

	obj.AddAttr(b)

	s := sprite.NewSprite(sprite.GetColor(0xff, 0xaa, 0x01), b)
	obj.AddAttr(s)

	c := collision.NewCollision(b, "wall")
	obj.AddAttr(c)

	return obj
}
