package buy

import (
	"log"

	"github.com/ozonmp/omp-bot/internal/app/commands/buy/basket"
	"github.com/ozonmp/omp-bot/internal/app/path"
	basketservice "github.com/ozonmp/omp-bot/internal/service/buy/basket"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BuyCommander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type commander struct {
	bot             *tgbotapi.BotAPI
	basketCommander basket.BasketCommander
}

func NewBuyCommander(bot *tgbotapi.BotAPI) BuyCommander {
	basketService := basketservice.NewService()

	return &commander{
		bot:             bot,
		basketCommander: basket.NewBasketCommander(bot, basketService),
	}
}

func (c *commander) HandleCallback(
	callback *tgbotapi.CallbackQuery,
	callbackPath path.CallbackPath,
) {

	switch callbackPath.Subdomain {
	case "subdomain":
		c.basketCommander.HandleCallback(callback, callbackPath)
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
		c.basketCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
