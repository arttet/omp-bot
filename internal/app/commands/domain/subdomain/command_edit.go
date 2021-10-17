package subdomain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *commander) Edit(
	inputMessage *tgbotapi.Message,
) tgbotapi.MessageConfig {

	return tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Not Implemented",
	)
}
