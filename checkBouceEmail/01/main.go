package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	err := checkmail.ValidateHost("email@x-unkown-domain.com")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("email ok")
	}
}
