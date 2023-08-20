package main

//interfaces are collections of method signatures -- a type "implements" an interface if it has all of the methods of the given interface
//defined on it.

//both structs below, "implement" the shape interface

// type shape interface{
// 	area() float64
// 	perimeter() float64
// }

// type rect struct {
// 	width, height float64
// }
// func (r rect) area() float64 {
// 	return r.width * r.height
// }
// func (r rect) perimeter() float64 {
// 	return 2*r.width + 2*r.height
// }

// type circle struct {
// 	radius float64
// }
// func (c circle) area() float64{
// 	return math.Pi * c.radius * c.radius
// }
// func (c circle) perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func sendMessage(msg message) {
// 	fmt.Println(msg.getMessage())
// }

// type message interface {
// 	getMessage() string
// }

// type BirthdayMessage struct {
// 	birthdayTime  time.Time
// 	recipientName string
// }

// func (bm BirthdayMessage) getMessage() string {
// 	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime)
// }

// type sendingReport struct {
// 	reportName    string
// 	numberOfSends int
// }

// func (sr sendingReport) getMessage() string {
// 	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v something`, sr.reportName, sr.numberOfSends)
// }

//because both fufill "message", we can feed both structs into sendMessage.
//That's the idea... huh...
//interfaces implemented implicitly (unique to go, java for example, needs to say it implements the interface)

// type employee interface {
// 	getName() string
// 	getSalary() int
// }

// type contractor struct {
// 	name string
// 	hourlyPay int
// 	hoursPerYear int
// }
// func (c contractor) getName() string {
// 	return c.name
// }
// func (c contractor) getSalary() int {
// 	return c.hourlyPay * c.hoursPerYear
// }

//how is an interface fufilled?
//implicitly if a type has all the required methods defined on it

//can a type fufill multiple interfaces?
//yes

//go uses the ___ keyword to show that a type implements an interface
//interfaces are impl,emented implicitly so no keyword

// func (e email) cost() float64{
// 	//?
// }

// func (e email) print() float64{

// }

// type expense interface {
// 	cost() float64
// }
// type printer interface{
// 	print()
// }

//are you required to name the arguments of an interface?
//no, but we should anyway -- Copy(sourceFile string, destinationFile string)

//why name them?
//for readability

///////////////////////////////////////////////////////////////////
//Type assertion??
// c, ok := s.(circle) //is s a circle? if it is, c = a circle struct instance and ok = true otherwise its false...???

//below, it works like this... email and sms are different structs. that both fufill the interface expense that has "cost()" on it.
//We know both fufill expense, but maybe we wanna do something different with them depending on which struct it actually is.

// func getExpenseReport(e expense) (string, float64){
// 	em,ok := e.(email)
// 	if ok {
// 		return em.toAddress, em.cost()
// 	}
// 	s, ok := e.(sms)
// 	if ok {
// 		return s.toPhoneNumber, s.cost()
// 	}
// 	return "", 0.0
// }

//can also use switch for this.
// func getExpenseReport (e expense) (string, float64){
// 	switch v := e.(type){
// 	case email:
// 		return v.toAddress, v.cost()
// 	case sms:
// 		return v.toPhoneNumber, v.cost()
// 	default:
// 		return "", 0.0
// 	}
// }

//Interfaces should have as _ methods as possible
//FEW methods as possible
