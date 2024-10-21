package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado \n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
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
	guilherme.Desativar()
	minhaEmpresa := Empresa{}
	Desativacao(guilherme)
	Desativacao(minhaEmpresa)
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", guilherme.Nome, guilherme.Idade, guilherme.Ativo)
}
