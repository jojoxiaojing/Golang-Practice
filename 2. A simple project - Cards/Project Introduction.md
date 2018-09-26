* newDeck():         create a list of playing cards. an array of strings
* print():           log out the contents of a deck of cards
* shuffle():         shuffles all the cards in a deck
* deal():            create a hand of cards
* saveToFile():      save a list of cards to a file on the local machine
* newDeckFromFile(): load a list of cards from the local machine

### How to declaire a variable
    var card string = "Ace of Spades"
    //declare a variable with a type

    card := "Ace of Spades" 
    //don't have to use type, only use ":" for the first initialization

    card = "Five of diamonds"
    //for the second time, don't use :. Otherwise it will go wrong

## How to declaire a function
    func newCard() string {
	return "Five of Diamonds"
}

## Basic type
* bool    true false
* string  "Five of Diamonds"
* int     1
* float64 10.0009

## Array and Slice
* Array is fixed size
    var a [2]string
	a[0] = "Hello"
	a[1] = "World"
    fmt.Println(a[0], a[1])
	fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
    
* Slice is not fixed
    cards := []string{"Ace of Diamonds", newCard()}
    cards = append(cards, "Six of Spades")

## For loop
    for i, card := range cards {
		fmt.Println(i, card)
	}