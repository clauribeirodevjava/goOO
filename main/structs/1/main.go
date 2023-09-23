package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	CanFly bool
	Name   string
}

func (c Car) info() string {
	return fmt.Sprintf("Car : %s\nYear: %d\nColor: %s", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{
		Name:  "Corolla",
		Year:  2017,
		Color: "White",
	}
	car2 := Car{
		"bmw",
		2017,
		"White",
	}
	sCar := SuperCar{
		Car:    car1,
		CanFly: true,
		Name:   "Elantra",
	}

	fmt.Println(car1.Name)
	fmt.Println(car2.Name)
	fmt.Println(car2.info())
	fmt.Println(sCar)
	fmt.Println("==========")
	fmt.Println(sCar.Name)

}
