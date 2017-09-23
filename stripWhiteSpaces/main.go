package main

import (
	"fmt"
	"strings"
)

func main() {
	x := stripSpaces("Hello World!")
	fmt.Println(x)
}

func stripSpaces(words string) string {
	x := strings.Replace(words, " ", "", -1)
	return x
}
