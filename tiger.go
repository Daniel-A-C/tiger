package main

import (
	"fmt"
	"os"
	"github.com/gdamore/tcell/v2"

	"github.com/mattn/go-runewidth"
)

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func testCorners(s tcell.Screen) {
	s.Clear()
	w, h := s.Size()
	emitStr(s, 0, 0, tcell.StyleDefault, "1")
	emitStr(s, w-1, 0, tcell.StyleDefault, "2")
	emitStr(s, 0, h-1, tcell.StyleDefault, "3")
	emitStr(s, w-1, h-1, tcell.StyleDefault, "4")
	s.Show()
}

func initScreen() (sc tcell.Screen) {
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	s.Clear()

	return s
}

func runTestCorners(s tcell.Screen) { 
	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:
			//fmt.Print(" ", string(ev.Rune()))
			if ev.Key() == tcell.KeyEscape || string(ev.Rune()) == "q" || ev.Key() == tcell.KeyCtrlC {
				s.Fini()
				os.Exit(0)
			}
			if string(ev.Rune()) == "p" {
				emitStr(s, 0, 10, tcell.StyleDefault, "Hello!")
				s.Show()
			}
		case *tcell.EventResize:
			testCorners(s)
		}

	}
}

func main() {
	s := initScreen()
	testCorners(s)
	runTestCorners(s)
}


//boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple) defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

//emitStr(s, 10, 10, defStyle, "Test")
//s.Show()

//time.Sleep(3 * time.Second)
