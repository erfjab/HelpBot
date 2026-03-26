package handlers

import (
	"helpbot/internal/keyboards"
	Middlewares "helpbot/internal/middlewares"

	"github.com/erfjab/egobot/core"
	"github.com/erfjab/egobot/state"
)

func RegisterHandlers(b *core.Bot) {
	userGroup := core.NewHandlerGroup("handlers")
	userGroup.UseMiddleware(Middlewares.AuthMiddlewares)

	userGroup.OnCommand("start", StartCommandHandler, state.IgnoreState())
	userGroup.OnCallbackData("home", HomeCallbackHandler, state.IgnoreState())

	userGroup.OnCallbackStruct(keyboards.ItemInfoCB{}, ItemInfoCallbackHandler, state.IgnoreState())

	userGroup.OnCallbackStruct(keyboards.ItemCreateCB{}, ItemCreateCallbackHandler, state.IgnoreState())
	userGroup.OnText(ItemCreateTitleHandler, state.InState(stateItemCreateAwaitTitle))
	userGroup.OnText(ItemCreateContentHandler, state.InState(stateItemCreateAwaitContent))

	userGroup.OnCallbackStruct(keyboards.ItemUpdateCB{}, ItemUpdateCallbackHandler, state.IgnoreState())
	userGroup.OnText(ItemUpdateTitleHandler, state.InState(stateItemUpdateAwaitTitle))
	userGroup.OnText(ItemUpdateContentHandler, state.InState(stateItemUpdateAwaitContent))

	userGroup.OnCallbackStruct(keyboards.ItemRemoveCB{}, ItemRemoveCallbackHandler, state.IgnoreState())

	b.RegisterGroup(userGroup)
}