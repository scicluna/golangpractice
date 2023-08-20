package main

//structured data types (very cool)
// type car struct {
// 	Make   string
// 	Model  string
// 	Height int
// 	Width  int
// }

// type messageToSend struct {
// 	phoneNumber int
// 	message     string
// }

// func test(m messageToSend) {
// 	fmt.Printf("Sending message: '%s' to %v\n", m.message, m.phoneNumber)
// 	fmt.Println("=============================")
// }

// func main() {
// 	test(messageToSend{
// 		phoneNumber: 148255510981,
// 		message:     "thanks for signing up",
// 	})
// }

//pretty cool. their keys can hold any type including other structs
//can assign variables to structs like...
//myCar := car{}

// type messageToSend struct {
// 	message   string
// 	sender    user
// 	recipient user
// }

// type user struct {
// 	name   string
// 	number int
// }

// func canSendMessage(mToSend messageToSend) bool {
// 	if mToSend.sender.name == "" || mToSend.recipient.name == "" {
// 		return false
// 	}
// 	if mToSend.sender.number == 0 || mToSend.recipient.number == 0 {
// 		return false
// 	}
// 	return true
// }

//What's an anonymous struct?
//myCar := struct {
// 	Make string
// 	Model string
// } {
// 	Make: 'tesla'
// 	Model: 'model 3'
// }

//When to use anonymous structs?
//When the object is a one-of
//But probably just always name them tbh

//////////////////////////////////////////////////////////
//can embed structs
//like super kind of -> just prevents us from needing to rewrite the make/model in this example

//type car struct {
// 	make string
// 	model string
// }
// type truck struct {
// 	car
// 	bedSize int
// }

//nested structs (key same name as type)
// lanesTruck := truck{
// 	bedSize: 10,
// 	car: car{
// 		make:"toyota",
// 		model: "camry",
// 	},
// }

//however, . operator applies directly
//lanesTruck.make == "toyota"
//lanesTruck.model == "camry"

//example of embedded struct
// type sender struct {
// 	user
// 	rateLimit int
// }

// type user struct {
// 	name   string
// 	number int
// }

//public/private fields -- private fields embedded for easy removal later

///////////////////////////////////////////////////////////////////////
//Methods on structs

// type rect struct {
// 	width int
// 	height int
// }

// //weird "receiver" thing where (r rect) is a special thing... idk... its what makes it a method though
// func (r rect) area () int {
// 	return r.width & r.height
// }

// rec := rect{
// 	width:5,e
// 	height:10,
// }
// fmt.Println(rec.area())

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("=======================")
}

func (authInfo authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", authInfo.username, authInfo.password)
}

func asdasdsadsadasd() {
	authInfo := authenticationInfo{
		username: "bob",
		password: "cool",
	}
	test(authInfo)
}
