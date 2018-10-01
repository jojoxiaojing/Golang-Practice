# Golang Practice

# Chapter One -- Hello World

### How do we run the code:

`go run main.go`
 run the main directly

`go build main.go`
make the file into main and run ./main

### What does 'package main' mean?

* Package == Project == Workspace
* Types of packages: 

(a) Executable: generates a file that we can run, must contain a function called "main" 

and (b)reusable: helpers, put reusable logic
* package main -> go build -> produce main

### What does import "fmt" means
to get another package

### what is 'func' means
short for function

### How is the main.go file organized
* Package declaration : 

        `package main`
* Import other packages that we need: 

        `import "fmt"`
* Declare functions, tell Go to do things:
```
        func main(){
            fmt.Println("hi there")
        }
 ```

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


	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+"of"+suit)
		}
	}

## Declare a new class
Any variable with a deck type can get access to print method now

    type deck []string

    func newDeck() deck {

	cards := deck{}
	return cards
}

    //d:    actual copy of the deck
    //deck: Every variable of type 'deck' can get access to print method now
    func (d deck) print() {
        for i, card := range d {
            fmt.Println(i, card)
        }
    }

## Range of Slice

 [starIndexIncluding, endIndexNotIncluding]

## byte slice = string
ASCII table ascii.com

"Hi there!" = [72 105 32 116 104 101 114 101 33]

    []byte("Hi there!")
