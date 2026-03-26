package main

import (
	"helpbot/internal/config"
	"helpbot/internal/database"
	"log"
)

func main() {
	log.Printf("Starting HelpBot...")
	_, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}
	log.Printf("Config loaded successfully")

	err = database.LoadDB("helpbot.db")
	if err != nil {
		log.Printf("Error loading database: %v", err)
		return
	}
	log.Printf("Database loaded successfully")

	err = database.PingDB()
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return
	}
	log.Printf("Database connection successful")

	err = database.AutoMigrate()
	if err != nil {
		log.Printf("Error auto-migrating database: %v", err)
		return
	}
	log.Printf("Database auto-migration successful")

}
