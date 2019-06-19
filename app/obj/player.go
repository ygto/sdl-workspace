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

func NewPlayer(b *body.Body) *Player {
	p := new(Player)
	p.GameObject = src.NewGameObject()
	p.acc = 0.05

	p.AddAttr(b)

	s := sprite.NewSprite(sprite.GREEN)
	p.AddAttr(s)

	c := collision.NewCollision(b, 1, "player")
	p.AddAttr(c)

	return p
}

func (player *Player) Update() {

	b := player.GetAttr("body").(*body.Body)

	col := player.GetAttr("collision").(*collision.Collision)
	if col.TagCollision("wall") != nil {
		player.acc *= -1
	}
	if b.X < -1 {
		player.acc *= -1
	}
	b.X += player.acc

	/*if b.X > 199 && src.GetDirector().GetActiveScene().GetName() != "second" {
		src.GetDirector().SetActiveScene("second")
	}*/
}
