package handlers

import (
	"log"

	tb "github.com/090809/telebot"
	"playit-bot/buttons"
	"playit-bot/user"
	"playit-bot/utils"
)

type ProjectHandler struct {
	Handler
}

var textProjectMain = "Расскажи о своем проекте на миллион (текст не менее 100 символов) " +
	"и получи за это N баллов!"
var goodProjectIdea = "Крутая идея, сейчас я передам ее нашим организаторам и " +
	"они обговорят ее с тобой, а пока мы тебе начислили N баллов"

var badProjectIdea = "Попробуй описать свою идею более подробно"

var errorUtils = "Во время обработки твоей идеи произошла ошибка, попробуй отправить ее чуть позже"

func (h *ProjectHandler) Handle(m *tb.Message) {
	u := h.repository.Find(m.Sender.Recipient())
	if u == nil {
		return
	}

	if _, ok := u.Completed[user.Project]; ok {
		if _, err := h.bot.Send(m.Sender, "Вы уже выполняли это задание!"); err != nil {
			log.Printf("[ERROR] %v", err)
		}
		return
	}
	_, err := h.bot.Send(m.Sender, textProjectMain, &tb.ReplyMarkup{
		ReplyKeyboardRemove: true,
	})
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}

	u.Phase = user.ProjectPhase
	h.repository.Save(u)
}

func (h *TextHandler) HandleProjectText(m *tb.Message) {
	if len(m.Text) < 100 {
		_, err := h.bot.Send(m.Sender, badProjectIdea)
		if err != nil {
			log.Printf("[ERROR] %v", err)
		}
		return
	}

	u := h.repository.Find(m.Sender.Recipient())
	if err := utils.Confirm("idea", *u.HashTag); err != nil {
		_, err := h.bot.Send(m.Sender, errorUtils)
		if err != nil {
			log.Printf("[ERROR] %v", err)
		}
		return
	}

	_, err := h.bot.Send(m.Sender, goodProjectIdea, &tb.ReplyMarkup{
		ReplyKeyboard: buttons.MainReplyButtons,
	})
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}

	u.Phase = user.MainPhase
	u.ProjectText = m.Text
	u.Completed[user.Project] = true
	h.repository.Save(u)
}
