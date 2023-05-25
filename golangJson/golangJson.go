package golangJson

import (
	"encoding/json"
	"fmt"
)

type Ninja struct {
	Name   string
	Weapon string
	Level  int
}

func MakeJson() {
	fmt.Println("")

	sFrom := `{"name": "Wallace", "weapon": "ninja star", "level" : 1}`
	fmt.Println(sFrom)
	var Wallace Ninja
	err := json.Unmarshal([]byte(sFrom), &Wallace)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Wallace)

	sTo, err := json.Marshal(Wallace)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sTo))
	fmt.Printf("%T", sTo)
}
