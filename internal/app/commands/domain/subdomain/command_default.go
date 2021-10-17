package subdomain

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *commander) Default(
	inputMessage *tgbotapi.Message,
) tgbotapi.MessageConfig {

	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	return tgbotapi.NewMessage(inputMessage.Chat.ID, "I do not understand: "+inputMessage.Text)
}
