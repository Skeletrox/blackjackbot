# Blackjack Bot

Bot that plays Blackjack with 2 approaches: Radical (constantly trying new things) and conservative (Preferring to stay true to roots). Actual behavior is modeled by pitting a radical agent against a conservative one for 100,000 games.

## Execution

`go run ./main.go` will run the main program and demonstrate how it works. As Blackjack is a *very* random game, with odds often against the player, results may vary wildly at the end, but you can see how the bot approximates rewards across multiple actions over multiple states.

## Are PRs welcome?

YES! I'd appreciate any input on how I could use my GPU to accelerate execution.