package obj

import (
	"../../src"
)

type Controller struct {
	*src.GameObject
	updateFunc func(obj src.GameObjectInterface)
}

func NewController() *Controller {
	obj := new(Controller)
	obj.GameObject = src.NewGameObject()

	return obj
}
func (obj *Controller) SetFunc(f func(obj src.GameObjectInterface)) {
	obj.updateFunc = f
}

func (obj *Controller) Update() {
	obj.updateFunc(obj)
}
