package basket

import (
	"encoding/json"
	"log"

	"github.com/ozonmp/omp-bot/internal/app/path"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *commander) List(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	outputMsgText := "Here all the products: \n\n"

	var (
		cursor uint64
		limit  uint64
	)

	// products
	_, err := c.service.List(cursor, limit)
	if err != nil {
		log.Printf("failed to list the products: %v", err)
		return tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, err := json.Marshal(CallbackListData{
		Offset: 21,
	})
	if err != nil {
		log.Printf("failed to list the products: %v", err)
		return tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
	}

	callbackPath := path.CallbackPath{
		Domain:       "buy",
		Subdomain:    "subdomain",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	return msg
}
