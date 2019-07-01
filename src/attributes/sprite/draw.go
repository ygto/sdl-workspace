package sprite

import (
	"../../../src"
	"github.com/veandco/go-sdl2/sdl"
	"sort"
)

type drawCollect []*draw

type draw struct {
	*sdl.Rect
	color uint32
	z     float32
}

func (d drawCollect) Len() int {
	return len(d)
}
func (d drawCollect) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d drawCollect) Less(i, j int) bool {
	return d[i].z < d[j].z
}

var draws = make(drawCollect, 0, 0)

func addDrawQueue(z float32, rect *sdl.Rect, color uint32) {
	d := draw{rect, color, z}
	draws = append(draws, &d)
}

func Draw() {
	sort.Sort(draws)
	surface := src.GetSurface()
	surface.FillRect(nil, 0)
	for _, d := range draws {
		surface.FillRect(d.Rect, d.color)
	}
	src.GetWindow().UpdateSurface()
	draws = make(drawCollect, 0, 0)
}
