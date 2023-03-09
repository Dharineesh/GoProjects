package main

import "fmt"

func main() {
	card := newDeck()
	// handCard, remainingCard := deal(card, 3)

	// fmt.Println(handCard.toString())
	// fmt.Println(remainingCard.toString())

	// card.saveFileToDrive("new_deck")
	// var newCard deck = deckFromDrive("new_deck")
	// fmt.Println(newCard)
	// sampleString := "Hi there!"
	fmt.Println("Before")
	card.printValue()
	card.shuffle()
	fmt.Println("After")
	card.printValue()

	// fmt.Println([]byte(sampleString))
}
