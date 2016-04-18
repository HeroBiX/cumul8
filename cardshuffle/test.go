package main

import (
	"fmt"
)

type Card struct {
	Number int `json:"number"`
}

var cards = []Card{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}, {11}, {12}}

func main() {
	var deck1 []Card
	var deck2 []Card

	fmt.Println("Cards :", cards)
	fmt.Println("Cards Len/2 :", len(cards)/2)
	deck1 = append(deck1, cards[:len(cards)/2]...)
	fmt.Println("deck1: ", deck1)
	deck2 = append(deck2, cards[len(cards)/2:]...)
	fmt.Println("deck2: ", deck2)
}
