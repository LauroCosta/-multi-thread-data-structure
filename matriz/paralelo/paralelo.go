package paralelo

import (
	"sync"
)

func Multiplica(a, b [][]int) [][]int {

	resultado := make([][]int, len(a))
	for i := 0; i < len(a); i++ {

		resultado[i] = make([]int, len(a))

	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 0; i < len(a)/2; i++ {
			for j := 0; j < len(a); j++ {
				for k := 0; k < len(a); k++ {
					resultado[i][j] += a[i][k] * b[k][j]
				}
			}
		}
	}()

	go func() {
		defer func() {
			wg.Done()
		}()
		for i := len(a) / 2; i < len(a); i++ {
			for j := 0; j < len(a); j++ {
				for k := 0; k < len(a); k++ {
					resultado[i][j] += a[i][k] * b[k][j]
				}
			}
		}
	}()

	wg.Wait()

	return resultado
}

func Paralelo(i int, a, b, c [][]int) [][]int {
	for j := 0; j < len(a); j++ {
		for k := 0; k < len(a); k++ {
			c[i][j] += a[i][k] * b[k][j]
		}
	}
	return c
}
