package generics

import "fmt"

func MakeGenerics() {

	// 3main features
	// 1. Type parameter (with constraint)

	// 2. Type inference ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ 물결만 넣어주면 해결됨. 내가 만든 float64 가 있다면

	// 3. Type set

	fmt.Println(minInt(1, 2))
	fmt.Println(minFloat64(0.1, 0.2))

	// fmt.Println(min[int](1, 2))
	fmt.Println(min[float64](0.1, 0.2))

	type superFloat float64
	var sf superFloat = 0.3
	fmt.Println(min(sf, 0.2))

}

type minTypes interface {
	~float64 | int
}

func min[T minTypes](a T, b T) T {
	if a < b {
		return a
	}

	return b
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func minFloat64(a float64, b float64) float64 {
	if a < b {
		return a
	}

	return b
}
