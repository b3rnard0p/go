package main

import "fmt"

func main(){
	anoNasc := map[string]int{
		"Ana": 2000,
		"Be": 2004,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Ana"])
	fmt.Println(anoNasc["Be"])
	anoNasc["Pedro"] = 2005
	fmt.Println(anoNasc)
}