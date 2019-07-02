package scenes

import "../../src"
import "../../src/event"
import "../../src/attributes/body"
import "../../src/attributes/sprite"
import "../obj"

func MainScene() *src.Scene {
	mainScene := src.NewScene("main")

	player := obj.NewPlayer("player", body.NewBody(70, 100, 1.1, 10, 10))
	mainScene.AddObj(player)

	ai := obj.NewAI("ai", body.NewBody(250, 100, 1, 10, 10))
	mainScene.AddObj(ai)

	wall1 := obj.NewWall("wall", body.NewBody(10, 0, 1, 10, 200))
	mainScene.AddObj(wall1)

	wall2 := obj.NewWall("wall", body.NewBody(300, 0, 1, 10, 200))
	mainScene.AddObj(wall2)

	wallTop := obj.NewHiddenWall("wall", body.NewBody(10, 0, 1, 300, 5))
	mainScene.AddObj(wallTop)

	bottomWall := obj.NewHiddenWall("wall", body.NewBody(10, 200, 1, 300, 5))
	mainScene.AddObj(bottomWall)

	wall3 := obj.NewWall("front", body.NewBody(150, 0, 2, 10, 200))
	wall3.GetAttr("sprite").(*sprite.Sprite).SetColor(sprite.GetColor(0xff, 0, 0))
	mainScene.AddObj(wall3)

	wall4 := obj.NewWall("back", body.NewBody(180, 0, 0, 10, 200))
	wall4.GetAttr("sprite").(*sprite.Sprite).SetColor(sprite.GetColor(0xff, 0, 0))
	mainScene.AddObj(wall4)

	controller := obj.NewController()
	controller.SetFunc(func(obj src.GameObjectInterface) {
		switch event.KeyDown(event.KEY_ESC) {
		case event.RELEASED:
			src.GetDirector().SetActiveScene("menu")
		}

	})
	mainScene.AddObj(controller)

	return mainScene
}
