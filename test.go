package main

import "fmt"

//Basic hello world
// func main() {
// 	//comments
// 	fmt.Println("starting Textio server")
// }

//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////

//:= makes a new variable and assigns it
//we also can declare the type of variable upfront before the statement
// func main() {
// 	messagesFromDoris := []string{
// 		"You doing anything later??",
// 		"hey lol",
// 		"Don't leave me hanging...",
// 		"Cool",
// 	}
// 	numMessages := float64(len(messagesFromDoris))
// 	costPerMessage := .02

// 	totalCost := costPerMessage * numMessages

// 	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
// }

//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////

//go code generally runs FASTER than interpreted languages and compiles FASTER than other compiled languages like C and Rust
//does go execute faster than rust? NO

//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////

//main.go -> gobuild -> helloworld.exe

//variable types
func main() {
	//var username string = "wagslane"
	//var password string = "20947382822"
	username := "wagslane"
	password := "20947382822"

	fmt.Println("Authorization: Basic", username+":"+password)
}
