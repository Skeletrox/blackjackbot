package game

import (
	"../bot"
	"../card"
)

func (g *Game) Init(Dealer, Other *bot.Bot) {
	g.Dealer = Dealer
	g.Other = Other
	g.Cards = &card.Deck{}
	g.Cards.Init()
}

func (g *Game) Run() uint8 {
	g.Dealer.Hit(g.Cards)
	g.Other.Hit(g.Cards)
	g.Dealer.Hit(g.Cards)
	g.Other.Hit(g.Cards)
	if g.Other.Score == 21 {
		return 0
	}
	if g.Dealer.Score == 21 {
		return 1
	}
	for g.Dealer.Score < 22 && g.Other.Score < 22 {
		g.Other.PerformAction(g.Cards)
		g.Dealer.Hit(g.Cards)
		if g.Other.Score == 21 && g.Dealer.Score == 21 {
			// draw
			return 2
		} else if g.Dealer.Score == 21 {
			// Dealer wins
			return 1
		} else if g.Other.Score == 21 {
			// Player wins
			return 0
		}
	}
	if g.Dealer.Score > 22 {
		return 0
	}
	return 1
}
