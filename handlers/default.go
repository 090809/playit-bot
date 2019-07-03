package handlers

import (
	"context"

	"github.com/090809/telebot"
	"playit-bot/user"
)

type Handler struct {
	bot     *telebot.Bot
	context context.Context
	repository *user.Repository
}

func NewHandler(bot *telebot.Bot, context context.Context, repository *user.Repository) *Handler {
	return &Handler{bot: bot, context: context, repository: repository}
}

