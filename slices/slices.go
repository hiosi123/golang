package slices

import "fmt"

func MakeSlices() {

	fixed := [...]int{0, 1, 2}
	fmt.Println(fixed)

	a := []int{0, 1}
	fmt.Println(a)
	fmt.Println(len(a))

	a = []int{0, 1, 2}
	fmt.Println(a)
	a = append(a, 1, 2, 3)
	fmt.Println(a)

	fmt.Println(cap(a), len(a))

	b := make([]int, 5)
	fmt.Println(b)

	fmt.Println(a)

	a = a[0:6]

	fmt.Println(a)
	fmt.Println(len(a))

	if a == nil {
	}

}
