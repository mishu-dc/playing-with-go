package main

import "fmt"

func Excercise_004() {
	for i := 1; i <= 10000; i++ {
		//fmt.Println(i)
	}

	for i := 65; i < 'A'+26; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

	year := 1982
	for year <= 2019 {
		fmt.Println(year)
		year++
	}

	year = 1982
	for {
		if year > 2019 {
			break
		}
		fmt.Println(year)
		year++
	}

	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}

	var result []int
	result = append(result, 1)
	result = append(result, 2)
	fmt.Println(result)

	switch {
	case 1 == 1:
		fmt.Println("1==1")
	case 1 == 2:
		fmt.Println("1==2")
	}

	fav := "Bond"
	switch fav {
	case "Bond", "007", "James":
		fmt.Println("Bond")
	case "World":
		fmt.Println("World")
	}

}
