package main

import "fmt"

func main(){

	sum := 0
	for i := 0; i < 10; i++{ //Go só tem o for de loop
		sum += i
	}
	fmt.Println("Sum:", sum)

	// for {
	// 	fmt.Println("Loop infinito")
	// }

	frutas := []string{"Melancia", "Laranja", "Limão"}
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}
	
}