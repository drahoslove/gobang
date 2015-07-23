// go game in go lang

package main

import (
	"fmt"

	"time"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/gxfont"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"

	// . "gobang/consts"
	"gobang/logic"
)

// GLOBALS
const (
	VERSION     = "0.1"
	APPNAME     = "Gobang!"
	DESCRIPTION = "Goban in Golang"
	AUTHOR      = "Drahoslav Bednář"
	WEBPAGE     = "go.yo2.cz"
)

// font ids
const (
	H1 = -(1 + iota)
	H2
	H3
	NORMAL
	CODE
)

var (
	app logic.Application // controler
)

/* Theme type */

type Theme struct {
	gxui.Theme
	fonts   map[int]gxui.Font
	bgColor gxui.Color
}

func createTheme(theme gxui.Theme) Theme {
	t := Theme{Theme: theme, fonts: make(map[int]gxui.Font)}

	t.addFont(H1, gxfont.Default, 36)
	t.addFont(H2, gxfont.Default, 24)
	t.addFont(H3, gxfont.Default, 18)
	t.addFont(NORMAL, gxfont.Default, 14)
	t.addFont(CODE, gxfont.Monospace, 13)

	return t
}

func (t *Theme) addFont(id int, bytes []byte, size int) {
	font, err := t.Driver().CreateFont(bytes, size)
	if err != nil {
		panic(err)
	}
	t.fonts[id] = font
}

/* gui componennts */

func createScreen(theme Theme) gxui.LinearLayout {
	screen := theme.CreateLinearLayout()
	screen.SetVisible(false)
	screen.SetSizeMode(gxui.Fill)
	screen.SetDirection(gxui.TopToBottom)
	screen.SetHorizontalAlignment(gxui.AlignCenter)
	screen.SetPadding(math.CreateSpacing(5))
	screen.SetBackgroundBrush(gxui.CreateBrush(theme.bgColor))
	return screen
}

func createIntro(theme Theme) gxui.LinearLayout {
	screen := createScreen(theme)

	// set texts
	var label gxui.Label

	texts := []struct {
		fontId int
		text   string
	}{
		{H1, APPNAME},
		{H3, DESCRIPTION},
		{CODE, fmt.Sprintf(
			"ver %s\n"+
				"\n"+
				"%s\n"+
				"\n"+
				"© 2015 %s. All rights reserved\n"+
				"And some other stuff to make this look pretty\n"+
				"interesting. Lorem ipsum, also I love you!\n"+
				"",
			VERSION, WEBPAGE, AUTHOR,
		)},
	}

	for _, s := range texts {
		label = theme.CreateLabel()
		label.SetMultiline(true)
		font := theme.fonts[s.fontId]
		label.SetFont(font)
		M := font.Size()
		label.SetMargin(math.Spacing{0, M * 3 / 2, M / 2, 0})
		label.SetText(s.text)
		screen.AddChild(label)
	}

	return screen
}

// func createMenu(theme Theme) gxui.LinearLayout {
// 	menu := createScreen(theme)

// 	label := theme.CreateLabel()
// 	label.SetText("choose")

// 	menu.AddChild(label)

// 	return menu
// }

func createEntrance(theme Theme) gxui.LinearLayout {
	screen := createScreen(theme)

	label := theme.CreateLabel()
	label.SetFont(theme.fonts[H3])
	label.SetText("I am ")

	input := theme.CreateTextBox()
	input.SetFont(theme.fonts[H2])
	input.SetText("")
	input.OnAttach(func() {
		gxui.SetFocus(input)
	})

	buttext := theme.CreateLabel()
	buttext.SetFont(theme.fonts[H3])
	buttext.SetText("Let me in!")

	button := theme.CreateButton()
	button.AddChild(buttext)

	screen.AddChild(label)
	screen.AddChild(input)
	screen.AddChild(button)

	return screen
}

// app start point
func appMain(driver gxui.Driver) {
	theme := createTheme(dark.CreateTheme(driver))
	theme.bgColor = gxui.Color{0.25, 0.25, 0.5, 1}

	window := theme.CreateWindow(400, 350, APPNAME)
	window.SetBackgroundBrush(gxui.CreateBrush(theme.bgColor))

	window.OnClose(driver.Terminate)
	window.OnKeyDown(func(e gxui.KeyboardEvent) {
		if e.Key == gxui.KeyF11 {
			window.SetFullscreen(!window.Fullscreen())
		}
	})

	// screens //

	// login screen
	entrance := createEntrance(theme)
	window.AddChild(entrance)

	// intro view
	intro := createIntro(theme)
	intro.SetVisible(true)
	window.AddChild(intro)

	// intro actions rules
	func() {

		enter := func() {
			app.Entrance(func(_ interface{}) {
				intro.SetVisible(false)
				entrance.SetVisible(true)
				// show if F1 press
				window.OnKeyDown(func(e gxui.KeyboardEvent) {
					if e.Key == gxui.KeyF1 {
						intro.SetVisible(true)
					}
				})
				// hide on F1 release
				window.OnKeyUp(func(e gxui.KeyboardEvent) {
					if e.Key == gxui.KeyF1 {
						intro.SetVisible(false)
					}
				})
			})
		}
		// hide after some time
		time.AfterFunc(2*time.Second, enter)
		// or hide on click
		intro.OnClick(func(_ gxui.MouseEvent) {
			enter()
		})
		// or hide on esc
		window.OnKeyDown(func(e gxui.KeyboardEvent) {
			if e.Key == gxui.KeyEscape {
				enter()
			}
		})

	}()

}

func main() {

	fmt.Println("start")
	gl.StartDriver(appMain)
	fmt.Println("end")

}
