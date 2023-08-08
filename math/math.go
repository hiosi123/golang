package math

import (
	"fmt"
	"math"
)

func MyMath() {
	var x int = 1 // int32 int64
	var x32 int32 = 1
	var x64 int64 = 1
	fmt.Println(x)
	fmt.Println(x32)
	fmt.Println(x64)

	fmt.Printf("%T, %T, %T", x, x32, x64)

	var unsinged uint = 1
	fmt.Println(unsinged)

	pi := 3.141
	fmt.Printf("%T", pi)

	var pi32 float32 = 3.141
	fmt.Printf("%T", pi32)

	pi = float64(3.14)

	//conversion between the two
	a := 1
	b := 3.141

	fmt.Println(a)
	fmt.Println(b)

	a = int(b)
	fmt.Println(a)

	fmt.Println(math.Ceil(4.123123))
	fmt.Println(math.Min(1, 0))
	fmt.Println(math.Max(1, 0))
	fmt.Println(math.Abs(-1))
	fmt.Println(math.Sqrt(100))
	fmt.Println(math.Pow(2, 3))

	fmt.Println(complex(1, 2))

}
