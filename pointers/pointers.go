package pointers

import "fmt"

//package level
var five = 5

func PointerFunc() {

	x := 1
	p := &x

	fmt.Println(p)
	fmt.Println(*p)

	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius = 0

	c = celsius((f - 32) * 5 / 32)

	fmt.Println(f, c)

}

// block => package  =>  universe
