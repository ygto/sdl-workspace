package obj

import (
	"../../src"
	"../../src/attributes/body"
	"../../src/attributes/collision"
	"../../src/attributes/sprite"
)

type AI struct {
	activeColor uint32
	color       uint32
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

func NewAI(layer string, b *body.Body) *AI {
	obj := new(AI)
	obj.GameObject = src.NewGameObject()
	obj.accX = 0.05
	obj.color = sprite.GetColor(0xb7, 0x07, 0xaa)
	obj.activeColor = sprite.GetColor(0xf7, 0x47, 0xfa)

	obj.AddAttr(b)

	s := sprite.NewSprite(obj.color, b)
	obj.AddAttr(s)

	c := collision.NewCollision(obj, layer)
	obj.AddAttr(c)

	return obj
}

func (obj *AI) Update() {

	obj.Body().X += obj.accX
	if obj.Collision().TagCollision("wall") != nil {
		obj.accX *= -1
		obj.Body().RevertChanges()
	}
	if obj.Collision().TagCollision("player") != nil {
		obj.GetAttr("sprite").(*sprite.Sprite).SetColor(obj.activeColor)
	} else {
		obj.GetAttr("sprite").(*sprite.Sprite).SetColor(obj.color)
	}
}
