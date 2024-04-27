package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

func RunCounter(s tcell.Screen) { 
	run := true
	count1 := 0
	w, h := s.Size()
	for run {
		EmitStr(s, w/2, h/2, tcell.StyleDefault, fmt.Sprintf("%d", count1))
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

func RunWave(s tcell.Screen) { 
	run := true
	count1 := 0
	period := 10
	wavePos := 0
	w, h := s.Size()
	waveChar := "*"
	for run {
		count1 += 1

		if count1 >= period {
			count1 = 0
			
			for i := wavePos; i > 0; i-- {
				EmitStr(s, i, wavePos-i, tcell.StyleDefault, waveChar)
			}
			s.Show()

			wavePos += 1
			if wavePos >= w + h {
				wavePos = 0
				if waveChar == "*" {
					waveChar = "-"
				} else {
					waveChar = "*"
				}
			}
		}

		if s.HasPendingEvent() {
			switch ev := s.PollEvent().(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || string(ev.Rune()) == "q" || ev.Key() == tcell.KeyCtrlC {
					run = false
				}
			}
		}

		time.Sleep(1 * time.Millisecond)

	}
}

func TestCorners(s tcell.Screen) {
	s.Clear()
	w, h := s.Size()
	EmitStr(s, 0, 0, tcell.StyleDefault, "1")
	EmitStr(s, w-1, 0, tcell.StyleDefault, "2")
	EmitStr(s, 0, h-1, tcell.StyleDefault, "3")
	EmitStr(s, w-1, h-1, tcell.StyleDefault, "4")
	s.Show()
}

func RunTestCorners(s tcell.Screen) { 
	TestCorners(s)
	run := true
	for run {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || string(ev.Rune()) == "q" || ev.Key() == tcell.KeyCtrlC {
				run = false
			}
			if string(ev.Rune()) == "p" {
				EmitStr(s, 0, 10, tcell.StyleDefault, "Hello!")
				s.Show()
			}
		case *tcell.EventResize:
			TestCorners(s)
		}

	}
}
