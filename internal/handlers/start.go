package handlers


import (
	"github.com/erfjab/egobot/core"
	"github.com/erfjab/egobot/models"
)

func StartCommandHandler(b *core.Bot, update *models.Update, ctx *core.Context) error {
	_, err := b.SendMessage(&models.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Hello! I'm HelpBot, your friendly assistant.",
	})
	if err != nil {
		return err
	}
	return nil
}
