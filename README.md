# goLang-Practice

# Chapter 1 Hello World

### How do we run the code:

* go run main.go

    run the main directly
* go build main.go

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

        package main
* Import other packages that we need: 

        import "fmt"
* Declare functions, tell Go to do things:

        func main(){
            fmt.Println("hi there")
        }