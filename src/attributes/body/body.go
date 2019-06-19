package body

import "../../../src"

type Body struct {
	src.Attribute
	X float32
	Y float32
	Z float32
	W float32
	H float32
}

func NewBody(x float32, y float32, z float32, w float32, h float32) *Body {
	attr := new(Body)
	attr.SetBody(x, y, z, w, h)
	return attr
}
func (attr *Body) GetName() string {
	return "body"
}

func (attr *Body) SetBody(x float32, y float32, z float32, w float32, h float32) {
	attr.X = x
	attr.Y = y
	attr.Z = z
	attr.W = w
	attr.H = h
}
func (attr *Body) Init(obj src.GameObjectInterface) {

}

func (attr *Body) Update(obj src.GameObjectInterface) {

}
