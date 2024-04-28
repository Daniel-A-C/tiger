package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

func RunSnackGame(s tcell.Screen) {
	run := true
	count1 := 0
	w, h := s.Size()
	for run {
		EmitStr(s, w/2, h/2, tcell.StyleDefault, fmt.Sprintf("hi %d", count1))
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
