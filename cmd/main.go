package main

import (
	"dojo/linkedList"
)

func main() {
	// result := function.Sum(1, 2)
	// fmt.Println(result)

	// result1, ok := function.GeertingOk("hello")
	// if ok {
	// 	fmt.Println(result1)
	// }

	// result3 := function.Fibonacci(10)
	// fmt.Println(result3)

	// param := function.ReturnParam(2)
	// defer fmt.Println(param)

	// function := function.ReturnFunction(1, function.ReturnParam)
	// defer fmt.Println(function)

	l := linkedList.LinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Remove(1)
	

	// math.MyMath()
}
