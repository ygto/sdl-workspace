package sprite

import (
	"../../../src"
	"github.com/veandco/go-sdl2/sdl"
	"sort"
)

var rects = make(map[float32]map[*sdl.Rect]uint32)

func addDrawQueue(layer float32, rect *sdl.Rect, color uint32) {
	if _, ok := rects[layer]; !ok {
		rects[layer] = make(map[*sdl.Rect]uint32)
	}
	rects[layer][rect] = color
}
func Draw() {
	keys := []int{}
	for k := range rects {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	surface := src.GetSurface()
	surface.FillRect(nil, 0)
	for k := range keys {
		for rect, color := range rects[float32(k)] {
			surface.FillRect(rect, color)
		}
	}
	src.GetWindow().UpdateSurface()
	rects = make(map[float32]map[*sdl.Rect]uint32)
}
