package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
	middleware "gopkg.in/telebot.v3/middleware"
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TELE_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Use(middleware.Logger())
	b.Use(middleware.AutoRespond())
	
	b.Handle("/cancel", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
