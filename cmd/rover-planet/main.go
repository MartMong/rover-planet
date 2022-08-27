package main

import (
	"os"
	"strings"

	"github.com/MartMong/rover-planet/internal/ports/cli"
)

func main() {
	option := strings.Join(os.Args[1:], ",")
	switch option {
	case "-cli":
		roverCli := cli.NewRoverCLI()
		roverCli.Start()
	}
}
