package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

type cricket int

var c cricket

var w int

func Excercise_001() {
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)

	c = 25
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	w = int(c)
	fmt.Println("w =", w)
	fmt.Printf("type of w = %T", w)
}
