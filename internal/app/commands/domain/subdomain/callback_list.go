package subdomain

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ozonmp/omp-bot/internal/app/path"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *commander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("DemoSubdomainCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)

		return
	}

	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		fmt.Sprintf("Parsed: %+v\n", parsedData),
	)

	if _, err = c.bot.Send(msg); err != nil {
		log.Printf("DemoSubdomainCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
