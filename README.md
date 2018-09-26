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

# Chapter Two -- Begin a simple project
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
