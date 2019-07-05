package main

import (
	"context"
	"log"
	"os"

	//"fmt"

	tb "github.com/090809/telebot"
	"github.com/joho/godotenv"
	"playit-bot/handlers"
	"playit-bot/user"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		if _, err := os.Stat("./.env.example"); os.IsExist(err) {
			log.Fatalln("File .env.example exists! Fill it and rename to .env, before we go")
		}
		log.Fatalf("Error .env loading: %v", err.Error())
	} else {
		log.Println(".env file loaded")
	}

	var (
		listenURL = os.Getenv("URL")
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
		token     = os.Getenv("TOKEN")      // you must add it to your config vars
	)

	webhook := &tb.Webhook{
		Listen:   listenURL + ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	settings := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(settings)
	if err != nil {
		log.Fatalln(err)
	}

	ur := user.NewRepository(nil, nil)

	defHandler := handlers.NewHandler(b, ctx, ur)
	startHandler := handlers.StartHandler{Handler: *defHandler}
	textHandler := handlers.TextHandler{Handler: *defHandler}
	callbackHandler := handlers.CallbackHandler{Handler: *defHandler}
	projectHandler := handlers.ProjectHandler{Handler: *defHandler}

	b.Handle(tb.OnText, textHandler.Handle)
	b.Handle("/start", startHandler.Handle)
	b.Handle(tb.OnCallback, callbackHandler.Handle)
	b.Handle("🐾 Это что за покемон? 🐾", textHandler.HandleTestStart)
	b.Handle("💰Проект на миллион💰", projectHandler.Handle)

	b.Start()
}
