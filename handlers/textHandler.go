package handlers

import (
	"errors"
	"log"
	"regexp"

	tb "github.com/090809/telebot"
	"playit-bot/user"
)

var notValidated = errors.New("token not validated")
var notParsed = errors.New("command not parsed")

type TextHandler struct {
	Handler
}

func (h *TextHandler) Handle(m *tb.Message) error {
	userId := m.Sender.Recipient()

	u := h.repository.Find(userId); if u == nil {
		if h.bot != nil {
			if _, err := h.bot.Send(m.Sender, "Отправьте /start для начала работы с ботом!"); err != nil {
				log.Fatalln(err)
				return nil
			}
		}
	}

	switch u.Phase {
	case user.LoggingPhase:
		if err := h.HandleLogging(m, u); err != nil {
			return err
		}
		return nil
	}

	if h.bot != nil {
		if _, err := h.bot.Send(m.Sender, "Я не распознал вашу команду"); err != nil {
			log.Fatalln(err)
		}
	}
	return notParsed
}

func (h *TextHandler) HandleLogging(m *tb.Message, u *user.User) error {
	text := m.Text; if !h.validateLoggingText(text) {
		if h.bot != nil {
			if _, err := h.bot.Send(m.Sender, "Необходимо ввести правильный токен. Пример: \"#000000\""); err != nil {
				log.Fatalln(err)
			}
		}
		return notValidated
	}

	u.HashTag = &text
	u.Phase = user.MainPhase
	h.repository.Save(u)
	return nil
}

func (h *TextHandler) validateLoggingText(text string) bool {
	var validHash = regexp.MustCompile(`[#][0-9]{6}`)
	return validHash.MatchString(text)
}