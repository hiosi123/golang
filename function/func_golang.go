package function

import (
	"errors"
	"fmt"
)

func Function() {
	fmt.Println("hello")

}

func Variadic(intergers ...int) {
	for _, integer := range intergers {
		fmt.Println(integer)
	}
}

func Sum(a, b int) int {
	return a + b
}

func NameLength(name string) (string, int) {
	return name, len(name)
}

func Greeting(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("empty name")
	} else {
		return "Hello " + name, nil
	}
}

func GeertingOk(name string) (string, bool) {
	if len(name) == 0 {
		return "", false
	} else {
		return "Hello " + name, true
	}
}

func Fibonacci(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return Fibonacci(i-1) + Fibonacci(i-2)
	}
}

func WithPointer(i *int) {
	*i = *i + 1
}

func ReturnParam(i int) int {
	return i
}

func ReturnFunction(i int, f func(int) int) int {
	return f(i)
}

func First() {
	fmt.Println("first")
}

func Second() {
	fmt.Println("Second")
}

func CrazyFunction(f func(func()) func()) func(func()) func() {
	return f
}
