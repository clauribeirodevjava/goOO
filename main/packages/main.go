package main

import (
	"fmt"

	"github.com/clauribeirodevjava/goroutine.git/main/packages/car"
)

func main() {
	car := car.Car{"Gol", "Black"}
	fmt.Println(car.Start())

}
