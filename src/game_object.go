package src

type GameObjectInterface interface {
	AddAttr(attr AttributeInterface)
	GetAttr(name string) AttributeInterface
	GetAttrs() map[string]AttributeInterface
	Update()
	AfterUpdate()
}
type GameObject struct {
	Attributes map[string]AttributeInterface
}

func NewGameObject() *GameObject {
	obj := new(GameObject)
	obj.Attributes = make(map[string]AttributeInterface)
	return obj
}

func (o *GameObject) AddAttr(attr AttributeInterface) {
	o.Attributes[attr.GetName()] = attr
	attr.Init(o)
}
func (o *GameObject) GetAttrs() map[string]AttributeInterface {
	return o.Attributes
}
func (o *GameObject) GetAttr(name string) AttributeInterface {
	return o.Attributes[name]
}

func (g *GameObject) Update() {

}
func (g *GameObject) AfterUpdate() {

}
func updateGameObject(obj GameObjectInterface) {
	for _, attr := range obj.GetAttrs() {
		attr.BeforeUpdate(obj)
	}

	for _, attr := range obj.GetAttrs() {
		attr.Update(obj)
	}
	obj.Update()

	for _, attr := range obj.GetAttrs() {
		attr.AfterUpdate(obj)
	}
}
