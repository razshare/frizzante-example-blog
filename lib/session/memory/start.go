package memory

var States = map[string]*State{}

func Start(id string) *State {
	v, ok := States[id]
	if !ok {
		States[id] = &State{}
		return States[id]
	}
	return v
}
