package utils

import (
	"bufio"
	"fmt"
	"strconv"

	roverDomain "github.com/MartMong/rover-planet/internal/domain/rover"
)

func ExecuteRoverWithScanner(scanner *bufio.Scanner, outputLog func(output string)) error {
	scanner.Scan()
	text := scanner.Text()
	mapSize, err := strconv.Atoi(text)
	if err != nil {
		return fmt.Errorf("invalid map size: %s", text)
	}

	rover := roverDomain.New(mapSize)
	outputLog(rover.GetState())

	for scanner.Scan() {
		text := scanner.Text()
		currentState, err := rover.Command(text)
		if err != nil {
			return err
		}

		outputLog(currentState)
	}

	return nil
}
