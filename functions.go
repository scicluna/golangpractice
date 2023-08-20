package main

import "fmt"

func cool() string {
	return "hi"
}

func main() {
	interestingString := cool()
	fmt.Printf(interestingString)
}
