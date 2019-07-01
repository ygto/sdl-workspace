package body

import (
	"../../../src"
)

type Body struct {
	Layer string

	src.Attribute
	X float32
	Y float32
	Z float32
	W float32
	H float32

	//previous body
	P_X float32
	P_Y float32
	P_Z float32
	P_W float32
	P_H float32
}

func NewBody(layer string, x float32, y float32, z float32, w float32, h float32) *Body {
	attr := new(Body)
	attr.SetBody(layer, x, y, z, w, h)
	return attr
}
func (attr *Body) GetName() string {
	return "body"
}

func (attr *Body) SetBody(layer string, x float32, y float32, z float32, w float32, h float32) {
	attr.Layer = layer
	attr.X = x
	attr.Y = y
	attr.Z = z
	attr.W = w
	attr.H = h
}

func (attr *Body) SetPBody() {
	attr.P_X = attr.X
	attr.P_Y = attr.Y
	attr.P_Z = attr.Z
	attr.P_W = attr.W
	attr.P_H = attr.H
}

func (attr *Body) RevertChanges() {
	attr.X = attr.P_X
	attr.Y = attr.P_Y
	attr.Z = attr.P_Z
	attr.W = attr.P_W
	attr.H = attr.P_H
}

func (attr *Body) Init(obj src.GameObjectInterface) {

}

func (attr *Body) Update(obj src.GameObjectInterface) {
}
func (attr *Body) BeforeUpdate(obj src.GameObjectInterface) {
}

func (attr *Body) AfterUpdate(obj src.GameObjectInterface) {

	attr.SetPBody()
}
