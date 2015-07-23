// definice přikazu volaných z UI

package logic

type AppState uint

const (
	NONE AppState = iota
	ENTRANCE
	MENU_ONLINE
	MENU_OFFLINE
)

type Application struct {
	state AppState
}

func (a *Application) State() AppState {
	return a.state
}

// switch to state ABOUT if possible and call cb on success
// func (a *Application) About(cb func()) (ok bool) {
// 	switch a.state {
// 	case NONE, ENTRANCE:
// 		a.state = ABOUT
// 		ok = true
// 		println("state is now ABOUT")
// 		cb()
// 	}
// 	return
// }

// switch to state ENTRANCE if possible and call cb on success
func (a *Application) Entrance(cb func(interface{})) (ok bool) {
	switch a.state {
	case NONE:
		a.state = ENTRANCE
		ok = true
		println("state is now ENTRANCE")
		cb(nil)
	}
	return
}
