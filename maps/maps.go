package maps

import "fmt"

func MakeMaps() {
	var m map[string]string
	fmt.Println(m == nil)
	m = map[string]string{}
	fmt.Println(m == nil)

	fmt.Println(m)
	fmt.Println(len(m))

	m = make(map[string]string, 5)
	fmt.Println(len(m))

	m = map[string]string{"Wallace": "Programmer"}
	fmt.Println(m)
	fmt.Println(len(m))

	m["Johnny"] = "Not a programmer"
	fmt.Println(m)

	fmt.Println(m["Johnny"])

	delete(m, "Johnny")
	fmt.Println(m)

	m["Wallace"] += " Ninja"
	fmt.Println(m)

	for name, title := range m {
		fmt.Println(name, title)
	}

	title, ok := m["Johnny"]
	if ok {
		fmt.Println(title)
	} else {
		fmt.Println("we didn't find johnny")
	}

}
