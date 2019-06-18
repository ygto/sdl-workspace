package src

type Director struct {
	scenes      map[string]*Scene
	activeScene string
}

func NewDirector() *Director {
	d := new(Director)
	d.scenes = make(map[string]*Scene)
	return d
}
func (d *Director) AddScene(s *Scene) {
	d.scenes[s.name] = s
}
func (d *Director) SetActiveScene(name string) {
	d.activeScene = name
}
func (d *Director) GetActiveScene() *Scene {
	return d.scenes[d.activeScene]
}
func (d *Director) GetScene(name string) *Scene {
	return d.scenes[name]
}
func (d *Director) Update() {
	updateScene(d.GetActiveScene())
}