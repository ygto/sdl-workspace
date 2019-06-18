package src

import "github.com/veandco/go-sdl2/sdl"

var window *sdl.Window
var surface *sdl.Surface
var director *Director

func SetDirector(d *Director) {
	director = d
}
func SetWindow(w *sdl.Window) {
	window = w
	s, err := w.GetSurface()
	if err != nil {
		panic(err)
	}
	surface = s

}
func GetWindow() *sdl.Window {
	return window
}

func GetSurface() *sdl.Surface {
	return surface
}

func GetDirector() *Director {
	return director
}
