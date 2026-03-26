package main

import (
	"helpbot/internal/config"
	"log"
)

func main() {
	log.Printf("Starting HelpBot...")
	_, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}
	log.Printf("Config loaded successfully")
}