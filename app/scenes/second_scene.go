package scenes

import "../../src"
import "../obj"
import "../../src/attributes/position"

func SecondScene() *src.Scene {
	s := src.NewScene("second")

	player := obj.NewPlayer()
	s.AddObj(player)

	pos:=player.GetAttr("position").(*position.Position)

	pos.Y=300

	player2 := obj.NewPlayer()
	s.AddObj(player2)

	return s
}
