package obj

import (
	"../../src"
	"../../src/attributes/body"
	"../../src/attributes/collision"
)

type HiddenWall struct {
	*src.GameObject
}

func NewHiddenWall(layer string, b *body.Body) *HiddenWall {
	obj := new(HiddenWall)
	obj.GameObject = src.NewGameObject()

	obj.AddAttr(b)

	c := collision.NewCollision(obj, layer)
	obj.AddAttr(c)

	return obj
}
