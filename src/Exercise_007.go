package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func mainyyy() {
	fmt.Println("Struct")

	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}

	a := struct {
		doors     int
		color     string
		awd       bool
		fwd       bool
		accidents map[int]string
		owners    []string
	}{
		doors:     4,
		color:     "black",
		awd:       true,
		fwd:       false,
		owners:    []string{"Faisal", "Nusrat", "Monika"},
		accidents: map[int]string{2019: "front end crash", 2018: "back end crash"},
	}

	fmt.Println(a)
	fmt.Println(t)
	fmt.Println(s)

	fmt.Println(s.doors, s.color, s.luxury)
	fmt.Println(t.doors, t.color, t.fourWheel)
}
