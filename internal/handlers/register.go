package handlers

import (
	Middlewares "helpbot/internal/middlewares"

	"github.com/erfjab/egobot/core"
	"github.com/erfjab/egobot/state"
)

func RegisterHandlers(b *core.Bot) {
	userGroup := core.NewHandlerGroup("handlers")
	userGroup.UseMiddleware(Middlewares.AuthMiddlewares)
	userGroup.OnCommand("start", StartCommandHandler, state.IgnoreState())
	b.RegisterGroup(userGroup)
}