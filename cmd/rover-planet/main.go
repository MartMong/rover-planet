package main

import "github.com/MartMong/rover-planet/internal/domain/rover"

func main() {
	rover := rover.New(24)
	rover.TurnRight()
	rover.Forward()
	rover.TurnLeft()
	rover.Forward()
	rover.TurnLeft()
	rover.TurnLeft()
	rover.Forward()
	rover.TurnRight()
}
