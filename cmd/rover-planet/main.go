package main

import (
	"fmt"
	"os"

	"github.com/MartMong/rover-planet/internal/ports/cli"
	"github.com/MartMong/rover-planet/internal/ports/file"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Printf("Error with message: %s\n", err.Error())
		}
	}()

	if len(os.Args) < 2 {
		fmt.Println("Please select mode: -cli, -file")
		return
	}

	switch os.Args[1] {
	case "-cli":
		roverCli := cli.NewRoverCLI()
		roverCli.Start()

	case "-file":
		if len(os.Args) < 3 {
			fmt.Println("Please in put file path")
			return
		}

		var roverFile *file.RoverFile
		roverFile, err = file.NewRoverFile(os.Args[2])
		if err != nil {
			return
		}

		err = roverFile.Start()
		if err != nil {
			return
		}
	}
}
