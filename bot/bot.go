package bot

import (
	"../card"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func (b *Bot) Display() []*card.Card {
	if b.IsDealer {
		// Dealer does not show his first card
		return b.Cards[1:]
	}
	return b.Cards
}

func (b *Bot) Hit(deck *card.Deck) {
	// deal a card from the deck and add to this bot
	card := deck.DealCard()
	b.Cards = append(b.Cards, card)
	if card.Value == 1 && b.Score < 10 {
		// Ace can be 11 or 1 depending on the context. We'll choose greedy and set it to bot's advantage.
		b.Score += 11
	} else {
		// All cards max out at 10
		b.Score += uint8(math.Min(float64(card.Value), 10.0))
	}
}

func (b *Bot) Stand() {
	// Literally does nothing. One could add a penalty-ish for this, to "rig" the game?
}

func (b *Bot) Init(wipeClean bool) {
	// Init ONLY initializes for a game, does not overwrite behavior, unless wipeClean is set.
	b.Score = 0
	b.IsDealer = false
	b.Cards = nil
	if wipeClean {
		b.ControlStruct.Init(b.ControlStruct.Alpha, b.ControlStruct.Gamma, b.ControlStruct.RandomProb, b.ControlStruct.TempDelta)
	}
}

func (b *Bot) ChooseBestAction(deck *card.Deck) int {
	// Choose the best action to make for the current score the bot is in.
	bestAction := 1
	if b.ControlStruct.Rewards[b.Score][0] > b.ControlStruct.Rewards[b.Score][1] {
		bestAction = 0
	}
	switch bestAction {
	case 0:
		b.Hit(deck)
	case 1:
		b.Stand()
	}
	return bestAction
}

func (b *Bot) ChooseRandomAction(deck *card.Deck) int {
	// Choose a random action for the bot.
	action := rand.Int() % 2
	switch action {
	case 0:
		b.Hit(deck)
	case 1:
		b.Stand()
	}
	b.LastAction = action
	return action
}

func (b *Bot) PerformAction(deck *card.Deck) int {
	// Store the old score to update the Q Values
	b.OldScore = b.Score
	var actionPerformed int
	// Seed everywhere. Randomize whenever you can.
	rand.Seed(time.Now().UnixNano())
	// A random value that determines whether you go brashly or carefully.
	value := rand.Float64()
	if value < b.ControlStruct.RandomProb {
		actionPerformed = b.ChooseRandomAction(deck)
	} else {
		actionPerformed = b.ChooseBestAction(deck)
	}
	// Q-Learning!
	reward := 0
	if actionPerformed == 0 {
		reward = 10
	}
	future := b.ControlStruct.Alpha * (float64(reward) + b.GetBestReward(b.Score)*b.ControlStruct.Gamma)
	b.ControlStruct.Rewards[b.OldScore][actionPerformed] = (1-b.ControlStruct.Alpha)*(b.ControlStruct.Rewards[b.OldScore][actionPerformed]) + future
	// Wisdom with experience
	b.ControlStruct.RandomProb *= b.ControlStruct.TempDelta
	return actionPerformed
}

func (b *Bot) GetBestReward(score uint8) float64 {
	if score > 21 {
		// You're done for, man
		return -100000
	}
	return math.Max(b.ControlStruct.Rewards[score][0], b.ControlStruct.Rewards[score][1])
}

func (b *Bot) PrintRewards() {
	fmt.Println(b.ControlStruct.Rewards)
	fmt.Println(b.ControlStruct.RandomProb)
}

func (b *Bot) UpdateVictory(didWin bool) {
	if didWin {
		b.ControlStruct.Rewards[b.OldScore][b.LastAction] = (1-b.ControlStruct.Alpha)*(b.ControlStruct.Rewards[b.OldScore][b.LastAction]) + 100000
	} else {
		b.ControlStruct.Rewards[b.OldScore][b.LastAction] = (1-b.ControlStruct.Alpha)*(b.ControlStruct.Rewards[b.OldScore][b.LastAction]) - 100000
	}
}
