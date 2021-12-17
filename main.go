package main

import "fmt"

func speak() {
	fmt.Print("yeah son!")
}

func main() {
	var m map[string]string
	fmt.Println(m == nil)
	m = map[string]string{}
	fmt.Println(m == nil)
	fmt.Println(m)
	fmt.Println(len(m))
	m = make(map[string]string)
	m = map[string]string{"Nick": "Programmer"}
	fmt.Print(m)
	fmt.Println(m)
	fmt.Println(len(m))
	m["Johnny"] = "Douche"
	fmt.Println(m)
	m["Johnny"] = "No longer douche"
	fmt.Println(m)
	fmt.Println(m["Johnny"])

	fmt.Println(m)
	m["Wallace"] += " Ninja"
	fmt.Println(m)
	m["Johnny"] = "No longer douche"

	for name, title := range m {
		fmt.Println(name, title)
	}
}
