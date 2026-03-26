package handlers

import (
	"context"
	"fmt"
	"html"

	"helpbot/internal/database"
	"helpbot/internal/keyboards"

	"github.com/erfjab/egobot/core"
	"github.com/erfjab/egobot/models"
)

func ItemInfoCallbackHandler(b *core.Bot, update *models.Update, ctx *core.Context) error {
	_ = b.StateManager.ForUser(update.CallbackQuery.Message.From.ID).ClearAll(context.Background())

	var cb keyboards.ItemInfoCB
	if !ctx.LoadCallbackData(&cb) {
		return fmt.Errorf("failed to parse ItemInfoCB")
	}

	item, err := database.GetItemByID(uint(cb.Id))
	if err != nil {
		return err
	}

	text := "<b>Title:</b> " + html.EscapeString(item.Title) + "\n\n<b>Content:</b> " + item.Content

	_, err = b.EditMessageText(&models.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.MessageID,
		Text:        text,
		ParseMode:   "HTML",
		ReplyMarkup: keyboards.UpdateItemKeyboard(cb.Id),
	})
	return err
}
