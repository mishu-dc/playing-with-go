package main

import "fmt"

func mainXX() {
	fmt.Println("Ninja level 5")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)

	fmt.Println(x)

	x = append(x, 53, 54, 55)

	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	states := make([]string, 50, 50)
	states = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(len(states))
	fmt.Println(cap(states))
	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}

	actors := [][]string{{"007", "James Bond"}, {"008", "Peter Parker"}}
	fmt.Println(actors)

	m := map[string][]string{}
	fmt.Println("map:", m)
	m["bond_james"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	m["moneypenny_miss"] = []string{`James Bond`, `Literature`, `Computer Science`}
	m["no_dr"] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	for k, v := range m {
		fmt.Println("for key", k)
		for i, val := range v {
			fmt.Printf("\t%d\t%s\n", i, val)
		}
		fmt.Println()
	}

	delete(m, "bond_james")

	fmt.Println("After delete")

	for k, v := range m {
		fmt.Println("for key", k)
		for i, val := range v {
			fmt.Printf("\t%d\t%s\n", i, val)
		}
		fmt.Println()
	}
}
