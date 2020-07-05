package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	divisaoConquista "trabalho/matriz/divisaoConquista"
	iterativo "trabalho/matriz/iterativo"
	paralelo "trabalho/matriz/paralelo"
)

func main() {

	a := GeraMatrix(256)
	b := GeraMatrix(256)

	testeAlgoritimo(1, a, b, 33)
	testeAlgoritimo(2, a, b, 33)
	testeAlgoritimo(3, a, b, 33)

}

func testeAlgoritimo(tipo int, A [][]int, B [][]int, qtdExe int) ([][]int, string) {

	total := 0.0
	var matriz [][]int
	i := 0

	for i < qtdExe {
		start := time.Now()

		switch tipo {
		case 1:
			matriz = iterativo.Multiplica(A, B)
		case 2:
			matriz = divisaoConquista.Multiplica(A, B)
		case 3:
			matriz = paralelo.Multiplica(A, B)
		}

		t := time.Now()
		elapsed := t.Sub(start)
		total = total + elapsed.Seconds()
		i++
	}

	switch tipo {
	case 1:
		fmt.Println("Multiplicação versão iterativa")
	case 2:
		fmt.Println("Multiplicação versão divisão e conquista")
	case 3:
		fmt.Println("Multiplicação versão em paralelo")
	}

	fmt.Println("tamanho Instancia: ", len(matriz))
	fmt.Printf("Total tempo de execução %.6fs\n", total)
	fmt.Printf("Média %.6fs\n\n", total/float64(qtdExe))

	return matriz, strconv.FormatFloat(total, 'f', 5, 64)
}

func GerarNumb(max int) int {

	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	return r.Intn(100)
}

func GeraMatrix(tamanho int) [][]int {

	a := make([][]int, tamanho)
	for i := 0; i < tamanho; i++ {

		a[i] = make([]int, tamanho)

	}

	for i := 0; i < tamanho; i++ {
		for j := 0; j < tamanho; j++ {
			a[i][j] = GerarNumb(100)
		}
	}

	return a

}
