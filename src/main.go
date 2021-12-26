package main

func main() {
	cards := deck{newCard(), "Ace of Spades"}
	cards = append(cards, "Ace of Diamonds")

	cards.print()

}

func newCard() string { //Specify data type
	return "Ace of Diamonds"
}
