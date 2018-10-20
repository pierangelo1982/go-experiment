package main

import "fmt"

type skill struct {
	Power string
}

type hero struct {
	Name  string
	Alias string
	skill
}

func main() {
	h1 := hero{
		Name:  "Bruce Wayne",
		Alias: "Batman",
		skill: skill{
			Power: "a lot of gadgets",
		},
	}
	h2 := hero{
		Name:  "Tony Stark",
		Alias: "Iron",
		skill: skill{
			Power: "hi-tech armor",
		},
	}
	fmt.Println("eroe:", h1.Name)
	fmt.Println("eroe:", h2.Name)

	fmt.Printf("the best skill of %s is %s \n", h1.Alias, h1.skill.Power)
	fmt.Printf("the best skill of %s is %s \n", h2.Alias, h2.skill.Power)

	s1 := skill{
		Power: "velocity",
	}

	h3 := hero{
		Name:  "Barry Allen",
		Alias: "Flash",
		skill: skill{
			Power: s1.Power,
		},
	}

	fmt.Printf("the best skill of %s is %s \n", h3.Alias, h3.skill.Power)
}
