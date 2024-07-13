// An executable package
package main

// Notes:

/*
*

  - Run the script

  - Open terminal in the root directory

  - go run main.go deck.go

    Go is not a POO language, so there's no classes concept inside it

    Slices: Every element in a slice must
    be of the same type.

    cards := []string

*
*/
func main() {
	// Same ways of defining a string variable
	//var card string = "Ace of Spades"

	// Only for new variables
	//card := newCard()

	cards := newDeck()

	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
