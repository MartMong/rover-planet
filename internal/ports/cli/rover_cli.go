package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	roverDomain "github.com/MartMong/rover-planet/internal/domain/rover"
)

type RoverCLI struct {
	scanner *bufio.Scanner
}

func NewRoverCLI() *RoverCLI {
	scanner := bufio.NewScanner(os.Stdin)
	return &RoverCLI{
		scanner: scanner,
	}
}

func (c *RoverCLI) Start() {
	var rover roverDomain.Interface

	fmt.Print("Please enter map size: ")
	for rover == nil {
		c.scanner.Scan()
		text := c.scanner.Text()

		mapSize, err := strconv.Atoi(text)
		if err != nil {
			fmt.Print("Please enter map size as integer: ")
			continue
		}

		rover = roverDomain.New(mapSize)
		fmt.Println(rover.GetState())
	}

	for {
		fmt.Print("Please enter command: ")

		c.scanner.Scan()
		text := c.scanner.Text()

		if text == "exit" {
			os.Exit(0)
			return
		}

		currentState, err := rover.Command(text)
		if err != nil {
			fmt.Printf("\terror: %s\n", err.Error())
		}

		fmt.Println(currentState)
	}
}
