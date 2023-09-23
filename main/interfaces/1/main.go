package main

import "fmt"

type Names []interface{}

func (n *Names) load() {
	*n = Names{
		"Wesley",
		"Sarah",
		"Davi",
	}
}

func (n Names) printNames() {
	fmt.Println(n)
}

func main() {
	var name Names
	name.load()
	name.printNames()

}
