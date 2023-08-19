package main

import "fmt"

//bool
//string
//ints, floats
//bools
//int8, int16, int32, int64
//uint, uint8, uint16, uint32, uint64

//wtf is the difference between a uint8 and an int8? Wtf does that mean? Why use this over int/uint ? wtf is a uint? for memory? okay
//byte = uint8
//rune = int32 (unicode code point)

//float32, float64 ???
// func main() {
// 	var number int
// 	var pi float64 = 3.14159

// 	fmt.Printf("%v \n", number)
// 	fmt.Printf("%f \n", pi)
// }

//infer variables
// func main() {
// 	congrats := "Congratulations!" //its a string strongly typed now, but inferred -- := initializes and assigns
// 	fmt.Println(congrats)
// }

//value shifting from inferred
// func main() {
// 	penniesPerText := 2.0                                            //change from 2 to 2.0 to make it a float64
// 	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText) //%T is type of (weird)
// }

// func main() {
// 	averageOpenRate, displayMessage := .23, "is the average open rate of your messages" //can assign two on the same line if we want
// 	fmt.Println(averageOpenRate, displayMessage)
// }

//int int8 int16 int32 int64 == whole numbers
//uint uint8 uint16 uint32 uint64 == positive whole numbers
//float32 float64 == decimal numbers
//complex64 complex128 == imaginary numbers

//always just use int/uint unless we are really hyperoptimizing
//always use int, uint, float64, bool, string

//can shift types like
//temp := 88
//tempFloat := float64(temp)

// func main() {
// 	accountAge := 2.6
// 	accountAgeInt := int(accountAge)

// 	fmt.Println("Your account has existed for", accountAgeInt, "years")
// }

//when should you elect not to use a default type?
//When performance and memory are a concern

//what does the size of a type indicate?
//Bits

//////////////////////////////////////////////////////////////////////
//constants

// func main() {
// 	const premiumPlanName = "Premium Plan"
// 	const basicPlanName = "Basic Plan"

// 	fmt.Println("plan:", premiumPlanName)
// 	fmt.Println("plan:", basicPlanName)
// }

//consts different than in JS
//in go, every value stored in a constant must be known or computed at compile time (not at run time)

//we can still use constants to compute other constants aslong as we dont recieve any info at run time
//ex.
//const firstName = "Lane"
//const lastName = "Wagner"
//const fullName = firstName + " " + lastName
//only valid cuz firstname and lastname are known at compile time.

// func main() {
// 	const secondsInMinute = 60
// 	const minutesInHour = 60
// 	const secondsInHour = secondsInMinute * minutesInHour

// 	fmt.Println("number of seconds in an hour:", secondsInHour)
// }

////////////////////////////////////////////////////////////////////////////////////
//string formatting

//fmt.Printf - prints a formatted string to standard out
//fmt.Sprintf() - returns the formatted string
//%v interpolates the default representation (aka the value placed second)
//ex. fmt.Printf("I am %v years old", 10) == //I am 10 years old

//%s interpolates a string...
//%d interpolates an integer in decimal "form"
//%f interpolates a decimal... but you can do %.2f to for example round it to 2 decimals

// func main() {
// 	const name = "Saul Goodman"
// 	const openRate = 30.5
// 	msg := fmt.Sprintf("hi, %v, your open rate is %.2f percent.", name, openRate) //why cant i use const msg here? cuz fmt.Sprintf is dont at runtime? Idk

// 	fmt.Println(msg)
// }

///////////////////////////////////////////////////////////////////////////////////
//conditionals
//just like in JS only without parens

func oldb() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen < maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println(("Message not sent"))
	}
}

//why would you use the initial section of an if statement?
//if length := getLength(email); length < 1{
// 	fmt.Println("Email is invalid")
// }

//Keeps the code concise and the scope limited (kind of nice syntactic sugar stuff) if we are only using that variable in the comparison
