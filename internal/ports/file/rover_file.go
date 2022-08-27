package file

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MartMong/rover-planet/internal/utils"
)

type RoverFile struct {
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
	return utils.ExecuteRoverWithScanner(c.scanner, func(output string) {
		fmt.Println(output)
	})
}
