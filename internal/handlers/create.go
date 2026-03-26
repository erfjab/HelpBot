package handlers

import (
	"context"
	"html"

	"helpbot/internal/database"
	"helpbot/internal/keyboards"

	"github.com/erfjab/egobot/core"
	"github.com/erfjab/egobot/models"
	"github.com/erfjab/egobot/state"
	"github.com/erfjab/egobot/tools"
)

var itemCreateStates = state.NewStateGroup("item_create")
var stateItemCreateAwaitTitle = itemCreateStates.Add("AwaitTitle")
var stateItemCreateAwaitContent = itemCreateStates.Add("AwaitContent")

func ItemCreateCallbackHandler(b *core.Bot, update *models.Update, ctx *core.Context) error {
	userID := update.CallbackQuery.From.ID
	bg := context.Background()

	if err := b.StateManager.ForUser(userID).SetState(bg, stateItemCreateAwaitTitle); err != nil {
		return err
	}

	_, err := b.EditMessageText(&models.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.MessageID,
		Text:        "✏️ Send the item <b>title</b>:",
		ParseMode:   "HTML",
		ReplyMarkup: keyboards.CancelKeyboard(),
	})
	return err
}

func ItemCreateTitleHandler(b *core.Bot, update *models.Update, ctx *core.Context) error {
	userID := update.Message.From.ID
	bg := context.Background()

	title := update.Message.Text
	userState := b.StateManager.ForUser(userID)

	if err := userState.SetDataValue(bg, "create_title", title); err != nil {
		return err
	}
	if err := userState.SetState(bg, stateItemCreateAwaitContent); err != nil {
		return err
	}

	_, err := b.SendMessage(&models.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "📝 Now send the item <b>content</b>.\n\nYou can use any formatting and premium emojis — they will be preserved.",
		ParseMode:   "HTML",
		ReplyMarkup: keyboards.CancelKeyboard(),
	})
	return err
}

func ItemCreateContentHandler(b *core.Bot, update *models.Update, ctx *core.Context) error {
	userID := update.Message.From.ID
	bg := context.Background()

	userState := b.StateManager.ForUser(userID)

	titleVal, err := userState.GetDataValue(bg, "create_title")
	if err != nil {
		return err
	}
	title, _ := titleVal.(string)

	content := tools.ParseMessageHTML(update.Message)

	if err := userState.ClearAll(bg); err != nil {
		return err
	}

	item, err := database.CreateItem(title, content)
	if err != nil {
		return err
	}

	_, err = b.SendMessage(&models.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "✅ Item <b>" + html.EscapeString(item.Title) + "</b> created successfully!",
		ParseMode:   "HTML",
		ReplyMarkup: keyboards.BackHomeKeyboard(),
	})
	return err
}


