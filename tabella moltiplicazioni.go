package main

import "fmt"

func main() {
	var n int
	fmt.Println("Inserisci un numero: ")
	fmt.Scanf("%d", &n)
	for i := 1; i <= 10; i++ {

		fmt.Println(n, "*", i, "=", n*i)

	}

}
