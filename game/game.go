package game

import (
	"../bot"
	"../card"
	"math/rand"
	"time"
)

func (g *Game) Init(Dealer, Other *bot.Bot) {
	rand.Seed(time.Now().UnixNano())
	g.Dealer = Dealer
	g.Dealer.Init(false)
	g.Dealer.IsDealer = true
	g.Other = Other
	g.Other.Init(false)
	g.Other.IsDealer = false
	g.Cards = &card.Deck{}
	g.Cards.Init()
}

func (g *Game) RunGame() uint8 {
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
		var actionPerformed int
		if !g.Other.HasDoubled {
			actionPerformed = g.Other.PerformAction(g.Cards)
		} else {
			actionPerformed = 1
		}
		if g.Dealer.Score < 17 {
			g.Dealer.Hit(g.Cards)
		} else if actionPerformed == 1 {
			if g.Other.Score > g.Dealer.Score {
				return 0
			}
			return 1
		}
		if g.Other.Score == 21 && g.Dealer.Score == 21 {
			// draw
			return 2
		} else if g.Other.Score == 21 {
			// Other wins
			return 0
		} else if g.Dealer.Score == 21 {
			// Dealer wins
			return 1
		}
	}
	if g.Other.Score > 22 {
		return 1
	}
	return 0
}

func (g *Game) Run() uint8 {
	res := g.RunGame()
	switch res {
	case 0:
		g.Dealer.UpdateVictory(false)
		g.Other.UpdateVictory(true)
	case 1:
		g.Dealer.UpdateVictory(true)
		g.Other.UpdateVictory(false)
	case 2:
		g.Dealer.UpdateVictory(false)
		g.Other.UpdateVictory(false)
	}
	return res
}
