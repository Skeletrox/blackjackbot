package game

import (
	"../bot"
	"../card"
)

type Game struct {
	Dealer, Other *bot.Bot
	Cards         *card.Deck
}
