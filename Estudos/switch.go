package main

import "fmt"

func main() {

	posicao := 2

	switch posicao {
	case 1:
		fmt.Println("Primeiro")
	case 2:
		fmt.Println("Segundo")
	case 3:
		fmt.Println("Terceiro")
	default:
		fmt.Println("Outro")
	}

	fruta := "maca"

	switch fruta {
	case "maca":
		fmt.Println("É uma maçã")
	case "Banana":
		fmt.Println("É uma Banana")
	case "laranja":
		fmt.Println("É uma laranja")
	default:
		fmt.Println("Não é uma fruta")
	}
}