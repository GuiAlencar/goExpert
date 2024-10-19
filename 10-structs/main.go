package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	guilherme := Cliente{
		Nome:  "Guilherme",
		Idade: 26,
		Ativo: true,
	}

	guilherme.Ativo = false
	guilherme.Endereco.Cidade = "Rialma"
	guilherme.Estado = "Goiás"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", guilherme.Nome, guilherme.Idade, guilherme.Ativo)
}
