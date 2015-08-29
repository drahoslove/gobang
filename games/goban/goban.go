// game := goban.newGoban(9)
// game.Place(x, y, color int)
// game.NewGroup(x,y int)
// group.HasLIberty()
// group.Remove()

package goban

import (
	"fmt"
)

/* consts */
type Color int

const (
	BLACK Color = 0
	WHITE Color = 1
)

const (
	OK = iota
	NOTOK
)

//////////////
/*  stone  */
type stone struct {
	x, y  int
	color Color
	goban Goban
}

func (s stone) hasLiberty() bool {
	if s.x > 0 && s.goban[s.x-1][s.y] == nil {
		return true
	}
	if s.y > 0 && s.goban[s.x][s.y-1] == nil {
		return true
	}
	if s.x < len(s.goban)-1 && s.goban[s.x+1][s.y] == nil {
		return true
	}
	if s.y < len(s.goban[s.x])-1 && s.goban[s.x][s.y+1] == nil {
		return true
	}
	return false
}

func (s *stone) swap() {
	switch s.color {
	case BLACK:
		s.color = WHITE
	case WHITE:
		s.color = BLACK
	}
}

//////////////
/*  Goban  */
type Goban [][]*stone

func newGoban(size int) Goban {
	g := make(Goban, size)
	for i := range g {
		g[i] = make([]*stone, size)
	}
	return g
}

func (g Goban) PlaceStone(x, y int, color Color) (ok int) {
	if g[x][y] != nil {
		return NOTOK
	}
	stone := stone{x: x, y: y, color: color, goban: g}

	g[x][y] = &stone

	return OK
}

func (g Goban) forEach(f func(*stone)) {
	for i := range g {
		for j := range g[i] {
			if g[i][j] != nil {
				f(g[i][j])
			}
		}
	}
}

func (g Goban) String() string {
	str := ""
	for i := range g {
		for j := range g[i] {
			switch {
			case g[i][j] == nil:
				str += "+"
			case g[i][j].color == BLACK:
				str += "☺"
			case g[i][j].color == WHITE:
				str += "☻"
			}
		}
		str += "\n"
	}
	return str
}

//////////////
/*  group  */
type group []*stone

func (gb Goban) NewGroup(x, y int) group {
	grp := make(group, 0)
	open := make(group, 0)
	open.push(gb[x][y])

	var point *stone
	var s stone

	for len(open) > 0 {
		point = open.pop()
		s = *point
		if !grp.contain(point) {
			if s.x > 0 && open.push(gb[s.x-1][s.y]) {
				// println("left")
			}
			if s.y > 0 && open.push(gb[s.x][s.y-1]) {
				// println("bot")
			}
			if s.x < len(gb)-1 && open.push(gb[s.x+1][s.y]) {
				// println("right")
			}
			if s.y < len(gb[s.x])-1 && open.push(gb[s.x][s.y+1]) {
				// println("top")
			}
		}
		grp.push(point)
	}
	return grp
}

func (g *group) push(point *stone) bool {
	if point != nil && !g.contain(point) && (len(*g) == 0 || point.color == g.top().color) {
		*g = append(*g, point)
		return true
	}
	return false
}

func (g *group) pop() *stone {
	s := (*g)[len(*g)-1]
	*g = (*g)[0 : len(*g)-1]
	return s
}

func (g *group) top() *stone {
	return (*g)[len(*g)-1]
}

func (g *group) contain(point *stone) bool {
	for _, p := range *g {
		if p == point {
			return true
		}
	}
	return false
}

func (g *group) HasLiberty() bool {
	for _, p := range *g {
		if p.hasLiberty() {
			return true
		}
	}
	return false
}

func (g *group) Remove() {
	for _, ps := range *g {
		ps.goban[ps.x][ps.y] = nil
	}
}

func (g group) String() string {
	str := ""
	for _, s := range g {
		str += fmt.Sprintf("[%v %v] ", s.x, s.y)
	}
	return str
}

///////////
/* Game */
type Game struct {
	color        Color // actual color
	goban        Goban
	sizeX, sizeY int
}

func New(size int) Game {
	var game Game
	game.sizeX = size
	game.sizeY = size
	game.color = BLACK
	game.goban = newGoban(size)
	return game
}

func (g *Game) Size() (int, int) {
	return g.sizeX, g.sizeY
}

func (g Game) Place(x, y int) error {
	g.goban.PlaceStone(x, y, g.color)
	return nil
}
