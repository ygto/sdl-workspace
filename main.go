package main

import (
	"./app/scenes"
	"./src"
	"./src/attributes/sprite"
	"./src/event"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	src.SetWindow(window)

	surface, err := window.GetSurface()

	surface.FillRect(nil, 0)

	/*rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)*/
	window.UpdateSurface()

	director := src.NewDirector()
	src.SetDirector(director)

	director.AddScene(scenes.MainScene())
	director.AddScene(scenes.MenuScene())
	director.SetActiveScene("main")

	running := true
	for running {
		event.FlushKeyboard()
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				event.FillKeyboard(e.(*sdl.KeyboardEvent))
				break
			default:
			}
		}

		director.Update()
		sprite.Draw()
	}
}
