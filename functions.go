package main

import "fmt"

// func cool() string {
// 	return "hi"
// }

// func main() {
// 	interestingString := cool()
// 	fmt.Printf(interestingString)
// }

// //functions are obviously the best
// //like TS but with no colons
// func sub(x int, y int) int {
// 	return x - y
// }

//Which of the following is the most succinct way to write a function signature
//func createUser(firstName, lastName string) // [can declare both to be strings like this apparently]

//What are we talking about when we discuss declaration syntax?
//The style of language used to create new variables, types, and functions

//Go reads left to right

//What is f func(func(int,int)int,int)int?
//A function that takes in a function and int and returns an int

////////////////////////////////////////////////////////////////////////////
//Something about Memory

//x:=5
//5 stored in memory -- x is a pointer to that 5
//y:=x actually allocates new mem for a y variable and gives it initial value x (copy of x)
//So in go, assigning y := x actually makes a whole new variable unlike in JS (where it would point to same spot still [reference])

// func main() {
// 	sendsSoFar := 430
// 	const sendsToAdd = 25
// 	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
// 	fmt.Println("you've sent", sendsSoFar, "messages")
// }

// func incrementSends(sendsSoFar int, sendsToAdd int) int {
// 	sendsSoFar = sendsSoFar + sendsToAdd
// 	return sendsSoFar
// }

//You can ignore values in go with an underscore
//ex. func getPoint()(x int, y int) int { -> okay... i guess (x int, y int) represents... two returns? or wtf is the other int then
// 	return 3,4
// }

//x, _ := getPoint()

func old55() {
	firstName, _ := getNames()
	fmt.Println("Welcom to textio, ", firstName)
}

func getNames() (string, string) {
	return "John", "Doe"
}

//can initialize functions with return values... uh... wtf...
//func getCoords() (x, y int){
// 	return // automatically returns x and y (and both will be 0)
// }

//Use when you want to make it easier to understand what you're returning

//explicit returns probably better in most scenarios
//func getCoords() (x, y int){
// 	return x, y
// }

//When should naked returns be used
//Never (lol)

//When should named returns be used
//When there are many values being returned (but really, there's no real reason not to always do it)

//Guard clauses provide a linear approach to logic trees :)
//Guard clauses are early returns - very common to use in go. I love them and always use them
