package makeP

import (
	"bytes"
	"fmt"
)

func Makemake() {
	// make() & new()
	// built-in functions to allocate memory for a given target type T
	// make(T, args) & new(T)

	//Data Types
	// make() only works with slices, maps, and channels & new() works with any

	s := make([]int, 0)
	m := make(map[int]int)
	c := make(chan bool)

	// := make(int)

	fmt.Println(s, m, c)

	i := new(int)
	fmt.Println(i)

	// Returned Types
	// make() returns the target type & new() returns the address
	var v map[int]int = make(map[int]int)
	var p *map[int]int = new(map[int]int)

	fmt.Println(v, p)

	//Memory Initialization
	// make() initializes memory & new() zeros memory
	var mMake map[int]int = make(map[int]int) // initializes a map
	var nNew *map[int]int = new(map[int]int)  // Zoeros a map pointer to nil

	mMake[0] = 1
	(*nNew)[0] = 1

	// Zero Value
	// Sometimes new()'s zero values are useful rather than just nil
	// The zero value for Buffer is an empty buffer ready to use.
	b1 := new(bytes.Buffer)
	var b2 bytes.Buffer
	fmt.Println(b1.Len())
	fmt.Println(b2.Len())

}

// For new(map[int]int), more precisely, it will return a pointer that
// points to a nil map object rather than return a nil map pointer.
// If you try to print out the exact thing that returns
// from new(map[int]int) it's a valid pointer,
// however it just points to a map object that is not initialized.
