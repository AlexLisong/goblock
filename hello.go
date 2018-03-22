package main

import (
	"fmt"
	"reflect"
)

var(
	name, course string
	module float64
	nickname, address, length = "Nigel", "qq Nicholson", 34
)
type Salutation struct{
	name string
	greeting string
}

func Greet(salu Salutation, printMe func(string)){
	printMe(salu.name)
	printMe(salu.greeting)

}
func printMe(msg string){
	fmt.Println(msg)	
}
func main(){
//	var s = Salutation("bbbb")
	var s = Salutation{"bbbb","qqq"}
	Greet(s,printMe)
	name = "Alex"
	fmt.Println("name is", name, "and is of type", reflect.TypeOf(name), nickname)
	message := "my message here"
	var pmessage = "ccc"
	var message_address *string = &message
	var pmessage_address *string = &pmessage
	fmt.Println(message, message_address, pmessage, pmessage_address)
	message = "my new message here"	
	fmt.Println(message, message_address, pmessage, pmessage_address)
	pmessage = "my 3rd message here"	
	fmt.Println(message, message_address, pmessage, pmessage_address)
}