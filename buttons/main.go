package buttons

import (
	tb "github.com/090809/telebot"
)

var MainReplyButtons = [][]tb.ReplyButton{
	{
		{
			Text: "🐾 Это что за покемон? 🐾",
		},
	},
	{
		{
			Text: "💰Проект на миллион💰",
		},
	},
}

var LoginInlineButtons = [][]tb.InlineButton{
	{
		{
			Unique: "login-restart",
			Text:   "Ввести код заного 🔙",
		},
		{
			Unique: "login-ok",
			Text:   "Подвердить ✅",
		},
	},
}
