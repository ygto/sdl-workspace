package src

type AttributeInterface interface {
	GetName() string
	Init(obj GameObjectInterface)
	BeforeUpdate(obj GameObjectInterface)
	Update(obj GameObjectInterface)
}
type Attribute struct {
}

func (attr *Attribute) Init(obj GameObjectInterface) {

}
func (attr *Attribute) BeforeUpdate(obj GameObjectInterface) {

}

func (attr *Attribute) Update(obj GameObjectInterface) {

}
