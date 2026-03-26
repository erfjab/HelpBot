package handlers

import (
	"context"
	"fmt"
	"html"
	"strconv"

	"helpbot/internal/database"
	"helpbot/internal/keyboards"

	"github.com/erfjab/egobot/core"
	"github.com/erfjab/egobot/models"
	"github.com/erfjab/egobot/state"
	"github.com/erfjab/egobot/tools"
)

var itemUpdateStates = state.NewStateGroup("item_update")
var stateItemUpdateAwaitTitle = itemUpdateStates.Add("AwaitTitle")
var stateItemUpdateAwaitContent = itemUpdateStates.Add("AwaitContent")

func ItemUpdateCallbackHandler(b *core.Bot, update *models.Update, ctx *core.Context) error {
	var cb keyboards.ItemUpdateCB
	if !ctx.LoadCallbackData(&cb) {
		return fmt.Errorf("failed to parse ItemUpdateCB")
	}

	userID := update.CallbackQuery.From.ID
	bg := context.Background()
	userState := b.StateManager.ForUser(userID)

	if cb.Remove {
		_, err := b.EditMessageText(&models.EditMessageTextParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.MessageID,
			Text:      "⚠️ Are you sure you want to <b>remove</b> this item?",
			ParseMode: "HTML",
			ReplyMarkup: keyboards.ConfirmRemoveKeyboard(cb.Id),
		})
		return err
	}

	if cb.Title {
		if err := userState.SetDataValue(bg, "update_item_id", strconv.Itoa(cb.Id)); err != nil {
			return err
		}
		if err := userState.SetState(bg, stateItemUpdateAwaitTitle); err != nil {
			return err
		}
		_, err := b.EditMessageText(&models.EditMessageTextParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.MessageID,
			Text:        "✏️ Send the new <b>title</b>:",
			ParseMode:   "HTML",
			ReplyMarkup: keyboards.BackItemKeyboard(cb.Id),
		})
		return err
	}

	if cb.Content {
		if err := userState.SetDataValue(bg, "update_item_id", strconv.Itoa(cb.Id)); err != nil {
			return err
		}
		if err := userState.SetState(bg, stateItemUpdateAwaitContent); err != nil {
			return err
		}
		_, err := b.EditMessageText(&models.EditMessageTextParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.MessageID,
			Text:        "📝 Send the new <b>content</b>.\n\nYou can use any formatting and premium emojis — they will be preserved.",
			ParseMode:   "HTML",
			ReplyMarkup: keyboards.BackItemKeyboard(cb.Id),
		})
		return err
	}

	return nil
}

func ItemUpdateTitleHandler(b *core.Bot, update *models.Update, ctx *core.Context) error {
	userID := update.Message.From.ID
	bg := context.Background()
	userState := b.StateManager.ForUser(userID)

	idVal, err := userState.GetDataValue(bg, "update_item_id")
	if err != nil {
		return err
	}
	idStr, _ := idVal.(string)
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return err
	}

	if err := userState.ClearAll(bg); err != nil {
		return err
	}

	item, err := database.UpdateItem(uint(id64), update.Message.Text, "")
	if err != nil {
		return err
	}

	_, err = b.SendMessage(&models.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "✅ Title updated to <b>" + html.EscapeString(item.Title) + "</b>.",
		ParseMode:   "HTML",
		ReplyMarkup: keyboards.BackItemKeyboard(int(item.Id)),
	})
	return err
}

func ItemUpdateContentHandler(b *core.Bot, update *models.Update, ctx *core.Context) error {
	userID := update.Message.From.ID
	bg := context.Background()
	userState := b.StateManager.ForUser(userID)

	idVal, err := userState.GetDataValue(bg, "update_item_id")
	if err != nil {
		return err
	}
	idStr, _ := idVal.(string)
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return err
	}

	content := tools.ParseMessageHTML(update.Message)

	if err := userState.ClearAll(bg); err != nil {
		return err
	}

	item, err := database.UpdateItem(uint(id64), "", content)
	if err != nil {
		return err
	}

	_, err = b.SendMessage(&models.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "✅ Content of <b>" + html.EscapeString(item.Title) + "</b> updated.",
		ParseMode:   "HTML",
		ReplyMarkup: keyboards.BackItemKeyboard(int(item.Id)),
	})
	return err
}

func ItemRemoveCallbackHandler(b *core.Bot, update *models.Update, ctx *core.Context) error {
	var cb keyboards.ItemRemoveCB
	if !ctx.LoadCallbackData(&cb) {
		return fmt.Errorf("failed to parse ItemRemoveCB")
	}

	if !cb.Confirm {
		return nil
	}

	if err := database.DeleteItem(uint(cb.Id)); err != nil {
		return err
	}

	_ = b.StateManager.ForUser(update.CallbackQuery.Message.From.ID).ClearAll(context.Background())

	_, err := b.EditMessageText(&models.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.MessageID,
		Text:        "🗑️ Item removed. Here are your remaining items:",
		ReplyMarkup: keyboards.BackHomeKeyboard(),
	})
	return err
}
