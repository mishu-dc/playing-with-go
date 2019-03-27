package main

import "fmt"

const (
	Male   = 1
	Female = 2
)

const (
	KB int = 1024
	MB int = KB * 1024
	GB int = MB * 1024
)

const (
	year1 = iota + 2015
	year2 = iota + 2015
	year3 = iota + 2015
	year4 = iota + 2015
)

func Excercise_002() {
	title := `Ninja Level 2
	Keep going`

	fmt.Println(title)
	x := 911
	fmt.Printf("Decimal: %d\n", x)
	fmt.Printf("Binary: %b\n", x)
	fmt.Printf("Hex: %#X\n", x)

	a := (42 == 42)
	b := (42 <= 42)
	c := (42 >= 42)
	d := (42 != 42)
	e := (42 < 42)
	f := (42 > 42)

	fmt.Println(a, b, c, d, e, f)

	fmt.Printf("Male= %d Type=%T Female = %d Type=%T\n", Male, Male, Female, Female)
	fmt.Printf("KB= %d byte MB = %d byte GB=%d byte\n", KB, MB, GB)

	y := 10
	z := y >> 1

	fmt.Printf("Decimal: %d\t Binary: %b\t Hex: %#X\n", y, y, y)
	fmt.Printf("Decimal: %d\t Binary: %b\t Hex: %#X\n", z, z, z)
	fmt.Printf("%d\t%d\t%d\t%d\n", year1, year2, year3, year4)

}
