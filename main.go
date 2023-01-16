package main

import (
	"log"
	"os"
	"time"

	cron "github.com/robfig/cron/v3"
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

	savedMsg := ""

	b.Use(middleware.Logger())
	b.Use(middleware.AutoRespond())

	cronjob := cron.New()
	cronjob.Start()

	b.Handle("/register", func(c tele.Context) error {
		if c.Text()[10:] != "" {
			savedMsg = c.Text()[10:]
			cronjob.AddFunc("@TZ=Asia/Bangkok 20 04 * * * *", func() {
				b.Send(c.Sender(), "Wanna /cancel ?")
			})
			cronjob.AddFunc("@TZ=Asia/Bangkok 30 04 * * * *", func() {
				b.Send(c.Sender(), savedMsg)
			})
			return c.Send("Registered!")
		} else {
			return c.Send("Please enter a msg")
		}
	})

	b.Handle("/cancel", func(c tele.Context) error {
		cronjob.Stop()
		c.Send("Canceled!")
		time.Sleep(1 * time.Hour)
		cronjob.Start()
		return nil
	})

	b.Start()
}
