package src

type Attribute interface {
	GetName() string
	Init(obj GameObjectInterface)
	Update(obj GameObjectInterface)
}
