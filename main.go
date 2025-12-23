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

	//set current user and update config file on disk
	//if err := cfg.SetUser("charlie"); err != nil {
	//	fmt.Println("error setting user", err)
	//}

	// read config again and print contents of struct to terminal
	//updatedCfg, err := config.Read()
	//if err != nil {
	//	fmt.Println("error reading updated config", err)
	//}

	fmt.Println(s)
}
