package main

import "fmt"

var numberArray [10]int

var x int
var y int

func main() {
	x = 0
	y = 1
	for i := 0; i < len(numberArray); i++ {
		numberArray[i] = y + x
		x = y
		y = numberArray[i]
	}

	fmt.Println(numberArray)
}
