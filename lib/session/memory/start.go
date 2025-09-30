package memory

import "time"

var States = map[string]*State{}

func init() {
	go func() {
		for {
			for id, state := range States {
				if !state.LoggedIn {
					delete(States, id)
					continue
				}

				if time.Since(state.LastActivity) > 30*time.Minute {
					delete(States, id)
					continue
				}
			}
			time.Sleep(10 * time.Minute)
		}
	}()
}

func Start(id string) *State {
	state, ok := States[id]
	if !ok {
		States[id] = &State{}
		return States[id]
	}
	return state
}
