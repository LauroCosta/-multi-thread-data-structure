package iterativo

func Multiplica(a, b [][]int) [][]int {

	resultado := make([][]int, len(a))
	for i := 0; i < len(a); i++ {

		resultado[i] = make([]int, len(a))

	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			for k := 0; k < len(a); k++ {
				resultado[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return resultado
}
