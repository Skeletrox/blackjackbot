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
	for i := 0; i < 10000; i++ {
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
	}
	cautionBot.PrintRewards()
	brashBot.PrintRewards()
	fmt.Println("Number of games won by dealer:", dealerWins)
}
