package card

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"example.com/lib"
)

type Suit string

const (
	Spade   = Suit("Spade")
	Club    = Suit("Club")
	Diamond = Suit("Diamond")
	Heart   = Suit("Heart")
)

type Face string

const (
	King  = Face("King")
	Queen = Face("Queen")
	Jack  = Face("Jack")
	Ten   = Face("Ten")
	Nine  = Face("Nine")
	Eight = Face("Eight")
	Seven = Face("Seven")
	Six   = Face("Six")
	Five  = Face("Five")
	Four  = Face("Four")
	Three = Face("Three")
	Two   = Face("Two")
	Ace   = Face("Ace")
)
const BurstThreshold = 22

type Facenum struct {
	face Face
	// for Ace
	nums []int
}

type Card struct {
	suit    Suit
	facenum Facenum
}

func Init0() []Card {
	faces := []Face{
		Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King,
	}
	nums := [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 11, 11}

	var facenum []Facenum
	for i := 0; i < 13; i++ {
		facenum = append(facenum, Facenum{face: faces[i], nums: []int{nums[i]}})
	}
	facenum[0].nums = append(facenum[0].nums, 11)

	suits := []Suit{Spade, Club, Diamond, Heart}
	var stack []Card
	for i := 0; i < 13; i++ {
		for j := 0; j < 4; j++ {
			stack = append(stack, Card{suit: suits[j], facenum: facenum[i]})
		}
	}
	return stack
}

func Shuffle(stack []Card) {
	for i := range stack {
		j := rand.Intn(i + 1)
		stack[i], stack[j] = stack[j], stack[i]
	}
}

func Pick(stack []Card, hands *[]Card) []Card {
	len := len(stack)
	*hands = append(*hands, stack[0])
	return stack[1:len]
}

func ReadAction(stack []Card, hands *[]Card) (bool, []Card) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("pick up card? [y/n]: ")
	scanner.Scan()
	txt := scanner.Text()
	flag := true
	switch txt {
	case "y":
		stack = Pick(stack, hands)
	case "n":
		flag = false
	case "":
	default:
		fmt.Println("type y or n!")
	}
	return flag, stack
}

func Calc(hands []Card) int {
	sums := []int{0}
	for _, card := range hands {
		var lst []int
		for s := range sums {
			for _, num := range card.facenum.nums {
				lst = append(lst, s+num)
			}
		}
		sums = lst
	}
	res := -1
	for _, num := range sums {
		if num >= BurstThreshold {
			continue
		} else {
			res = lib.Max(res, num)
		}
	}
	if res == -1 {
		res = BurstThreshold
	}
	return res
}
