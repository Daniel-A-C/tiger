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
	testCorners(s)
	run := true
	for run {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || string(ev.Rune()) == "q" || ev.Key() == tcell.KeyCtrlC {
				run = false
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

func runFunVisual(s tcell.Screen) { 
	run := true
	count1 := 0
	w, h := s.Size()
	for run {
		emitStr(s, w/2, h/2, tcell.StyleDefault, fmt.Sprintf("%d", count1))
		s.Show()
		count1 += 1

		if s.HasPendingEvent() {
			switch ev := s.PollEvent().(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || string(ev.Rune()) == "q" || ev.Key() == tcell.KeyCtrlC {
					run = false
				}
			}
		}

	}
}

func main() {
	s := initScreen()

	runFunVisual(s)

	s.Fini()
	os.Exit(0)
}
