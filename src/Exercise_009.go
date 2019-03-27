package main

import "fmt"

type member struct {
	firstName string
	lastName  string
}

type student struct {
	member
	roll int
}

type teacher struct {
	member
	designation string
}

type homosapien interface {
	speak()
}

func details(h homosapien) {
	switch h.(type) {
	case member:
		fmt.Println("member talking")
	case student:
		fmt.Println("student talking")
	case teacher:
		fmt.Println("teacher talking")

	}
}

func (t member) speak() {
	fmt.Printf("I am %T\n", t)
	fmt.Println(t)
}

func (t student) speak() {
	fmt.Printf("I am %T\n", t)
	fmt.Println(t)
}

func (t teacher) speak() {
	fmt.Printf("I am %T\n", t)
	fmt.Println(t)
}

func square(x int) func(x int) int {
	return func(x int) int {
		return x * x
	}
}

func getSum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}

func odd(f func(...int) int, xi ...int) int {
	yi := []int{}
	for _, v := range xi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {

	var temp int
	temp = 50
	fmt.Println(temp)
	fmt.Println(&temp)
	fmt.Println(*&temp)

	a := increment()
	b := increment()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := getSum(xi...)
	fmt.Println(sum)

	sum = odd(getSum, xi...)
	fmt.Println(sum)

	//Annonymous function
	func() {
		fmt.Println("I have no name :(")
	}()

	//Annonymous function with parameter
	func(val int) {
		fmt.Println("I have no name :( but have parameter :) ", val)
	}(100)

	f := func(val int) {
		fmt.Println(val)
	}

	//returning function from another function

	rf := square(5)
	x := rf(5)
	fmt.Println(x)

	f(100)

	m := member{
		firstName: "Faisal",
		lastName:  "Ahmed",
	}

	s := student{
		member: member{
			firstName: "Nusrat",
			lastName:  "Sharmin",
		},
		roll: 125,
	}

	t := teacher{
		member: member{
			firstName: "David",
			lastName:  "Copperfield",
		},
		designation: "Professor",
	}

	m.speak()
	s.speak()
	t.speak()

	details(m)
	details(s)
	details(t)
}
