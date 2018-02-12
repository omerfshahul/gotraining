package main

func main() {
	//cards := newDeck()

	//	cards.print()

	//hand, remainingCards := deal(cards, 5)

	//	hand.print()
	//	remainingCards.print()

	//	s := cards.toString()
	//	fmt.Println(s)
	//d := toDeck(s)
	//d.print()
	//d.saveToFile("deckfile.txt")

	// tmpDeck, err := newDeckFromFile("deckfile1.txt")

	// if err == nil {
	// 	tmpDeck.print()
	// } else {
	// 	fmt.Println("Error: " + err.Error())
	// 	os.Exit(1)
	// }
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
