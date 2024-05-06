package main

import "fmt"

type No struct {
	valor int
	prox  *No
}

type Fila struct {
	tamanho int
	inicio  *No
	fim     *No
}

func (f *Fila) percorre() {
	if f.inicio == nil {
		fmt.Println("Fila vazia")
	} else {
		no := f.inicio
		for no != nil {
			fmt.Print(no.valor, " -> ")
			no = no.prox
		}
		fmt.Println()
	}
}

func (f *Fila) insere(novoValor int) {
	novoNo := &No{valor: novoValor}

	if f.fim == nil {
		f.inicio = novoNo
		f.fim = novoNo
	} else {
		f.fim.prox = novoNo
		f.fim = novoNo
	}

	f.tamanho++
}

func (f *Fila) remove() error {
	if f.inicio == nil {
		return fmt.Errorf("Fila vazia")
	}

	noRemovido := f.inicio
	f.inicio := f.inicio.prox

	if f.inicio == nil {
		f.fim = nil
	}

	f.tamanho--
	return noRemovido
}

func main() {
	f := Fila{}

	f.insere(7)
	f.insere(2)

	f.percorre()
	fmt.Println("Tamanho da f:", f.tamanho)

	f.remove()

	f.percorre()
	fmt.Println("Tamanho da f:", f.tamanho)

	f.remove()
	f.remove()

	f.percorre()
	fmt.Println("Tamanho da f:", f.tamanho)

	err := f.remove()
	fmt.Println(err)
}
