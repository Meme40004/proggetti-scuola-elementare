package main

import "fmt"

func fattoriale(n int) int {
	x := 1
	for i := 1; i <= n; i++ {
		x = x * i
	}
	return x
}

func main() {
	var n int
	fmt.Println("Inserisci un numero: ")
	fmt.Scanf("%d", &n)
	fmt.Println(fattoriale(n))

}
