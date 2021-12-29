package main

func main() {
	// var card = "Ace of Spades"
	// card := "Ace of Spades"

	// 1
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// fmt.Println("------")
	// fmt.Println(cards.toString())

	// cards.saveToFile("/tmp/my_cards.txt")

	// 2
	// cards := newDeckFromFile("/tmp/my_cards.txt")
	// cards.print()

	// 3
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
