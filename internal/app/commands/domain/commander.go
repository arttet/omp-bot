package domain

import (
	"log"

	"github.com/ozonmp/omp-bot/internal/app/commands/domain/subdomain"
	command "github.com/ozonmp/omp-bot/internal/app/commands/domain/subdomain"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/domain/subdomain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type DomainCommander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type commander struct {
	bot                *tgbotapi.BotAPI
	subdomainCommander command.SubdomainCommander
}

func NewDomainCommander(
	bot *tgbotapi.BotAPI,
) DomainCommander {

	subdomainService := service.NewService()

	return &commander{
		bot:                bot,
		subdomainCommander: subdomain.NewSubdomainCommander(bot, subdomainService),
	}
}

func (c *commander) HandleCallback(
	callback *tgbotapi.CallbackQuery,
	callbackPath path.CallbackPath,
) {

	switch callbackPath.Subdomain {
	case "subdomain":
		c.subdomainCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DemoCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *commander) HandleCommand(
	msg *tgbotapi.Message,
	commandPath path.CommandPath,
) {

	switch commandPath.Subdomain {
	case "subdomain":
		c.subdomainCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
