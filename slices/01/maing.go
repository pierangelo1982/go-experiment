package main

import "fmt"

var sliceNumber []int

func main() {
	//fmt.Println(sliceNumber)
	/*
		for i := 0; i < len(sliceNumber); i++ {
			fmt.Println(i)
		}
	*/

	for i := 0; i < 30; i++ {
		sliceNumber := append(sliceNumber, i)
	}
	fmt.Println(sliceNumber)
}
