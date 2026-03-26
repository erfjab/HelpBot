package main

import (
	"helpbot/internal/config"
	"helpbot/internal/database"
	"helpbot/internal/handlers"
	"log"

	"github.com/erfjab/egobot/core"
)

func main() {
	log.Printf("Starting HelpBot...")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}
	log.Printf("Config loaded successfully")

	log.Printf("AdminIds: %v", cfg.TelegramAdminsID)

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

	bot := core.NewBot(cfg.TelegramToken)

	getBotInfo, err := bot.GetMe()
	if err != nil {
		log.Fatalf("error getting bot info: %v", err)
	}
	log.Printf("Telegram bot initialized: @%s (ID: %d)", getBotInfo.Username, getBotInfo.ID)

	handlers.RegisterHandlers(bot)
	log.Printf("Registered handlers")

	bot.StartPolling(&core.PollingOptions{
		Timeout:    60,
		Limit:      100,
		Async:      true,
		RetryDelay: 3,
	})
}
