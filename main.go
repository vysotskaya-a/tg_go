package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {

	botToken := "6576344046:AAGNgp6aHaHFfXNAcwpSbH8xZTk8seR0N_k"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Обработчик команды старт
		chatID := tu.ID(update.Message.Chat.ID)

		keyboard := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("обузьяны мемы"),
				tu.KeyboardButton("дети мемы"),
			),
			tu.KeyboardRow(
				tu.KeyboardButton("миньоны"),
				tu.KeyboardButton("армянская бабушка"),
			),
		)

		message := tu.Message(
			chatID,
			"Привет! Это бот со стикерами! Нажми нужное описание на клавиатуре)",
		).WithReplyMarkup(keyboard)

		_, _ = bot.SendSticker(
			tu.Sticker(
				chatID,
				tu.FileFromID("CAACAgIAAxkBAAEMszZmyobzVK14pqylq1-m7IJu2wVXhQACeTkAAgLC4EpPQVSMKtXtCzUE"),
			),
		)

		_, _ = bot.SendMessage(message)

	}, th.CommandEqual("start"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {

	}, th.AnyCommand())

	for update := range updates {
		if update.Message != nil {
			chatID := tu.ID(update.Message.Chat.ID)

			_, _ = bot.CopyMessage(
				tu.CopyMessage(
					chatID,
					chatID,
					update.Message.MessageID,
				),
			)

			_, _ = bot.SendSticker(
				tu.Sticker(
					chatID,
					tu.FileFromID("CAACAgIAAxkBAAEMszZmyobzVK14pqylq1-m7IJu2wVXhQACeTkAAgLC4EpPQVSMKtXtCzUE"),
				),
			)
		}
	}

	bh.Start()

}
