
package main

import (

	//  "gopkg.in/telegram-bot-api.v4"
	//  "log"
	//  "os"

	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "go/types"
	"gopkg.in/telegram-bot-api.v4"
	_ "io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	_ "os"
	"strings"
	//  "math/big"
	//   "net/url"
	"time"
)
type Config struct {
	TelegramBotToken string
}

func ByteSlice(b []byte) []byte { return b }
func GetStockQuotes(symbols string, dataFormat ...int) (quotes string, err error) {
	var buf []uint8
	var resp *http.Response
	var symbolsString string
	var urlString string


	//handle the stock symbols


		symbolsString = symbols


	//handle the data format, defaults to json
	urlString = fmt.Sprintf("https://od-api.oxforddictionaries.com/api/v2/entries/EN-US/%s",symbolsString)
	//пример:
	//https://api.tiingo.com/iex/?tickers=goog&token=f12dadf2c9b6f44e01699335d6b011f713e0d363&format=json
	client := &http.Client{Timeout: 30 * time.Second}
	request, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT x.y; Win64; x64; rv:10.0) Gecko/20100101 Firefox/10.0")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-rapidapi-host", "mashape-community-urban-dictionary.p.rapidapi.com")
	request.Header.Set("x-rapidapi-key", "1cb7afaf9amsh3f337ac58d6ef53p1e59e5jsn61b199c309a5")

	//fmt.Println(urlString)
	if resp, err = client.Do(request); err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if buf, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println(err)
		return
	}

	 arr := strings.SplitAfter(string(buf), ":")
	 //arr = strings.SplitAfter(arr[2], ".")

    quotes = arr[2]
	return
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




		quotes, err := GetStockQuotes(string(update.Message.Text))
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, quotes)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}



}
