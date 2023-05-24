package structs

import "fmt"

func StructPrep() {
	type Ninja struct {
		Name    string
		Weapons []string
		Level   int
	}

	wallace := Ninja{
		Name:    "wallace",
		Weapons: []string{"ninja star"},
		Level:   5,
	}

	wallace.Weapons = append(wallace.Weapons, "ninja sword")

	// fmt.Println(wallace)

	type Dojo struct {
		Name  string
		Ninja Ninja
	}

	golangDojo := Dojo{
		Name:  "golang dojo",
		Ninja: wallace,
	}

	// fmt.Println(golangDojo)
	// fmt.Println(golangDojo.Ninja.Name)
	golangDojo.Ninja.Level = 3

	fmt.Println(golangDojo.Ninja.Level)
	fmt.Println(golangDojo.Ninja.Name)

	jonnyPointer := new(Ninja)
	// fmt.Println(jonnyPointer)
	// fmt.Println(*jonnyPointer)

	jonnyPointer.Level = 10
	// fmt.Println(jonnyPointer)

	wallace.Level = 12
	// fmt.Println(wallace)

	// golangDojo1 := Dojo{
	// 	Name:  "golang Dojo2",
	// 	Ninja: jonnyPointer,
	// }

	// golangDojo1.Ninja.Name = "hello"
	// fmt.Println(golangDojo1.Ninja)

	// intern := ninjaIntern{
	// 	name:  "intern",
	// 	level: 1,
	// }
	// intern.sayHello()
}

type ninjaIntern struct {
	name  string
	level int
}

func (n *ninjaIntern) sayHello() {
	fmt.Println("hello")
}
