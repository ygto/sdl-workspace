package collision

import (
	"../../../src"
)
import "../body"

var items = make(map[*Collision]bool)

func addItem(c *Collision) {
	items[c] = true
}

type Collision struct {
	src.Attribute
	Tag        string
	obj        src.GameObjectInterface
	collisions map[*Collision]string
}

func NewCollision(obj src.GameObjectInterface, tag string) *Collision {
	c := new(Collision)
	c.obj = obj
	c.Tag = tag
	c.collisions = make(map[*Collision]string)
	addItem(c)

	return c
}

func (attr *Collision) GetName() string {
	return "collision"
}
func (attr *Collision) GetTag() string {
	return attr.Tag
}

func (attr *Collision) TagCollision(tag string) *Collision {

	s := src.GetDirector().GetActiveScene()
	for col, _ := range items {

		if col.Tag != tag || col.obj.GetScene() != s {
			continue
		}

		attrBody := attr.obj.GetAttr("body").(*body.Body)
		colBody := col.obj.GetAttr("body").(*body.Body)
		if collisionDetection(attrBody, colBody) {
			attr.collisions[col] = col.GetTag()
			return col
		}
	}

	return nil
}
func (attr *Collision) Init(obj src.GameObjectInterface) {
	if obj.GetAttr("body") == nil {
		pos := body.NewBody(0, 0, 0, 0, 0)
		obj.AddAttr(pos)
	}
	addItem(attr)
}

func (attr *Collision) BeforeUpdate(obj src.GameObjectInterface) {
	attr.collisions = make(map[*Collision]string)
}
func collisionDetection(b1 *body.Body, b2 *body.Body) bool {
	return b1.X < b2.X+b2.W &&
		b1.X+b1.W > b2.X &&
		b1.Y < b2.Y+b2.H &&
		b1.Y+b1.H > b2.Y
}
