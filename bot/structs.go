package bot

import "../card"

type (
	/*
		A Bot is a player.
	*/
	Bot struct {
		// The set of cards the bot has
		Cards []*card.Card
		// The score of the bot
		Score, OldScore uint8
		// Is the bot a dealer?
		IsDealer bool
		// The reinforcement values in a "bot-controlling" structure
		ControlStruct *Reinforcement
		// The last action.
		LastAction int
		// has this bot doubled its stakes?
		HasDoubled bool
	}
	/*
		A Reinforcement is a reinforcement structure.
	*/
	Reinforcement struct {
		// Learning Rate
		Alpha float64
		// Decay Rate for Reinforcement Learning
		Gamma float64
		// Probability of choosing a random value.
		RandomProb float64
		/*Possible moves are: Hit, Stay (we ignore double for now). Possible Scores range between 1 and 21*/
		Rewards [][]float64
		// Cool down factor for random action probability
		TempDelta float64
	}
)
