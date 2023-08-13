package exception

import "fmt"

func exception() {
	defer safeExit()

	const one = 1
	if one != 1 {
		panic("one is not 1, this isn't good")
	}

}

func safeExit() {
	if r := recover(); r != nil {
		fmt.Println("panic is recovered!")
	}
}
