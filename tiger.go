package main

import (
	"os"
)

func main() {
	s := InitScreen()

	//RunWave(s)
	//RunCounter(s)
	//RunTestCorners(s)
	RunSnackGame(s)


	s.Fini()
	os.Exit(0)
}
