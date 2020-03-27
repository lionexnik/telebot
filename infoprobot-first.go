


package main

import (
    "fmt"
    "gopkg.in/telegram-bot-api.v4"
    "log"
    "os"
    "encoding/json"
)

type Config struct {
    TelegramBotToken string
}

func main() {



file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}


	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)

if err != nil {
    log.Panic(err)
}

bot.Debug = false

log.Printf("Authorized on account %s", bot.Self.UserName)

u := tgbotapi.NewUpdate(0)
u.Timeout = 60

updates, err := bot.GetUpdatesChan(u)

if err != nil {
    log.Panic(err)
}
// В канал updates будут приходить все новые сообщения.
for update := range updates { 
    // Создав структуру - можно её отправить обратно боту
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
    msg.ReplyToMessageID = update.Message.MessageID
    bot.Send(msg)
}
}
