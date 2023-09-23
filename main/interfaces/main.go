package main

import "fmt"

type vehicle interface {
	start() string
}

type car struct {
	name string
}
type motorCycle struct {
	name string
}

func (c car) start() string {
	return "The car " + c.name + " has benn started"
}
func (mc motorCycle) start() string {
	return "The motorcycle " + mc.name + " has benn started"
}
func startVechicle(v vehicle) string {
	return v.start()
}
func main() {
	c := car{"Gol"}
	//fmt.Println(c.startCar())
	m := motorCycle{"xpto"}
	fmt.Println(startVechicle(c))
	fmt.Println(startVechicle(m))

}
