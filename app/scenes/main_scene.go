package scenes

import "../../src"
import "../../src/attributes/body"
import "../../src/attributes/sprite"
import "../../src/attributes/collision"
import "../obj"

func MainScene() *src.Scene {
	mainScene := src.NewScene("main")

	player := obj.NewPlayer(body.NewBody(70, 100, 10, 10))
	mainScene.AddObj(player)

	wall1 := obj.NewWall(body.NewBody(10, 0, 10, 200))
	mainScene.AddObj(wall1)

	wall2 := obj.NewWall(body.NewBody(300, 0, 10, 200))
	mainScene.AddObj(wall2)

	wall3 := obj.NewWall(body.NewBody(150, 0, 10, 200))
	wall3.GetAttr("sprite").(*sprite.Sprite).SetColor(sprite.GetColor(0xff, 0, 0))
	wall3.GetAttr("collision").(*collision.Collision).SetLayer(2)
	mainScene.AddObj(wall3)

	return mainScene
}
