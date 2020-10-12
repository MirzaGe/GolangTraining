go package main

import "fmt"

func main() {
	greet("Jane")
	greet("John")
}

func greet(name string) {
	fmt.Println(name)
}

// greet is declared with a paramete
// when calling greet, pass in an argument
