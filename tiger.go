package main

import (
	"fmt"
	"os"
	"github.com/gdamore/tcell/v2"
)

func main() {
	s, e := tcell.NewScreen()
	//w, h := s.Size()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "hello")
	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:
			fmt.Print(" ", string(ev.Rune()))
			if ev.Key() == tcell.KeyEscape || string(ev.Rune()) == "q" || ev.Key() == tcell.KeyCtrlC {
				s.Fini()
				os.Exit(0)
			}
		}
	}
}
