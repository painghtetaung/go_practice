package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

type admin struct {
	user  // embedded type
	level string
}

type notifier interface {
	notify()
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{name: "John", email: "john@example.com"}
	a := admin{
		user:  u,
		level: "super",
	}
	sendNotification(&a)
	a.notify()
	// fmt.Println(&21, "21")
}

// package main

// import "fmt"

// // Define interface
// type Animal interface {
//     Speak() string
// }

// // Type 1
// type Dog struct{}
// func (d Dog) Speak() string {
//     return "Woof!"
// }

// // Type 2
// type Cat struct{}
// func (c Cat) Speak() string {
//     return "Meow!"
// }

// // Function that uses the interface
// func makeItSpeak(a Animal) {
//     fmt.Println(a.Speak())
// }

// func main() {
//     var d Dog
//     var c Cat

//     makeItSpeak(d) // Woof!
//     makeItSpeak(c) // Meow!
// }
