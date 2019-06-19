package obj

import (
	"../../src"
	"../../src/attributes/body"
	"../../src/attributes/collision"
	"../../src/attributes/sprite"
)

type Player struct {
	*src.GameObject
	acc float32
}

func (p *Player) Collision() *collision.Collision {
	return p.GetAttr("collision").(*collision.Collision)
}
func (p *Player) Body() *body.Body {
	return p.GetAttr("body").(*body.Body)
}

func NewPlayer(b *body.Body) *Player {
	obj := new(Player)
	obj.GameObject = src.NewGameObject()
	obj.acc = 0.05

	obj.AddAttr(b)

	s := sprite.NewSprite(sprite.GREEN, b)
	obj.AddAttr(s)

	c := collision.NewCollision(b, "player")
	obj.AddAttr(c)

	return obj
}

func (obj *Player) Update() {
	if obj.Collision().TagCollision("wall") != nil {
		obj.acc *= -1
	}
	obj.Body().X += obj.acc
}
