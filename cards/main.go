package main

func main() {
	//var card string = "Ace of Spades"
	//card := newCard()
	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of spades")

	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
