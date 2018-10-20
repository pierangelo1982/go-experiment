package main

import "fmt"

type skill struct {
	Power string
}

type hero struct {
	Name  string
	Alias string
}

type heroSkill struct {
	hero
	skill
}

func main() {
	hs1 := heroSkill{
		hero: hero{
			Name:  "Clark Kent / Kal-El",
			Alias: "Superman",
		},
		skill: skill{
			Power: "a lot of cool superpowers",
		},
	}

	hs2 := heroSkill{
		hero: hero{
			Name:  "Steve Rogers",
			Alias: "Capitan America",
		},
		skill: skill{
			Power: "strong",
		},
	}

	fmt.Printf("the best skill of %s is %s \n", hs1.hero.Alias, hs1.skill.Power)
	// nota il diverso richiamo, senza dover richiamare la sub struct
	fmt.Printf("the best skill of %s is %s \n", hs2.Alias, hs2.Power)
}
