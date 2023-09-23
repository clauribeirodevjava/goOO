package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

func (c Car) info() string {
	return fmt.Sprintf("Car : %s\nYear: %d\nColor: %s", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{
		"Corolla",
		2017,
		"White",
	}
	car2 := Car{
		"bmw",
		2017,
		"White",
	}

	fmt.Println(car1.Name)
	fmt.Println(car2.Name)
	fmt.Println(car2.info())

}
