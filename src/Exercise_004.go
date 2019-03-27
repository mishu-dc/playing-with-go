package main

import "fmt"

func Xmain() {
	x := []int{4, 5, 6, 9, 8, 1, 2, 5, 1}

	for i, v := range x {
		fmt.Println(i, v)
	}
}
