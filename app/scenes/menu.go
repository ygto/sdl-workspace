package scenes

import "../../src"
import "../../src/event"
import "../../src/attributes/body"
import "../../src/attributes/sprite"
import "../obj"

func MenuScene() *src.Scene {
	menu := src.NewScene("menu")

	b1 := obj.NewWall("wall", body.NewBody(100, 100, 1, 50, 50))
	b1.GetAttr("sprite").(*sprite.Sprite).SetColor(sprite.GetColor(0xff, 0, 0))
	menu.AddObj(b1)

	b2 := obj.NewWall("wall", body.NewBody(100, 200, 1, 50, 50))
	b2.GetAttr("sprite").(*sprite.Sprite).SetColor(sprite.GetColor(0xff, 0, 0))
	menu.AddObj(b2)

	b3 := obj.NewWall("wall", body.NewBody(100, 300, 1, 50, 50))
	b3.GetAttr("sprite").(*sprite.Sprite).SetColor(sprite.GetColor(0xff, 0, 0))
	menu.AddObj(b3)

	controller := obj.NewController()
	controller.SetFunc(func(obj src.GameObjectInterface) {
		switch event.KeyDown(event.KEY_ESC) {
		case event.RELEASED:
			src.GetDirector().SetActiveScene("main")
		}

	})

	menu.AddObj(controller)

	return menu
}
