package main

func main() {
	cards := newDeck()
    cards.shuffle()
    cards.print()
    /*
    cards := getDeckFromFile("my_cards")
	cards := newDeck()

    cards.saveToFile("my_cards")
    //cards.print()       // Defined in deck.go
    hands, remainingCards := deal(cards, 5)

    fmt.Println("hands.........")
    hands.print()
    fmt.Println("remaining.........")
    remainingCards.print()
    */
}
