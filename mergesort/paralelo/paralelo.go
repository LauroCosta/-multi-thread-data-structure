package mergesortParalelo

import (
	"sync"
	Iterativo "trabalho/mergesort/iterativo"
)

func MergesortParalelo(vetor []int, limite int) []int {

	var l, r []int

	if len(vetor) == 1 {
		return vetor
	}
	if limite == 0 {

		l = Iterativo.Mergesort(vetor[:(len(vetor) / 2)])
		r = Iterativo.Mergesort(vetor[(len(vetor) / 2):])
	} else {

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer func() {
				wg.Done()
			}()
			l = MergesortParalelo(vetor[:(len(vetor)/2)], limite-1)
		}()

		go func() {
			defer func() {
				wg.Done()
			}()
			r = MergesortParalelo(vetor[(len(vetor)/2):], limite-1)
		}()

		wg.Wait()
	}
	return MergeParalelo(r, l, limite-1)
}

func MergeParalelo(A []int, B []int, limite int) []int {

	if limite <= 0 {

		return Iterativo.Merge(A, B)

	} else {

		var wg sync.WaitGroup
		wg.Add(2)

		C := make([]int, len(A)+len(B))

		go func() {
			defer func() {
				wg.Done()
			}()
			for i, a := range A {
				C[i+buscaBinaria(a, B)] = a
			}
		}()

		go func() {
			defer func() {
				wg.Done()
			}()

			for i, b := range B {
				C[i+buscaBinaria(b, A)] = b
			}
		}()
		wg.Wait()
		return C
	}
}

func buscaBinaria(valor int, vetor []int) int {

	inicio := 0
	fim := len(vetor) - 1
	var meio int

	for inicio <= fim {
		meio = (inicio + fim) / 2

		if vetor[meio] == valor {
			return meio
		} else if vetor[meio] < valor {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return inicio
}
