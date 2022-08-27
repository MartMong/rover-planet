package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/MartMong/rover-planet/internal/domain/rover"
)

type RoverFile struct {
	rover   rover.Interface
	scanner *bufio.Scanner
	file    *os.File
}

func NewRoverFile(path string) (*RoverFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	return &RoverFile{
		file:    file,
		scanner: scanner,
	}, nil
}

func (c *RoverFile) Start() error {
	c.scanner.Scan()
	text := c.scanner.Text()

	mapSize, err := strconv.Atoi(text)
	if err != nil {
		return fmt.Errorf("invalid map size: %s", text)
	}

	c.rover = rover.New(mapSize)
	fmt.Println(c.rover.GetState())

	for c.scanner.Scan() {
		text := c.scanner.Text()
		currentState, err := c.rover.Command(text)
		if err != nil {
			return err
		}

		fmt.Println(currentState)
	}

	return nil
}
