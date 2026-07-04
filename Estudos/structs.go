package main

import "fmt"

//Criar uma estrutura de dados própria - Igual C++
//Vários tipos de dados diferentes

type Pessoa struct {
	Nome string
	Sobrenome string
	Idade int
	AnoNasc int
}

type Profissao struct{
	Pessoa
	Tipo string
}

func main(){
	fmt.Println(Pessoa{"Ana", "Sonza", 26, 2000})
	fmt.Println(Pessoa{Nome: "Bernardo", Sobrenome: "Paz", Idade: 21, AnoNasc: 2004}) //Serve para poder por em ordem diferente

	p1:= Pessoa{"Lucas", "Silva", 19, 2005}
	fmt.Println(p1)
	fmt.Println(p1.Nome, p1.Sobrenome)
	fmt.Println(p1.Idade, p1.AnoNasc)
	
	p1.Idade = 20
	fmt.Println(p1.Idade)

	p2 := Pessoa{"Julia", "Souza", 25, 1999}

	pessoas := []Pessoa{}

	pessoas = append(pessoas, p1, p2)
	
	fmt.Println(pessoas)

	//Struct + Map
	alunos := map[string][]Pessoa{}
	alunos["programacao"] = pessoas
	fmt.Println(alunos)

	var alunos2 = map[string][]Pessoa{
		"matematica" : []Pessoa{{Nome: "Carol", Sobrenome: "Nunes", Idade: 23, AnoNasc: 1999}, {Nome: "Pedro", Sobrenome: "Santos", Idade: 28, AnoNasc: 1996}},
		"programacao": []Pessoa{{Nome: "Bernardo", Sobrenome: "Paz", Idade: 21, AnoNasc: 2004}, {Nome: "Ana", Sobrenome: "Sonza", Idade: 26, AnoNasc: 2000}},
	}
	fmt.Println(alunos2)

	//Herança
	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Tipo)
}