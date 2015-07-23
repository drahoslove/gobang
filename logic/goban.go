// goban := newGoban(9)
// goban.Place(x, y, color int)
// goban.NewGroup(x,y int)
// group.HasLIberty()
// group.Remove()

package logic

import (
	"fmt"
	. "gobang/consts"
)

/*  stone  */
type stone struct {
	x, y  int
	color int
	goban goban
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

/*  goban  */
type goban [][]*stone

func newGoban(size int) goban {
	g := make(goban, size)
	for i := range g {
		g[i] = make([]*stone, size)
	}
	return g
}

func (g goban) Place(x, y int, color int) (ok int) {
	if g[x][y] != nil {
		return NOTOK
	}
	stone := stone{x: x, y: y, color: color, goban: g}

	g[x][y] = &stone

	return OK
}

func (g goban) forEach(f func(*stone)) {
	for i := range g {
		for j := range g[i] {
			if g[i][j] != nil {
				f(g[i][j])
			}
		}
	}
}

func (g goban) String() string {
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

/*  group  */
type group []*stone

func (gb goban) NewGroup(x, y int) group {
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

func swap(s *stone) {
	switch s.color {
	case BLACK:
		s.color = WHITE
	case WHITE:
		s.color = BLACK
	}
}
