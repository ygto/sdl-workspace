package src

type GameObjectInterface interface {
	AddAttr(attr Attribute)
	GetAttr(name string) Attribute
	GetAttrs() map[string]Attribute
	Update()
}
type GameObject struct {
	Attributes map[string]Attribute
}

func NewGameObject() *GameObject {
	obj := new(GameObject)
	obj.Attributes = make(map[string]Attribute)
	return obj
}

func (o *GameObject) AddAttr(attr Attribute) {
	o.Attributes[attr.GetName()] = attr
	attr.Init(o)
}
func (o *GameObject) GetAttrs() map[string]Attribute {
	return o.Attributes
}
func (o *GameObject) GetAttr(name string) Attribute {
	return o.Attributes[name]
}

func (g *GameObject) Update() {

}
func updateGameObject(obj GameObjectInterface) {
	for _, attr := range obj.GetAttrs() {
		attr.Update(obj)
	}
	obj.Update()
}
