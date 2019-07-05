package handlers

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	tb "github.com/090809/telebot"
	"playit-bot/buttons"
	"playit-bot/user"
	"playit-bot/utils"
)

var errorMessage = "Бип-бип! У нас произошла ошибка, подойди к ребятам на стенде, и покажи им это сообщение."

var loggedText = "Мы проверили твой код, тебя же зовут %s? Если нет, начни сначала и введи правильный код"

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
				log.Printf("[ERROR] %v", err)
			}
		}
		return
	}

	switch u.Phase {
	case user.LoggingPhase:
		h.HandleLogging(m, u)
	case user.MainPhase:
		h.HandleMain(m, u)
	case user.ProjectPhase:
		h.HandleProjectText(m)
	case user.TestPhase:
		h.HandleTest(m)
	default:
		if h.bot != nil {
			if _, err := h.bot.Send(m.Sender, "Я не распознал вашу команду"); err != nil {
				log.Printf("[ERROR] %v", err)
			}
		}
	}
}

func (h *TextHandler) HandleLogging(m *tb.Message, u *user.User) {
	tag := m.Text
	if !h.validateLoggingText(tag) {
		if h.bot != nil {
			if _, err := h.bot.Send(m.Sender, "Необходимо ввести правильный код. Пример: \"#000000\""); err != nil {
				log.Printf("[ERROR] %v", err)
			}
		}
		return
	}

	tagId := tag[1:]

	name, err := utils.CheckTag(tagId)
	if err != nil {
		if _, err := h.bot.Send(m.Sender, errorMessage); err != nil {
			log.Printf("[ERROR] %v", err)
		}
		return
	}

	u.HashTag = &tagId
	u.Phase = user.MainPhase
	h.repository.Save(u)

	ib := buttons.LoginInlineButtons

	if _, err := h.bot.Send(m.Sender, fmt.Sprintf(loggedText, name), &tb.ReplyMarkup{
		InlineKeyboard: ib,
	}); err != nil {
		log.Printf("[ERROR] %v", err)
	}
}

func (h *TextHandler) validateLoggingText(text string) bool {
	var validHash = regexp.MustCompile(`[#][0-9]{5,6}`)
	return validHash.MatchString(text)
}

func (h *TextHandler) HandleMain(m *tb.Message, u *user.User) {
	replyButtons := buttons.MainReplyButtons

	if _, err := h.bot.Send(m.Sender, "Команда не разобрана", &tb.ReplyMarkup{
		ReplyKeyboard: replyButtons,
	}); err != nil {
		log.Printf("[ERROR] %v", err)
	}
}
