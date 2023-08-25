package main

import "fmt"

type Carro struct {
	Fabricante string
	Modelo     string
	Ano        int
}

type Pessoa struct {
	Nome  string
	Idade int
	Carro Carro
}

func main() {

	pessoa := Pessoa{
		Nome:  "Hanniel",
		Idade: 17,
		Carro: Carro{
			Fabricante: "Toyota",
			Modelo:     "Supra MKV",
			Ano:        2023,
		},
	}

	fmt.Println(pessoa, pessoa.Carro.Fabricante)
}
