package main

import "fmt"

func main() {
fmt.Println(soma(10, 5))
fmt.Println(subtracao(10, 5))
// fmt.Println(printaNome("Bernardo"))
nome1, nome2 := printaNome("Bernardo")
fmt.Println(nome1, nome2)
fmt.Println(printaNome("Bernardo", "Paz"))
}
	
func soma(a int, b int) int {
	return a + b
}

func subtracao (a int, b int) int {
	return a - b
}	

// func printaNome(nome string) string{
// 	return nome
// }

func printaNome(nome string) (string, string) {
	return nome, nome
}

//Função com letra minuscula é PRIVADA (somente dentro do pacote) -> igual no python que usa _
func printaNomeCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}