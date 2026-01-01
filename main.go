package main

import (
	"fmt"
	"gator/internal/config"
	"os"
)

func main() {
	fmt.Println("Starting gator...")

	// Read Config
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config", err)
		os.Exit(1)
	}

	s := state{
		cfg: &cfg,
	}

	myCommands := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	myCommands.register("login", handlerLogin)

	cmdLineInput := os.Args
	if len(cmdLineInput) < 2 {
		fmt.Println("Requires a command name")
		os.Exit(1)
	}

	myCommand := command{
		name: cmdLineInput[1],
		args: cmdLineInput[2:],
	}

	err = myCommands.run(&s, myCommand)
	if err != nil {
		fmt.Println("error running command", err)
		os.Exit(1)
	}
}
