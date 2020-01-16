package main

import (
	"./bot"
	"./game"
	"fmt"
)

func main() {
	// Define some constants
	brashGamma := 0.8
	conservativeGamma := 0.99
	alpha := 0.8
	randomProb := 0.9
	brashCoolDown := 0.99
	cautiousCoolDown := 0.4

	cautionControl := &bot.Reinforcement{}
	cautionControl.Init(alpha, conservativeGamma, randomProb, cautiousCoolDown)
	brashControl := &bot.Reinforcement{}
	brashControl.Init(alpha, brashGamma, randomProb, brashCoolDown)

	cautionBot := &bot.Bot{}
	brashBot := &bot.Bot{}
	cautionBot.ControlStruct = cautionControl
	brashBot.ControlStruct = brashControl
	cautionBot.Init(true)
	brashBot.Init(true)
	dealerWins := 0
	cautionWins := 0
	brashWins := 0
	for i := 0; i < 100000; i++ {
		fmt.Println("Game number:", i+1)
		game := game.Game{}
		if i%2 == 0 {
			game.Init(brashBot, cautionBot)
		} else {
			game.Init(cautionBot, brashBot)
		}
		result := game.Run()
		if result == 1 {
			dealerWins++
		}
		if (int(result)+i)%2 == 0 {
			cautionWins++
		} else {
			brashWins++
		}
	}
	cautionBot.PrintRewards()
	brashBot.PrintRewards()
	fmt.Println("Number of games won by dealer in training phase:", dealerWins)
	fmt.Println("Games won by brashBot:", brashWins, " and cautionBot:", cautionWins)
	/*
		Some unbiased opponent acting as a dummy dealer.
	*/
	unBiased := &bot.Bot{}
	unBiasedControl := &bot.Reinforcement{}
	unBiased.ControlStruct = unBiasedControl

	game := game.Game{}

	brashWins = 0
	cautionWins = 0
	for i := 0; i < 10; i++ {
		unBiased.Init(true)
		game.Init(unBiased, brashBot)
		result := game.Run()
		if result != 1 {
			brashWins++
		}
		unBiased.Init(true)
		game.Init(unBiased, cautionBot)
		result = game.Run()
		if result != 1 {
			cautionWins++
		}
	}
	fmt.Println("Test result: Brashbot:", brashWins, "cautionBot:", cautionWins)
}
