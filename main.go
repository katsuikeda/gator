package main

import (
	"fmt"
	"log"

	"github.com/katsuikeda/gator/internal/config"
)

const userName = "katsu"

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser(userName)
	if err != nil {
		log.Fatalf("error setting the current user name: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	fmt.Printf("Read config again: %+v\n", cfg)
}
