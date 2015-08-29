package main

import (
	"fmt"
	"time"

	"strconv"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/gxfont"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"

	// . "gobang/consts"
	"github.com/Drahoslav7/gobang/logic"
)

// font ids
type fontId int

const (
	H1 fontId = -iota
	H2
	H3
	NORMAL
	CODE
)

/* Theme type */

type Theme struct {
	gxui.Theme
	fonts   map[fontId]gxui.Font
	bgColor gxui.Color
}

func createTheme(theme gxui.Theme) Theme {
	t := Theme{Theme: theme, fonts: make(map[fontId]gxui.Font)}

	t.addFont(H1, gxfont.Default, 36)
	t.addFont(H2, gxfont.Default, 24)
	t.addFont(H3, gxfont.Default, 18)
	t.addFont(NORMAL, gxfont.Default, 14)
	t.addFont(CODE, gxfont.Monospace, 13)

	return t
}

func (t *Theme) addFont(id fontId, bytes []byte, size int) {
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
	texts := []struct {
		fontId
		text string
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
		label := theme.CreateLabel()
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
	label.SetText("I am")

	input := theme.CreateTextBox()
	input.SetFont(theme.fonts[H2])
	input.SetText("")
	input.OnAttach(func() {
		gxui.SetFocus(input)
	})

	/* buttons */
	buttext := theme.CreateLabel()
	buttext.SetFont(theme.fonts[H3])
	buttext.SetText("Let me in!")

	buttonLogIn := theme.CreateButton()
	buttonLogIn.AddChild(buttext)

	buttonPlaySingle := theme.CreateButton()
	// buttonPlaySingle.SetFont(theme.fonts[H3])
	buttonPlaySingle.SetText("Practice")
	buttonPlaySingle.OnClick(func(_ gxui.MouseEvent) {
		fmt.Println("single clicked")
		app.Practice(nil)
	})

	screen.AddChild(label)
	screen.AddChild(input)
	screen.AddChild(buttonLogIn)
	screen.AddChild(buttonPlaySingle)

	return screen
}

func createSettings(theme Theme) gxui.LinearLayout {
	screen := createScreen(theme)

	label := theme.CreateLabel()
	label.SetText("size")

	input := theme.CreateTextBox()
	input.SetText("9")
	// input.OnAttach(func() {
	// 	gxui.SetFocus(input)
	// })

	button := theme.CreateButton()
	button.SetText("Play!")
	button.OnClick(func(_ gxui.MouseEvent) {
		size, _ := strconv.Atoi(input.Text())
		app.Play(size, nil)
	})

	screen.AddChild(label)
	screen.AddChild(input)
	screen.AddChild(button)

	return screen
}

func createPlayground(theme Theme) gxui.LinearLayout {
	screen := createScreen(theme)

	label := theme.CreateLabel()
	label.SetText("game") // TODO implement whole screen

	screen.AddChild(label)

	return screen
}

// app start point
func appMain(driver gxui.Driver) {

	theme := createTheme(dark.CreateTheme(driver))
	theme.bgColor = gxui.Color{0.25, 0.25, 0.5, 1}

	window := theme.CreateWindow(400, 350, APPNAME)
	window.SetBackgroundBrush(gxui.CreateBrush(theme.bgColor))

	window.OnClose(func() {
		driver.Terminate()
		app.Exit()
	})

	// fullscreen
	window.OnKeyDown(func(e gxui.KeyboardEvent) {
		if e.Key == gxui.KeyF11 {
			window.SetFullscreen(!window.Fullscreen())
		}
	})

	////////////////
	//  screens
	show := func(screen gxui.LinearLayout) {
		window.RemoveChild(screen)
		window.AddChild(screen)
		screen.SetVisible(true)
	}
	hide := func(screen gxui.LinearLayout) {
		screen.SetVisible(false)
	}

	///// playground screen
	playground := createPlayground(theme)
	window.AddChild(playground)
	app.OnState(logic.GAME_SINGLE, func() {
		show(playground)
	})

	///// settings screen
	settings := createSettings(theme)
	window.AddChild(settings)
	app.OnState(logic.PRACTICE_SETTINGS, func() {
		show(settings)
	})
	app.OnState(logic.DUEL_SETTINGS, func() {
		show(settings)
	})

	///// login screen
	entrance := createEntrance(theme)
	window.AddChild(entrance)
	func() {
		app.OnState(logic.ENTRANCE, func() {
			driver.Call(func() {
				show(entrance)
			})
		})
		enter := func() {
			app.Entrance(nil)
		}
		// hide after some time
		time.AfterFunc(2*time.Second, enter)
		// or hide on click
		window.OnClick(func(_ gxui.MouseEvent) {
			enter()
		})
		// or hide on esc
		window.OnKeyDown(func(e gxui.KeyboardEvent) {
			if e.Key == gxui.KeyEscape {
				enter()
			}
		})

	}()

	///// intro screen
	intro := createIntro(theme)
	window.AddChild(intro)
	intro.SetVisible(true)
	// show if F1 press
	window.OnKeyDown(func(e gxui.KeyboardEvent) {
		if e.Key == gxui.KeyF1 {
			show(intro)
		}
	})
	// hide on F1 release
	window.OnKeyUp(func(e gxui.KeyboardEvent) {
		if e.Key == gxui.KeyF1 {
			hide(intro)
		}
	})

}

func startGUI() {
	fmt.Println("GUI start")
	defer fmt.Println("GUI end")
	gl.StartDriver(appMain)
}
