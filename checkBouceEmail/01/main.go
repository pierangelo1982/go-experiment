package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	err := checkmail.ValidateHost("contact@mecadiesel.fr")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("email ok")
	}
}
