package handlers

import (
	"log"

	tb "github.com/090809/telebot"
	"playit-bot/user"
)

type StartHandler struct {
	Handler
}

const startMessage = "Привет, пришли мне свой хештэг-номер, который ты получил ранее, и мы приступим к заданиям =)"

func (h *StartHandler) Handle(m *tb.Message) {
	userId := m.Sender.Recipient()

	h.ProcessUser(userId)

	if _, err := h.bot.Send(m.Sender, startMessage); err != nil {
		log.Fatalln(err)
	}
}

func (h *StartHandler) ProcessUser(userId string) {
	var u *user.User

	u = h.repository.Find(userId); if u == nil {
		u = user.NewUser(userId, user.LoggingPhase, nil)
	}
	u.Phase = user.LoggingPhase
	h.repository.Save(u)
}
