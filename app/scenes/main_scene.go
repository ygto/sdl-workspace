package scenes

import "../../src"
import "../obj"

func MainScene() *src.Scene {
	mainScene := src.NewScene("main")

	player := obj.NewPlayer()
	mainScene.AddObj(player)

	return mainScene
}
