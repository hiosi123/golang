package typeAssertion

import "fmt"

type ninjaStar struct {
	owner string
}

type ninjaSword struct {
	owner string
}

func (ninjaStar) attack() {
	fmt.Println("Throwing ninja star")
}

func (ninjaSword) attack() {
	fmt.Println("Swinging ninja sword")
}

func (ninjaStar) pickUp() {
	fmt.Println("picking back up star")
}

type ninjaWeapone interface {
	attack()
}

func TypeAssertion() {
	weapons := []ninjaWeapone{
		ninjaStar{
			owner: "hong",
		},
		ninjaSword{
			owner: "king",
		},
	}

	for _, weapon := range weapons {
		// ns, ok := wepon.(ninjaStar) // type assertion can be checked as err OK
		// if ok {
		// 	ns.pickUp()
		// }
		weapon.attack()
		switch weapon.(type) {
		case ninjaStar:
			ns := weapon.(ninjaStar)
			ns.pickUp()
		case ninjaSword:
			fmt.Println("no need to pick it back up")
		}
	}

	go func() {
		fmt.Println("asd")
	}()

	go func() {

	}()

	go func() {

	}()

}
