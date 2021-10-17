package subdomain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *commander) Help(
	inputMessage *tgbotapi.Message,
) tgbotapi.MessageConfig {

	const help = `/help__domain__subdomain - help
/get__domain__subdomain - get a entity
/list__domain__subdomain - get a list of your entity
/delete__domain__subdomain - delete an existing entity
/new__domain__subdomain - create a new entity
/edit__domain__subdomain - edit a entity
`
	return tgbotapi.NewMessage(inputMessage.Chat.ID, help)
}
