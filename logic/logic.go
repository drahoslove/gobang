// definice přikazu volaných z UI

package logic

import (
	"github.com/Drahoslav7/gobang/games/goban"
	"log"
)

type AppState uint

func (state AppState) String() string {
	switch state {
	case NONE:
		return "NONE"
	case ENTRANCE:
		return "ENTRANCE"
	case MENU_ONLINE:
		return "MENU_ONLINE"
	case MENU_OFFLINE:
		return "MENU_OFFLINE"
	case PRACTICE_SETTINGS:
		return "PRACTICE_SETTINGS"
	case GAME_SINGLE:
		return "GAME_SINGLE"
	case DUEL_SETTINGS:
		return "DUEL_SETTINGS"
	case LAST_STATE:
		return "LAST_STATE"
	}
	return "UNKNOWN"
}

const (
	NONE AppState = iota
	ENTRANCE
	MENU_ONLINE
	MENU_OFFLINE
	PRACTICE_SETTINGS
	GAME_SINGLE
	DUEL_SETTINGS
	LAST_STATE
)

type CallBack func()

type Player struct {
	name   string
	secret string
	server string
}

type Game interface {
	// Size() (int, int)
	// SetSize(int, int)
	Place(int, int) error
	// Pass()
}

type Application struct {
	state      AppState
	player     Player
	game       Game
	onExitCbs  []CallBack
	onStateCbs map[AppState][]CallBack
	exiting    bool
}

func NewApp() (app Application) {
	app.onStateCbs = make(map[AppState][]CallBack)
	return
}

// support function it calls f if it is not null
func call(f CallBack) {
	if f != nil {
		f()
	}
}

// proto:
// func (a *Application) Name(cb CallBack) (ok bool) {
// 	call(cb)
// 	return
// }

func (a *Application) gotoState(state AppState) {
	a.state = state
	for _, cb := range a.onStateCbs[state] {
		call(cb)
	}
	log.Println("state is now", state)
	return
}

func (a *Application) State() AppState {
	return a.state
}

// switch to state ENTRANCE if possible and call cb on success
func (a *Application) Entrance(cb CallBack) (ok bool) {
	switch a.state {
	case NONE:
		a.gotoState(ENTRANCE)
		ok = true
		call(cb)
	}
	return
}

func (a *Application) SetPlayer(name string, cb CallBack) {
	a.player.name = name
	call(cb)
	return
}

func (a *Application) LoadSettings(cb CallBack) (ok bool) {
	// a.player = ??
	call(cb)
	return
}

func (a *Application) NewGame(size int, cb CallBack) (ok bool) {

	call(cb)
	return
}

func (a *Application) Practice(cb CallBack) (ok bool) {
	switch a.state {
	case MENU_OFFLINE, ENTRANCE:
		a.gotoState(PRACTICE_SETTINGS)
		ok = true
		call(cb)
	}
	return
}

func (a *Application) Play(size int, cb CallBack) (ok bool) {
	switch a.state {
	case PRACTICE_SETTINGS:
		a.gotoState(GAME_SINGLE)
		a.game = goban.New(size)
		ok = true
		call(cb)

	case DUEL_SETTINGS:
		// TODO: implement
	}
	return
}

func (a *Application) Exit() {
	if a.exiting {
		return
	}
	a.exiting = true
	for _, cb := range a.onExitCbs {
		call(cb)
	}
}

// event listenners registrations:

func (a *Application) OnExit(cb CallBack) {
	a.onExitCbs = append(a.onExitCbs, cb)
}

func (a *Application) OnState(state AppState, cb CallBack) {
	a.onStateCbs[state] = append(a.onStateCbs[state], cb)
	log.Println(a.onStateCbs)
}
