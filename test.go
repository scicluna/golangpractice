package main

import "fmt"

// func main() {
// 	//comments
// 	fmt.Println("starting Textio server")
// }

func main() {
	messagesFromDoris := []string{
		"You doing anything later??",
		"hey lol",
		"Don't leave me hanging...",
		"Cool",
	}
	numMessages := float64(len(messagesFromDoris))
	costPerMessage := .02

	totalCost := costPerMessage * numMessages

	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}
