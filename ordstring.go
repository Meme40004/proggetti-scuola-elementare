package main

import "fmt"

func main() {

	n := [6]int{1, 2, 3, 4, 5, 6}

	x := 0

	for i := 0; i < 6; i++ {

		x = x + n[i]

	}
	fmt.Println(x)

}
