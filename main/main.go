package main

import (
	"fmt"

	"example.com/card"
)

func printStatus(player []card.Card, dealer []card.Card) {
	fmt.Println("player: ", player)
	fmt.Println("dealer: ", dealer)
}

func printResult(player []card.Card, dealer []card.Card, winStatus card.WinStatus) {
	println("\n[[Result]]")
	printStatus(player, dealer)
	println("win: ", winStatus)
}

func main() {
	stack := card.Init0()
	card.Shuffle(stack)
	var player, dealer []card.Card
	for i := 0; i < 2; i++ {
		stack = card.Pick(stack, &player)
		stack = card.Pick(stack, &dealer)
	}
	winStatus := card.Player
	// player
	for {
		var flag bool
		printStatus(player, dealer)
		flag, stack = card.ReadAction(stack, &player)
		num := card.Calc(player)
		if card.IsBurst(num) {
			winStatus = card.Dealer
			break
		}
		if !flag {
			break
		}
	}
	// dealer
	for {
		num := card.Calc(dealer)
		if card.IsBurst(num) {
			break
		}
		if card.IsDealerStopped(num) {
			winStatus = card.GetWinStatus(player, dealer)
			break
		}
		stack = card.Pick(stack, &dealer)
	}
	printResult(player, dealer, winStatus)
}
