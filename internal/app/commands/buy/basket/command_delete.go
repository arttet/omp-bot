package basket

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *commander) Delete(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseInt(args, 10, 0)
	if err != nil {
		log.Println("wrong args", args)
		return tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			err.Error(),
		)
	}

	removed, err := c.service.Remove(uint64(idx))
	if err != nil {
		log.Printf("failed to remove the product with idx %d: %v", idx, err)
	}

	return tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("%v", removed),
	)
}
