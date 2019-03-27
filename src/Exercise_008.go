package main

import "fmt"

type human struct {
	firstName string
	lastName  string
}

func mainzzz() {
	fmt.Println("FUNCTION")

	fmt.Println(isEven(10))

	x, y := swap(20, 10)

	fmt.Println(x, y)

	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total := sum(vals...)

	defer done()

	fmt.Println(total)

	p := human{
		firstName: "James",
		lastName:  "Bond",
	}

	fmt.Println(p.speak())
}

func (p human) speak() string {
	return (`I am ` + p.firstName + ` ` + p.lastName)
}

func done() {
	fmt.Println("Done")
}

func sum(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}

	return total
}

func swap(x int, y int) (int, int) {
	return y, x
}

func isEven(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
