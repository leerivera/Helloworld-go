package main

import (
	"fmt"
)

func main() {
	fmt.Println("We going")

	type ninja struct {
		name    string
		weapons []string
		level   int
	}

	wallace := ninja{name: "wallace"}
	wallace = ninja{
		name:    "Wallace",
		weapons: []string{"Ninja Star"},
		level:   1,
	}
	fmt.Println(wallace)
	fmt.Println(wallace.name)
	fmt.Println(wallace.weapons)
	fmt.Println(wallace.level)
	wallace.level++
	wallace.weapons = append(wallace.weapons, "Bo Staff")
	fmt.Println(wallace)

	type dojo struct {
		name  string
		ninja ninja
	}

	deadlyKicks := dojo{
		name:  "Deadly Kicks",
		ninja: wallace,
	}
	fmt.Println(deadlyKicks)
	fmt.Println(deadlyKicks.ninja.level)

	type addressedDojo struct {
		name  string
		ninja *ninja
	}
	jonnyPointer := new(ninja)
	fmt.Println(jonnyPointer)
	fmt.Println(*jonnyPointer)
	(*jonnyPointer).name = "Jonny"
	jonnyPointer.weapons = []string{"Katana"}
	jonnyPointer.level = 1
	fmt.Println(*jonnyPointer)

	intern := ninjaIntern{name: "Intern", level: 1}
	intern.sayHello()
}

type ninjaIntern struct {
	name  string
	level int
}

func (ninjaIntern) sayHello() {
	fmt.Println("High Kick")
}
