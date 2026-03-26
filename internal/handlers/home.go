package handlers

import (
	"context"

	"helpbot/internal/database"
	"helpbot/internal/keyboards"

	"github.com/erfjab/egobot/core"
	"github.com/erfjab/egobot/models"
)

func StartCommandHandler(b *core.Bot, update *models.Update, ctx *core.Context) error {
	_ = b.StateManager.ForUser(update.Message.From.ID).ClearAll(context.Background())

	items, err := database.GetAllItems("")
	if err != nil {
		return err
	}
	_, err = b.SendMessage(&models.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Hello! I'm HelpBot, your friendly assistant.",
		ReplyMarkup: keyboards.HomeKeyboard(items),
	})
	if err != nil {
		return err
	}
	return nil
}

func HomeCallbackHandler(b *core.Bot, update *models.Update, ctx *core.Context) error {
	_ = b.StateManager.ForUser(update.CallbackQuery.From.ID).ClearAll(context.Background())

	items, err := database.GetAllItems("")
	if err != nil {
		return err
	}
	_, err = b.EditMessageText(&models.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.MessageID,
		Text:        "Welcome back to the home screen! Here are your items:",
		ReplyMarkup: keyboards.HomeKeyboard(items),
	})
	if err != nil {
		return err
	}
	return nil
}
