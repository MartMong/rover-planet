package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/MartMong/rover-planet/internal/domain/rover"
)

type RoverCLI struct {
	rover   rover.Interface
	scanner *bufio.Scanner
}

func NewRoverCLI() *RoverCLI {
	scanner := bufio.NewScanner(os.Stdin)
	return &RoverCLI{
		scanner: scanner,
	}
}

func (c *RoverCLI) Start() {
	fmt.Print("Please enter map size: ")
	for c.rover == nil {
		c.scanner.Scan()
		text := c.scanner.Text()

		mapSize, err := strconv.Atoi(text)
		if err != nil {
			fmt.Print("Please enter map size as integer: ")
			continue
		}

		c.rover = rover.New(mapSize)
		fmt.Println(c.rover.GetState())
	}

	for {
		fmt.Print("Please enter command: ")

		c.scanner.Scan()
		text := c.scanner.Text()

		if text == "exit" {
			os.Exit(0)
			return
		}

		currentState, err := c.rover.Command(text)
		if err != nil {
			fmt.Printf("\terror: %s\n", err.Error())
		}

		fmt.Println(currentState)
	}
}
