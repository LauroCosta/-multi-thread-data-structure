package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	iterativo "trabalho/mergesort/iterativo"
	paralelo "trabalho/mergesort/paralelo"
	escrever "trabalho/utils/escreverArquivo"
)

func main() {

	conteudoArquivo := ""

	amostra := 10
	_, tempoNormal := testeAlgoritimo(false, amostra, 33)
	_, tempoParalelo := testeAlgoritimo(true, amostra, 33)

	conteudoArquivo = conteudoArquivo + strconv.Itoa(amostra) + "," + tempoNormal + "," + tempoParalelo + "\n"

	amostra = 100
	_, tempoNormal = testeAlgoritimo(false, amostra, 33)
	_, tempoParalelo = testeAlgoritimo(true, amostra, 33)
	conteudoArquivo = conteudoArquivo + strconv.Itoa(amostra) + "," + tempoNormal + "," + tempoParalelo + "\n"

	amostra = 1000
	_, tempoNormal = testeAlgoritimo(false, amostra, 33)
	_, tempoParalelo = testeAlgoritimo(true, amostra, 33)
	conteudoArquivo = conteudoArquivo + strconv.Itoa(amostra) + "," + tempoNormal + "," + tempoParalelo + "\n"

	amostra = 10000
	_, tempoNormal = testeAlgoritimo(false, amostra, 33)
	_, tempoParalelo = testeAlgoritimo(true, amostra, 33)
	conteudoArquivo = conteudoArquivo + strconv.Itoa(amostra) + "," + tempoNormal + "," + tempoParalelo + "\n"

	amostra = 100000
	_, tempoNormal = testeAlgoritimo(false, amostra, 33)
	_, tempoParalelo = testeAlgoritimo(true, amostra, 33)
	conteudoArquivo = conteudoArquivo + strconv.Itoa(amostra) + "," + tempoNormal + "," + tempoParalelo + "\n"

	amostra = 1000000
	_, tempoNormal = testeAlgoritimo(false, amostra, 33)
	_, tempoParalelo = testeAlgoritimo(true, amostra, 33)
	conteudoArquivo = conteudoArquivo + strconv.Itoa(amostra) + "," + tempoNormal + "," + tempoParalelo + "\n"

	fmt.Println(conteudoArquivo)
	escrever.EscreverTexto(conteudoArquivo, "mergesort.txt")
}

func compararVetor(A []int, B []int) bool {

	i := 0
	for i < len(A) {
		if A[i] != B[i] {
			return false
		}
		i++
	}
	return true
}

func testeAlgoritimo(isParalelo bool, amostra int, qtdExe int) ([]int, string) {

	total := 0.0
	var vetor []int
	i := 0

	for i < qtdExe {
		start := time.Now()

		if isParalelo {
			vetor = paralelo.MergesortParalelo(rand.Perm(amostra), 3)
		} else {
			vetor = iterativo.Mergesort(rand.Perm(amostra))
		}

		t := time.Now()
		elapsed := t.Sub(start)
		total = total + elapsed.Seconds()
		i++
	}

	if isParalelo {
		fmt.Println("Execução algoritimo paralelo")
	} else {
		fmt.Println("Execução algoritimo iterativo")
	}
	fmt.Println("tamanho amostra: ", amostra)
	fmt.Printf("Total tempo de execução %.6fs\n", total)
	fmt.Printf("Média %.6fs\n\n", total/float64(qtdExe))

	return vetor, strconv.FormatFloat(total, 'f', 5, 64)
}
