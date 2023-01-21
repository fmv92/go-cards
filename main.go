package main

func main() {
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	newDeck := newDeckFromFile("my_cards")
	newDeck.print()
	//fmt.Println(newDeck.toString())
}