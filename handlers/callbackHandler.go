package handlers

import (
	tb "github.com/090809/telebot"
)

type CallbackHandler struct {
	Handler
}

func (h *CallbackHandler) Handle(m *tb.Message) {
	_, _ = h.bot.Send(m.Sender, "Amma callback handler")
}
