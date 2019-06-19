package collision

import (
	"../../../src"
)
import "../body"

var items = make(map[byte]map[*Collision]bool)

type Collision struct {
	src.Attribute
	Layer      uint8
	Tag        string
	body       *body.Body
	collisions map[*Collision]string
}

func NewCollision(b *body.Body, layer uint8, tag string) *Collision {
	c := new(Collision)
	c.body = b
	c.Layer = layer
	c.Tag = tag
	c.collisions = make(map[*Collision]string)

	if _, ok := items[layer]; !ok {
		items[layer] = make(map[*Collision]bool)
	}

	return c
}

func (attr *Collision) GetName() string {
	return "collision"
}
func (attr *Collision) GetTag() string {
	return attr.Tag
}
func (attr *Collision) SetLayer(layer uint8) {
	delete(items[attr.Layer], attr)
	attr.Layer = layer
	if _, ok := items[layer]; !ok {
		items[layer] = make(map[*Collision]bool)
	}
	items[attr.Layer][attr] = true
}
func (attr *Collision) TagCollision(tag string) *Collision {

	for c, t := range attr.collisions {
		if tag == t {
			return c
		}
	}

	return nil
}
func (attr *Collision) Init(obj src.GameObjectInterface) {
	if obj.GetAttr("body") == nil {
		pos := body.NewBody(0, 0, 0, 0)
		obj.AddAttr(pos)
	}
	items[attr.Layer][attr] = true
}

func (attr *Collision) BeforeUpdate(obj src.GameObjectInterface) {

	attr.collisions = make(map[*Collision]string)

	for col, _ := range items[attr.Layer] {

		if attr.body == col.body {
			continue
		}
		if col == nil {
			continue
		}

		if collisionDetection(attr.body, col.body) {
			attr.collisions[col] = col.GetTag()
		}
	}

}
func (attr *Collision) Update(obj src.GameObjectInterface) {
}

func collisionDetection(b1 *body.Body, b2 *body.Body) bool {
	return b1.X < b2.X+b2.W &&
		b1.X+b1.W > b2.X &&
		b1.Y < b2.Y+b2.H &&
		b1.Y+b1.H > b2.Y
}
