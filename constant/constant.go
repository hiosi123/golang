package constant

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func Constant() {

	const five int = 5
	fmt.Println(five)

	var seven = 7
	fmt.Println(seven)

	var six int = 6
	fmt.Println(six)

	six = 10
	fmt.Println(six)

	const pi = 3.14
	fmt.Println(pi)

	const (
		a = 1
		b
		c = 3
		d
	)

	fmt.Println(a, b, c, d)

	const (
		zero int = iota + 5

		one

		two
	)

	fmt.Println(zero, one, two)

	const (
		_ = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)

	fmt.Println(kb, mb, gb, tb)

	fmt.Println(math.Pow(2, 100))

	//const largeNum = math.Pow(2, 100)

	fmt.Println(math.Pi)
	fmt.Println(time.February)
	fmt.Println(time.Second)
	fmt.Println(time.UTC)
	fmt.Println(big.MaxExp)
	fmt.Println(big.MinExp)
}
