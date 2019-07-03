package main

import (
	"context"
	"testing"

	tb "github.com/090809/telebot"
	"playit-bot/handlers"
	"playit-bot/user"
)

func Test_StartHandler_ProcessUser(t *testing.T) {
	ctx := context.Background()
	ur := user.NewRepository(nil, nil)

	defHandler := handlers.NewHandler(nil, ctx, ur)
	startHandler := handlers.StartHandler{Handler: *defHandler}

	startHandler.ProcessUser("123")
}

func Test_TextHandler_HandleLogging(t *testing.T)  {
	ctx := context.Background()
	ur := user.NewRepository(nil, nil)

	defHandler := handlers.NewHandler(nil, ctx, ur)
	textHandler := handlers.TextHandler{Handler: *defHandler}

	var emptyTestMessage = tb.Message{
		ID: 123,
		Sender: &tb.User{
			ID: 123,
		},
	}

	if err := textHandler.Handle(&emptyTestMessage); err == nil {
		t.Error("Не должно было пройти валидацию")
	}

	var wrongTestMessage = tb.Message{
		ID: 123,
		Sender: &tb.User{
			ID: 123,
		},
		Text: "wrongtext",
	}

	if err := textHandler.Handle(&wrongTestMessage); err == nil {
		t.Error("Не должно было пройти валидацию")
	}

	var rightTestMessage = tb.Message{
		ID: 123,
		Sender: &tb.User{
			ID: 123,
		},
		Text: "#123456",
	}

	if err := textHandler.Handle(&rightTestMessage); err != nil {
		t.Error(err)
	}
}