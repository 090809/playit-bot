package handlers

import (
	"fmt"
	"log"

	tb "github.com/090809/telebot"
	"playit-bot/buttons"
	"playit-bot/user"
	"playit-bot/utils"
)

var testStartMessage = "Время по гадальному хрусталю и простенькому тесту определить, " +
	"к какой банде ты относишься.\n" +
	"Первый вопрос:\n"

func (h *TextHandler) HandleTestStart(m *tb.Message) {
	u := h.repository.Find(m.Sender.Recipient())
	if u == nil {
		if _, err := h.bot.Send(m.Sender, "Пожалуйста, начните работу с ботом командой /start"); err != nil {
			log.Printf("[ERROR] %v", err)
		}
		return
	}

	if _, ok := u.Completed[user.Test]; ok {
		if _, err := h.bot.Send(m.Sender, "Вы уже выполняли это задание!"); err != nil {
			log.Printf("[ERROR] %v", err)
		}
		return
	}

	_, err := h.bot.Send(m.Sender, testStartMessage)
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}

	u.Phase = user.TestPhase
	h.testUser(m, u)
	h.repository.Save(u)
}

type question struct {
	Question string
	Answers  []answer
}

type answer struct {
	Text  string
	Value []user.AnswerValue
}

var testMap = map[int]question{
	1: {Question: "1. Свое утро я начинаю с:", Answers: []answer{
		{"Кофе", []user.AnswerValue{user.SMM}},
		{"Энергетиков", []user.AnswerValue{user.Back}},
		{"Полноценного завтрака", []user.AnswerValue{user.Mobile}},
		{"Воды", []user.AnswerValue{user.Design}},
		{"Ничего не ем", []user.AnswerValue{user.Manager}},
		{"Кола и чипсы", []user.AnswerValue{user.Front}},
	}},
	2: {Question: "2. Я не могу работать без:", Answers: []answer{
		{"Телефона", []user.AnswerValue{user.Mobile}},
		{"Быстрого интернета", []user.AnswerValue{user.SMM}},
		{"Планера", []user.AnswerValue{user.Manager}},
		{"Moose", []user.AnswerValue{user.Design}},
		{"Сетки", []user.AnswerValue{user.Front}},
		{"Тостера", []user.AnswerValue{user.Back}},
	}},
	3: {Question: "3. Я спокойно могу работать без:", Answers: []answer{
		{"Второго монитора", []user.AnswerValue{user.SMM}},
		{"Мышки", []user.AnswerValue{user.Manager, user.Front}},
		{"Макета", []user.AnswerValue{user.Back, user.Design}},
		{"Правок", []user.AnswerValue{user.Mobile}},
	}},
	4: {Question: "4. Мое тотемное животное:", Answers: []answer{
		{"Ягуар-оборотень", []user.AnswerValue{user.SMM}},
		{"Хамелеон", []user.AnswerValue{user.Design}},
		{"Клещ", []user.AnswerValue{user.Manager}},
		{"Слон", []user.AnswerValue{user.Back}},
		{"Петушок", []user.AnswerValue{user.Mobile}},
		{"Котопёс", []user.AnswerValue{user.Front}},
	}},
	5: {Question: "5. Что ты возьмёшь с собой на свидание?", Answers: []answer{
		{"Ноутбук", []user.AnswerValue{user.SMM, user.Design, user.Manager, user.Back, user.Mobile, user.Front}},
	}},
	6: {Question: "6. Моя любимая одежда:", Answers: []answer{
		{"Тёплый свитер", []user.AnswerValue{user.Back}},
		{"Деловой костюм", []user.AnswerValue{user.Manager}},
		{"Модные кросы", []user.AnswerValue{user.Design}},
		{"Ремень", []user.AnswerValue{user.Front}},
		{"Носки с сандалями", []user.AnswerValue{user.Mobile}},
	}},
	7: {Question: "7. Мой любимый фильм:", Answers: []answer{
		{"Матрица", []user.AnswerValue{user.Back}},
		{"Терминатор", []user.AnswerValue{user.Mobile}},
		{"Мстители. Финал", []user.AnswerValue{user.Front}},
		{"Престиж", []user.AnswerValue{user.Manager}},
		{"Ешь. Молись. Люби.", []user.AnswerValue{user.SMM}},
		{"Автостопом по галактике", []user.AnswerValue{user.Design}},
	}},
	8: {Question: "8. Мой любимый цвет:", Answers: []answer{
		{"Красный", []user.AnswerValue{user.Back}},
		{"Оранжевый", []user.AnswerValue{user.Front}},
		{"Синий", []user.AnswerValue{user.Mobile}},
		{"Темно-зеленый", []user.AnswerValue{user.Design}},
		{"Светло-зеленый", []user.AnswerValue{user.SMM}},
	}},
	9: {Question: "9. Мой талисман:", Answers: []answer{
		{"Бубен", []user.AnswerValue{user.Back}},
		{"Чехол на телефон", []user.AnswerValue{user.SMM}},
		{"Ручка", []user.AnswerValue{user.Manager}},
		{"Робот", []user.AnswerValue{user.Mobile}},
		{"Икона", []user.AnswerValue{user.Design}},
		{"Зеркало", []user.AnswerValue{user.Front}},
	}},
	10: {Question: "10. В душе я:", Answers: []answer{
		{"Фронтер", []user.AnswerValue{user.Front}},
		{"Бэкер", []user.AnswerValue{user.Back}},
		{"Дизайнер", []user.AnswerValue{user.Design}},
		{"Манагер", []user.AnswerValue{user.Manager}},
		{"Android-разработчик", []user.AnswerValue{user.Mobile}},
		{"СММ-щик", []user.AnswerValue{user.SMM}},
	}},
}

func (h *TextHandler) HandleTest(m *tb.Message) {
	u := h.repository.Find(m.Sender.Recipient())
	if u == nil {
		return
	}

	vMap := testMap[u.TestPhase].Answers

	for _, val := range vMap {
		if val.Text == m.Text {
			h.processTest(m, u, val.Value)
			return
		}
	}

	_, err := h.bot.Send(m.Sender, "Используй клавиатуру снизу для ответов")
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
}

func (h *TextHandler) processTest(m *tb.Message, u *user.User, values []user.AnswerValue) {
	for _, val := range values {
		u.Answers[val]++
	}

	if u.TestPhase < 10 {
		h.testUser(m, u)
		return
	}

	h.testResult(m, u)
}

func (h *TextHandler) testUser(m *tb.Message, u *user.User) {
	u.TestPhase++

	var rk [][]tb.ReplyButton
	var rkb []tb.ReplyButton

	for i, val := range testMap[u.TestPhase].Answers {
		if i%2 == 0 {
			rk = append(rk, rkb)
			rkb = []tb.ReplyButton{}
		}
		b := tb.ReplyButton{Text: val.Text}
		rkb = append(rkb, b)
	}
	rk = append(rk, rkb)

	_, err := h.bot.Send(m.Sender, testMap[u.TestPhase].Question, &tb.ReplyMarkup{
		ReplyKeyboard: rk,
	})
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}

	h.repository.Save(u)
}

func (h *TextHandler) testResult(m *tb.Message, u *user.User) {
	var text string
	switch getMaxValue(u.Answers) {
	case user.Design:
		text = "ты повелитель цвета и формы, можешь сочетать не сочетаемое и не сочетать сочетаемое.\n\nТы - дизайнер"
	case user.Mobile:
		text = "телефончики твоя жизнь, ты готов сделать для них все, даже приложение.\n\nТы - мобильный разработчик"
	case user.Front:
		text = "ты фронтер. Посмотри в интернете."
	case user.Manager:
		text = "ты любишь заставлять людей работать, а твой уровень харизмы прокачен на 100% еще с детсва.\n\nТы - менеджер"
	case user.SMM:
		text = "ты просто гуру соцсетей, все инстагёрлы завидуют количеству твоих подписчиков, а сам Толстой не написал бы постов лучше, чем ты.\n\nТы - СММ-щик"
	case user.Back:
		text = "ты суровый программист, вооруженный бубном, а вместо крови у тебя по венам течет энергетик.\n\nТы - бэкендер"
	}

	_, err := h.bot.Send(m.Sender, fmt.Sprintf("Поздравлям, %s\n\nЗа это задание ты получил 150 devcoin-ов!", text), &tb.ReplyMarkup{
		ReplyKeyboard: buttons.MainReplyButtons,
	})
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}

	u.Completed[user.Test] = true
	if err := utils.Confirm("test", *u.HashTag); err != nil {
		log.Printf("[ERROR] %v", err)
	}
	h.repository.Save(u)
}

func getMaxValue(m map[user.AnswerValue]int) user.AnswerValue {
	var mKey = user.SMM
	var max = 0
	for key, val := range m {
		if val > max {
			max = val
			mKey = key
		}
	}
	return mKey
}
