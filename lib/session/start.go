package session

var Sessions = map[string]*Session{}

func Start(id string) *Session {
	v, ok := Sessions[id]
	if !ok {
		Sessions[id] = &Session{}
		return Sessions[id]
	}
	return v
}
