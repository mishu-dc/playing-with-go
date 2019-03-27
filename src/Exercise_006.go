package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	flavours  []string
}

func mainXXX() {
	m := make(map[string]person, 100)

	fmt.Println("STRUCT")

	persons := []person{{
		firstName: "Faisal",
		lastName:  "Ahmed",
		flavours:  []string{"vanila", "chocolate"},
	},
		{
			firstName: "Nusrat",
			lastName:  "Sharmin",
			flavours:  []string{"strawberry"},
		},
	}

	m["bond"] = person{
		firstName: "James",
		lastName:  "Bond",
		flavours:  []string{"chocokolate"},
	}

	m["pond"] = person{
		firstName: "James",
		lastName:  "Pond",
		flavours:  []string{"Vanila"},
	}

	fmt.Println("Map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	for _, p := range persons {
		fmt.Println(p.firstName)
		fmt.Println(p.lastName)
		for _, f := range p.flavours {
			fmt.Printf("\t%s\t", f)
		}
		fmt.Println()
	}
}
