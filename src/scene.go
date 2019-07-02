package src

type Scene struct {
	name    string
	objects []GameObjectInterface
}

func NewScene(name string) *Scene {
	scene := new(Scene)
	scene.name = name
	scene.objects = make([]GameObjectInterface, 0, 0)
	return scene
}

func (s *Scene) GetName() string {
	return s.name
}
func (s *Scene) AddObj(o GameObjectInterface) {
	o.SetScene(s)
	s.objects = append(s.objects, o)
}

func updateScene(s *Scene) {
	for _, o := range s.objects {
		updateGameObject(o)
	}
}
