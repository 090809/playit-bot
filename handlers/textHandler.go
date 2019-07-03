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

func (h *TextHandler) Handle(m *tb.Message) {
	log.Print("TextHandler")

	userId := m.Sender.Recipient()

	log.Printf("UserID: %v", userId)

	u := h.repository.Find(userId)
	if u == nil {
		if h.bot != nil {
			if _, err := h.bot.Send(m.Sender, "Отправьте /start для начала работы с ботом!"); err != nil {
				log.Fatalln(err)
				return
			}
		}
	}

	switch u.Phase {
	case user.LoggingPhase:
		h.HandleLogging(m, u)
	case user.MainPhase:
		h.HandleMain(m, u)
	default:
		if h.bot != nil {
			if _, err := h.bot.Send(m.Sender, "Я не распознал вашу команду"); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func (h *TextHandler) HandleLogging(m *tb.Message, u *user.User) {
	text := m.Text
	if !h.validateLoggingText(text) {
		if h.bot != nil {
			if _, err := h.bot.Send(m.Sender, "Необходимо ввести правильный код. Пример: \"#000000\""); err != nil {
				log.Fatalln(err)
			}
		}
	}

	u.HashTag = &text
	u.Phase = user.MainPhase
	h.repository.Save(u)
	if _, err := h.bot.Send(m.Sender, "Код принят, теперь можно развлекаться"); err != nil {
		log.Fatalln(err)
	}
}

func (h *TextHandler) validateLoggingText(text string) bool {
	var validHash = regexp.MustCompile(`[#][0-9]{6}`)
	return validHash.MatchString(text)
}

func (h *TextHandler) HandleMain(m *tb.Message, u *user.User) {
	ib := [][]tb.InlineButton{
		{
			{
				Unique: "123",
				Text:   "Some Text",
			},
			{
				Unique: "321",
				Text:   "Some Text #2",
			},
		},
	}

	if _, err := h.bot.Send(m.Sender, "Уиииииииииииииииииии", &tb.ReplyMarkup{
		InlineKeyboard: ib,
	}); err != nil {
		log.Fatalln(err)
	}
}
