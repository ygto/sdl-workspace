package position

import "../../../src"

type Position struct {
	X float32
	Y float32
}

func NewPosition(x float32, y float32) *Position {
	attr := new(Position)
	attr.SetPosition(x, y)
	return attr
}
func (attr *Position) GetName() string {
	return "position"
}

func (attr *Position) SetPosition(x float32, y float32) {
	attr.X = x
	attr.Y = y
}
func (attr *Position) GetPosition() (float32, float32) {
	return attr.X, attr.Y
}
func (attr *Position) Init(obj src.GameObjectInterface) {

}

func (attr *Position) Update(obj src.GameObjectInterface) {

}
