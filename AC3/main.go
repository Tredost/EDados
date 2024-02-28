package main

import (
	"fmt"
)

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}

func adicionarContato(contato Contato, arrayContatos []*Contato) {
	for i, c := range arrayContatos {
		if c == nil {
			arrayContatos[i] = &contato
			fmt.Println("Contato adicionado com sucesso!")
			return
		}
	}
	fmt.Println("Array está cheio. Delete um contato primeiro!")
}

func excluirContato(arrayContatos []*Contato) {
	for i := len(arrayContatos) - 1; i >= 0; i-- {
		if arrayContatos[i] != nil {
			arrayContatos[i] = nil
			fmt.Println("Último contato excluído com sucesso!")
			return
		}
	}
	fmt.Println("Não há contatos para excluir.")
}

func listarContatos(arrayContatos []*Contato) {
	fmt.Println("\nLista de contatos:")
	for _, c := range arrayContatos {
		if c != nil {
			fmt.Printf("Nome: %s, Email: %s\n", c.Nome, c.Email)
		}
	}
}

func modificarEmail(arrayContatos []*Contato) {
	var nome, novoEmail string
	fmt.Print("Nome do contato: ")
	fmt.Scanln(&nome)
	fmt.Print("Novo email: ")
	fmt.Scanln(&novoEmail)

	for _, c := range arrayContatos {
		if c != nil && c.Nome == nome {
			c.AlterarEmail(novoEmail)
			fmt.Println("E-mail modificado com sucesso!")
			return
		}
	}
	fmt.Println("Contato não encontrado.")
}

func main() {
	arrayContatos := make([]*Contato, 5)

	for {
		var opcao int
		fmt.Println("\nEscolha uma opção:")
		fmt.Println("1. Adicionar contato")
		fmt.Println("2. Excluir último contato")
		fmt.Println("3. Listar contatos")
		fmt.Println("4. Modificar e-mail de um contato")
		fmt.Println("5. Sair")

		fmt.Print("Opção: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			var nome, email string
			fmt.Print("Nome do contato: ")
			fmt.Scanln(&nome)
			fmt.Print("Email do contato: ")
			fmt.Scanln(&email)
			novoContato := Contato{Nome: nome, Email: email}
			adicionarContato(novoContato, arrayContatos)
		case 2:
			excluirContato(arrayContatos)
		case 3:
			listarContatos(arrayContatos)
		case 4:
			modificarEmail(arrayContatos)
		case 5:
			fmt.Println("Programa encerrado!")
			return
		default:
			fmt.Println("Opção inválida! Tente novamente.")
		}
	}
}
