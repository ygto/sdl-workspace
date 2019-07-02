package obj

import (
	"../../src"
	"../../src/attributes/body"
	"../../src/attributes/collision"
	"../../src/attributes/sprite"
	"../../src/event"
)

type Player struct {
	*src.GameObject
	accX      float32
	accY      float32
	speedLow  float32
	speedFast float32
	rollback  bool
}

func (p *Player) Collision() *collision.Collision {
	return p.GetAttr("collision").(*collision.Collision)
}
func (p *Player) Body() *body.Body {
	return p.GetAttr("body").(*body.Body)
}

func NewPlayer(layer string, b *body.Body) *Player {
	obj := new(Player)

	obj.speedLow = 0.05
	obj.speedFast = 0.1

	obj.GameObject = src.NewGameObject()

	obj.AddAttr(b)

	s := sprite.NewSprite(sprite.GREEN, b)
	obj.AddAttr(s)

	c := collision.NewCollision(obj, layer)
	obj.AddAttr(c)

	return obj
}

func (obj *Player) Update() {

	obj.accX = 0
	obj.accY = 0

	switch event.KeyDown(event.KEY_DOWN) {
	case event.PRESSED:
		obj.accY = obj.speedLow
	case event.REPEATED:
		obj.accY = obj.speedFast
	}
	switch event.KeyDown(event.KEY_UP) {
	case event.PRESSED:
		obj.accY = -obj.speedLow
	case event.REPEATED:
		obj.accY = -obj.speedFast
	}

	switch event.KeyDown(event.KEY_LEFT) {
	case event.PRESSED:
		obj.accX = -obj.speedLow
	case event.REPEATED:
		obj.accX = -obj.speedFast
	}

	switch event.KeyDown(event.KEY_RIGHT) {
	case event.PRESSED:
		obj.accX = obj.speedLow
	case event.REPEATED:
		obj.accX = obj.speedFast
	}

	obj.Body().X += obj.accX
	obj.Body().Y += obj.accY

	if c := obj.Collision().TagCollision("wall"); c != nil {
		obj.Body().RevertChanges()
	}
}
