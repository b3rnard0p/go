package main

import "fmt"

func main(){

	valor := 1

	if valor == 1 {
		fmt.Println("Valor é 1")
	} else {
		fmt.Println("Valor não é 1")
	}

	numero := 2

	if numero == 1 {
		fmt.Println("É 1")
	} else if numero == 2 {
		fmt.Println("É 2")
	} else {
		fmt.Println("Não é 1 nem 2")
	}

	if 7%2 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é impar")
	}

	nome:= "Bernardo"
	if nome == "Bernardo"{
		fmt.Println("Nome é Bernardo")
	} else {
		fmt.Println("Nome não é Bernardo")
	}
}