package subdomain

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *commander) Get(
	inputMessage *tgbotapi.Message,
) tgbotapi.MessageConfig {

	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseInt(args, 10, 0)
	if err != nil {
		log.Println("wrong args", args)
	}

	// product
	_, err = c.service.Describe(uint64(idx))
	if err != nil {
		log.Printf("failed to describe a product with idx %d: %v", idx, err)
	}

	return tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"product.Title",
	)
}
