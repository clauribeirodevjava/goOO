package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"Model"`
	Year  int    `json:"-"`
	color string //a lowercase letter
}

func main() {
	car := Car{"Gol", 2017, "Yellow"}

	result, _ := json.Marshal(car)
	fmt.Println(string(result))
}
