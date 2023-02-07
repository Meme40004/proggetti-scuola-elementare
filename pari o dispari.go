package main

import "fmt"

func pod(n int) string {
	var result string
	if n%2 == 1 {
		result = "dispari"
	} else {
		result = "pari"
	}
	return result
}
func main() {
	var n int
	fmt.Println("Inserisci un numero: ")
	fmt.Scanf("%d", &n)
	fmt.Println(pod(n))

}
