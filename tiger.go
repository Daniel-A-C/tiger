package main

import (
	"os"
)

func main() {
	s := InitScreen()

	//RunCounter(s)
	RunTestCorners(s)

	s.Fini()
	os.Exit(0)
}
