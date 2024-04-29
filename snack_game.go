package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Position struct {
    X int 
    Y int 
}

func updatePlayerPosition(player *Position, dx, dy int, s tcell.Screen) {
    EmitStr(s, player.X, player.Y, tcell.StyleDefault, " ")
    player.X += dx
    player.Y += dy
    EmitStr(s, player.X, player.Y, tcell.StyleDefault, "B")
    s.Show()
}

func RunSnackGame(s tcell.Screen) {
	const updateTime = 16666666 // The amount of time the game loop needs to wait in nanoseconds

	player := Position{X: 10, Y: 10}

	run := true

	//w, h := s.Size()
	EmitStr(s, player.X, player.Y, tcell.StyleDefault, fmt.Sprintf("B"))
	s.Show()

	var gameTimer, updateTimer Timer
	gameTimer.StartTimer()
	updateTimer.StartTimer()


	for run {

		if updateTimer.ElapsedTime() > updateTime {
			updateTimer.StartTimer()

			if s.HasPendingEvent() {
				switch ev := s.PollEvent().(type) {
				case *tcell.EventKey:
					if ev.Key() == tcell.KeyEscape || string(ev.Rune()) == "q" || ev.Key() == tcell.KeyCtrlC {
						run = false
					} else if ev.Key() == tcell.KeyF3 { // Debug information display
						EmitStr(s, 0, 0, tcell.StyleDefault, fmt.Sprintf("Game Time: %d", (gameTimer.ElapsedTime() / time.Second)))
						//EmitStr(s, 20, 0, tcell.StyleDefault, fmt.Sprintf("Updates: %d", updates / 60))
						s.Show()
					} else if string(ev.Rune()) == "h" {
						updatePlayerPosition(&player, -2,  0, s)
					} else if string(ev.Rune()) == "j" {
						updatePlayerPosition(&player,  0,  1, s)
					} else if string(ev.Rune()) == "k" {
						updatePlayerPosition(&player,  0, -1, s)
					} else if string(ev.Rune()) == "l" {
						updatePlayerPosition(&player,  2,  0, s)
					}
				}
			}
		}

	}

}
