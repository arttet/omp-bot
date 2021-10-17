package subdomain

import (
	"log"

	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/domain/subdomain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type SubdomainCommander interface {
	Help(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig
	Get(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig
	List(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig
	Delete(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig

	New(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig
	Edit(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig

	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(callback *tgbotapi.Message, commandPath path.CommandPath)
}

type commander struct {
	bot     *tgbotapi.BotAPI
	service service.SubdomainService
}

func NewSubdomainCommander(
	bot *tgbotapi.BotAPI,
	service service.SubdomainService,
) SubdomainCommander {

	return &commander{
		bot:     bot,
		service: service,
	}
}

func (c *commander) HandleCallback(
	callback *tgbotapi.CallbackQuery,
	callbackPath path.CallbackPath,
) {

	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *commander) HandleCommand(
	msg *tgbotapi.Message,
	commandPath path.CommandPath,
) {
	var replay tgbotapi.MessageConfig

	switch commandPath.CommandName {
	case "help":
		replay = c.Help(msg)
	case "list":
		replay = c.List(msg)
	case "get":
		replay = c.Get(msg)
	case "delete":
		replay = c.Delete(msg)
	case "new":
		replay = c.New(msg)
	case "edit":
		replay = c.Edit(msg)
	default:
		replay = c.Default(msg)
	}

	if _, err := c.bot.Send(replay); err != nil {
		log.Printf("Domain.Subdomain.Commander: error sending reply message to chat - %v", err)
	}
}
