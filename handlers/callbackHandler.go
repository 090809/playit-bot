package handlers

import (
	"log"
	"strings"

	tb "github.com/090809/telebot"
	"playit-bot/buttons"
	"playit-bot/user"
	"playit-bot/utils"
)

var welcomeText = "Добро пожаловать в систему, юный хацкер!\n\n" +
	"Теперь ты имеешь доступ к прохождению наших заданий."

type CallbackHandler struct {
	Handler
}

func (h *CallbackHandler) Handle(c *tb.Callback) {
	switch strings.TrimSpace(c.Data) {
	case "login-ok":
		h.loginOk(c)
	case "login-restart":
		h.loginRestart(c)
	}
}

func (h *CallbackHandler) loginOk(c *tb.Callback) {
	replyButtons := buttons.MainReplyButtons

	if err := h.bot.Delete(c.Message); err != nil {
		log.Printf("[ERROR] %v", err)
	}

	if _, err := h.bot.Send(c.Sender, welcomeText, &tb.ReplyMarkup{
		ReplyKeyboard: replyButtons,
	}); err != nil {
		log.Printf("[ERROR] %v", err)
	}

	u := h.repository.Find(c.Sender.Recipient())
	u.Phase = user.MainPhase
	h.repository.Save(u)

	if err := utils.Confirm("tag", *u.HashTag); err != nil {
		log.Printf("[ERROR] %v", err)
	}
}

func (h *CallbackHandler) loginRestart(c *tb.Callback) {
	_, err := h.bot.Edit(c.Message, "Введи хэш-тег в формате \"#12345\".\nХэш-тег можно получить у ребят на стенде.")
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}

	u := h.repository.Find(c.Sender.Recipient())
	u.Phase = user.LoggingPhase
	h.repository.Save(u)
}
