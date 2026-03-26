package Middlewares

import (
	"log"

	"helpbot/internal/config"

	"github.com/erfjab/egobot/core"
	"github.com/erfjab/egobot/models"
)

func AuthMiddlewares(b *core.Bot, update *models.Update, ctx *core.Context, next core.NextFunc) {
	user := extractUser(update)
	if user == nil {
		log.Printf("No user information found in the update")
		return
	}

	if !isAdmin(user.ID) {
		log.Printf("Unauthorized access attempt by user %d", user.ID)
		return
	}
	next()
}

func extractUser(update *models.Update) *models.User {
	if update.Message != nil {
		return update.Message.From
	} else if update.CallbackQuery != nil {
		return &update.CallbackQuery.From
	}
	return nil
}

func isAdmin(userID int64) bool {
	for _, adminID := range config.Cfg.TelegramAdminsID {
		if userID == adminID {
			return true
		}
	}
	return false
}