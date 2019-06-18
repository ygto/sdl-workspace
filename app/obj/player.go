package obj

import (
	"../../src"
	"../../src/attributes/position"
	"../../src/attributes/sprite"
)

type Player struct {
	*src.GameObject
	acc float32
}

func NewPlayer() *Player {
	p := new(Player)
	p.GameObject = src.NewGameObject()
	p.acc = 0.05

	pos := position.NewPosition(0, 100)
	p.AddAttr(pos)

	s := sprite.NewSprite(10, 10)
	p.AddAttr(s)

	return p
}

func (player *Player) Update() {

	pos := player.GetAttr("position").(*position.Position)
	pos.X += player.acc

	if pos.X > 200 {
		player.acc *= -1
	}
	if pos.X < -1 {
		player.acc *= -1
	}

	if pos.X > 199 && src.GetDirector().GetActiveScene().GetName() != "second" {
		src.GetDirector().SetActiveScene("second")
	}
}
