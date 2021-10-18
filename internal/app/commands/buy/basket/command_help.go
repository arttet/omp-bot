package basket

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *commander) Help(
	inputMessage *tgbotapi.Message,
) tgbotapi.MessageConfig {

	const help = `/help__buy__basket - help
/get__buy__basket - get a entity
/list__buy__basket - get a list of your entity
/delete__buy__basket - delete an existing entity
/new__buy__basket - create a new entity
/edit__buy__basket - edit a entity
`
	return tgbotapi.NewMessage(inputMessage.Chat.ID, help)
}
