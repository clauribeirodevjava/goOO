package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	fmt.Println(m["a"])
	delete(m, "c")
	fmt.Println(m["c"])
	_, exists := m["b"]
	fmt.Println(exists)
	value, exists := m["a"]
	println(value)
	var x = map[string]int{}
	fmt.Println(x)

	if value, exists := m["c"]; exists {
		fmt.Println("achou", value)
	} else {
		fmt.Println("ops")
	}

}
