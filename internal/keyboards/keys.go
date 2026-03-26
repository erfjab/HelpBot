package keyboards

import (
	"helpbot/internal/database"

	"github.com/erfjab/egobot/models"
	"github.com/erfjab/egobot/tools"
)

func HomeKeyboard(items []database.Items) *models.InlineKeyboardMarkup {
	kb := tools.NewInlineKeyboard()
	for _, item := range items {
		kb.Row(
			tools.MustCallbackButton(item.Title, ItemInfoCB{
				Id: int(item.Id),
			}),
		)
	}
	kb.Row(
		tools.MustCallbackButton("➕ Create Item", ItemCreateCB{}),
	)
	return kb.Build()
}

func BackHomeKeyboard() *models.InlineKeyboardMarkup {
	kb := tools.NewInlineKeyboard()
	kb.Row(
		tools.Button("🔙 Back to Home", "home"),
	)
	return kb.Build()
}

func CancelKeyboard() *models.InlineKeyboardMarkup {
	kb := tools.NewInlineKeyboard()
	kb.Row(
		tools.Button("❌ Cancel", "home"),
	)
	return kb.Build()
}

func BackItemKeyboard(itemId int) *models.InlineKeyboardMarkup {
	kb := tools.NewInlineKeyboard()
	kb.Row(
		tools.MustCallbackButton("🔙 Back to Item", ItemInfoCB{
			Id: itemId,
		}),
	)
	return kb.Build()
}

func UpdateItemKeyboard(itemId int) *models.InlineKeyboardMarkup {
	kb := tools.NewInlineKeyboard()
	kb.Row(
		tools.MustCallbackButton("✏️ Update Title", ItemUpdateCB{
			Id:    itemId,
			Title: true,
		}),
	)
	kb.Row(
		tools.MustCallbackButton("✏️ Update Content", ItemUpdateCB{
			Id:      itemId,
			Content: true,
		}),
	)
	kb.Row(
		tools.MustCallbackButton("🗑️ Remove Item", ItemUpdateCB{
			Id: itemId,
			Remove: true,
		}),
	)
	kb.Row(
		tools.Button("🔙 Back to Home", "home"),
	)
	return kb.Build()
}

func ConfirmRemoveKeyboard(itemId int) *models.InlineKeyboardMarkup {
	kb := tools.NewInlineKeyboard()
	kb.Row(
		tools.MustCallbackButton("✅ Confirm Remove", ItemRemoveCB{
			Id:      itemId,
			Confirm: true,
		}),
	)
	kb.Row(
		tools.MustCallbackButton("❌ Cancel", ItemInfoCB{
			Id: itemId,
		}),
	)
	return kb.Build()
}