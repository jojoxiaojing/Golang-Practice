# goLang-Practice

## Chapter 1 Hello World

#How do we run the code:

(1) go run main.go 
    run the main directly
(2) go build main.go 
    make the file into main and run ./main

#What does 'package main' mean?

(1) Package == Project == Workspace
(2) Types of packages: (a) Executable: generates a file that we can run, must contain a function called "main" and (b)reusable: helpers, put reusable logic
(3) package main -> go build -> produce main

# What does import "fmt" means
to get another package

# what is 'func' means
short for function

# How is the main.go file organized
1.Package declaration : 
        package main
2.Import other packages that we need: 
        import "fmt"
3.Declare functions, tell Go to do things:
        func main(){
            fmt.Println("hi there")
        }