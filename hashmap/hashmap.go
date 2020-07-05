package main

import (
	"errors"
	"fmt"

	LDA "trabalho/utils/ListaDuplamenteEncadeada"
)

type HashMap struct {
	tamanhoHash  int
	qtdElementos int
	buckets      []LDA.ListaEncadeada
}

func main() {

	h := NovoHashMap(10)

	h.Add(1, 33)
	h.Add(2, 33)
	h.Add(3, 33)
	h.Add(4, 33)
	h.Add(5, 33)
	h.Add(6, 33)
	h.Add(7, 33)

	fmt.Println(h.Buscar(5))
	fmt.Println(h.Remover(5))
	fmt.Println(h.Remover(6565))

}

func (h *HashMap) hash(chave int) int {

	x := chave * 37 / 64
	return x % h.tamanhoHash

}

func NovoHashMap(tamanho int) *HashMap {

	h := new(HashMap)

	h.tamanhoHash = tamanho
	h.qtdElementos = 0
	h.buckets = make([]LDA.ListaEncadeada, tamanho)

	return h
}

func (h *HashMap) Add(chave int, valor int) error {

	indice := h.hash(chave)

	if h.buckets[indice].Tamanho == 0 {
		h.buckets[indice].AddInicio(chave, valor)
		h.qtdElementos++
	} else if h.buckets[indice].Buscar(chave) == -1 {
		h.buckets[indice].AddInicio(chave, valor)
		h.qtdElementos++
	} else {
		return errors.New("Já existe um elemento com está chave")
	}

	return nil
}

func (h *HashMap) Remover(chave int) error {

	indice := h.hash(chave)

	if h.buckets[indice].Tamanho == 0 {
		return errors.New("Lista está vazia ")
	}

	if h.buckets[indice].Buscar(chave) != -1 {
		h.buckets[indice].Remover(chave)
		h.qtdElementos--
	} else {
		return errors.New("Não existe esta chave na hash")
	}

	return nil
}

func (h *HashMap) Buscar(chave int) int {

	indice := h.hash(chave)
	valor := h.buckets[indice].Buscar(chave)
	if valor != -1 {
		return valor
	} else {
		return -1
	}

}
