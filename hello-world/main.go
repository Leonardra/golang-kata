package main

import "fmt"

func main() {

	//Using the var keyword to delare a variable
	var whatToSay string = "Hello, World! - Var Keyword"
	greet(whatToSay)

	//Using the := shorthand to declare a variable
	//It implicitly infers the type by the value assigned
	shortHandGreeting := "Hello, World! - Short Hand"
	greet(shortHandGreeting)
}

func greet(greeting string) {
	fmt.Println(greeting)
}
