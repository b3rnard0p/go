// Arrays e Slices - Homogêneos - Todos os elementos tem o mesmo tipo
// [1, 2, 3, 4, 5] //int
// ["Be","Paz","Freitas"] //string


// Maps - Heterogêneos - Pode misturar tipos (chave-valor)
// map[string] int { "Be": 20}

// map[string] string {"nome": "Be","sobrenome": "Paz"}

//Array
//Tamanho fixo
//Para acessar os valores com o indice: a[0], a[1]
//Função len() embutida - retorna o tamanho

//Slice
//Sem tamanho fixo
//Acessar igual o array
//Função append() para adicionar e tbm tem o len()

package main

import "fmt"

func main() {
	
	//Array
	var array [2]string
	array[0] = "B"
	array[1] = "P"
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array)

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0:4])

	//Slices
	slice := make([]string, 5) //Mesmo que não seja fixo, precisa iniciar com um tamanho, quando for criado vazio
	numPares := []int{0,2,4,6,8,10} // Aqui já é criado com valores, não precisa passar tamanho
	fmt.Println(numPares)
	numPares = append(numPares, 12)
	fmt.Println(numPares)
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice)

	fmt.Println(len(slice))
}