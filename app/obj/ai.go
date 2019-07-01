package obj

import (
	"../../src"
	"../../src/attributes/body"
	"../../src/attributes/collision"
	"../../src/attributes/sprite"
)

type AI struct {
	*src.GameObject
	accX float32
	accY float32
}

func (p *AI) Collision() *collision.Collision {
	return p.GetAttr("collision").(*collision.Collision)
}
func (p *AI) Body() *body.Body {
	return p.GetAttr("body").(*body.Body)
}

func NewAI(b *body.Body) *AI {
	obj := new(AI)
	obj.GameObject = src.NewGameObject()
	obj.accX = 0.05

	obj.AddAttr(b)

	s := sprite.NewSprite(sprite.GetColor(0xb7, 0x07, 0xaa), b)
	obj.AddAttr(s)

	c := collision.NewCollision(b, "ai")
	obj.AddAttr(c)

	return obj
}

func (obj *AI) Update() {

	obj.Body().X += obj.accX
	if obj.Collision().TagCollision("wall") != nil {
		obj.accX *= -1
		obj.Body().RevertChanges()
	}
}
