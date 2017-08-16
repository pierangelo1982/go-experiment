package main

import "fmt"

const p float64 = 3.14 // pi greco

var radius float64
var diameter float64
var tot float64
var idCalculation int

func main() {
	switchCalculation()
}

// choice between radius or diameter
func switchCalculation() {
	fmt.Println("Select 1 for Radius\nSelect 2 for Diameter")
	fmt.Scan(&idCalculation)
	switch {
	case idCalculation == 1:
		fmt.Println(withRadius())
	case idCalculation == 2:
		fmt.Println(withDiameter())
	}
}

// calcolate with radius
func withRadius() float64 {
	fmt.Println("Insert radius: ")
	fmt.Scan(&radius)
	tot := (radius + radius) * p
	return tot
}

// calculate with diameter
func withDiameter() float64 {
	fmt.Println("Insert diameter:")
	fmt.Scan(&diameter)
	tot := diameter * p
	return tot
}
