package sessions

var Sessions = map[string]*Session{}

//func init() {
//	go func() {
//		for {
//			for id, state := range Sessions {
//				if !state.LoggedIn {
//					delete(Sessions, id)
//					continue
//				}
//
//				if time.Since(state.LastActivity) > 30*time.Minute {
//					delete(Sessions, id)
//					continue
//				}
//			}
//			time.Sleep(10 * time.Minute)
//		}
//	}()
//}

func Start(id string) *Session {
	session, ok := Sessions[id]
	if !ok {
		Sessions[id] = &Session{}
		return Sessions[id]
	}
	return session
}
