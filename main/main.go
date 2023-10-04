package main

import (
	"fmt"

	"example.com/card"
)

func main() {
	stack := card.Init0()
	card.Shuffle(stack)
	var player, dealer []card.Card
	for i := 0; i < 2; i++ {
		stack = card.Pick(stack, &player)
		stack = card.Pick(stack, &dealer)
	}
	for {
		var flag bool
		fmt.Println("player: ", player)
		fmt.Println("dealer: ", dealer)
		flag, stack = card.ReadAction(stack, &player)
		num := card.Calc()
		if !flag {
			break
		}
	}
}
