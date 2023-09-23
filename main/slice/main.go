package main

import "fmt"

func main() {
	//slice := make([]int, 2, 5)
	//slice[0] = 1
	//slice = append(slice, 10, 2, 5, 6)
	//fmt.Println(slice)
	//fmt.Println(len(slice))
	//fmt.Println(cap(slice))
	/* 	for i := 0; i < 20; i++ {
		slice = append(slice, i)
		fmt.Println("Len:", len(slice), "Cap:", cap(slice))
	} */
	//sliceTest := slice
	//slice = append(slice, 1, 2, 3, 4, 5, 6)
	//como dobrou de tamanho agora slice aponta
	//para um novo array
	//slice[0] = 10

	//fmt.Println(sliceTest)
	//fmt.Println(slice)
	sliceString := []string{
		"helow",
		"World",
		"Much",
		"Better",
	}
	fmt.Println(sliceString)
	fmt.Println(sliceString[2:4])
	fmt.Println(sliceString[2:])

}
