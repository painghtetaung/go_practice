package main

import (
	"fmt"
)

var MAX_CHICKEN_PRICE float32 = 5

type data struct {
	name string
	age int
}

func (d data) displayName() {
	fmt.Printf("\nName: %s, Age: %d", d.name, d.age)
}

func (d *data) setAge(age int) {
	d.age = age
}

func main() {
	var d = data{name: "John", age: 20}
	f1 := d.displayName
	d.name = "Jane"
	f2 := d.setAge
	f2(21)
	fmt.Println("After setAge", d)
	f1()
}
