package divisaoConquista

func NovaMatrix(tamanho int) [][]int {
	a := make([][]int, tamanho)
	for i := 0; i < tamanho; i++ {

		a[i] = make([]int, tamanho)

	}
	return a
}
func NovaMatrix1(tamanho int, tamanho2 int) [][]int {
	a := make([][]int, tamanho)

	for i := 0; i < tamanho; i++ {

		a[i] = make([]int, tamanho2)

	}
	return a
}

func IkjAlgoritmo(a, b [][]int) [][]int {

	resultado := NovaMatrix(len(a))

	for i := 0; i < len(a); i++ {
		for k := 0; k < len(a); k++ {
			for j := 0; j < len(a); j++ {
				resultado[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return resultado
}

func Multiplica(matrix1, matrix2 [][]int) [][]int {
	tamanho := len(matrix1)
	leafSize := 8

	if tamanho == leafSize {
		return IkjAlgoritmo(matrix1, matrix2)
	} else {
		novoTamanho := tamanho / 2

		a11 := NovaMatrix(novoTamanho)
		a12 := NovaMatrix(novoTamanho)
		a21 := NovaMatrix(novoTamanho)
		a22 := NovaMatrix(novoTamanho)

		b11 := NovaMatrix(novoTamanho)
		b12 := NovaMatrix(novoTamanho)
		b21 := NovaMatrix(novoTamanho)
		b22 := NovaMatrix(novoTamanho)

		aResult := NovaMatrix(novoTamanho)
		bResult := NovaMatrix(novoTamanho)

		for i := 0; i < novoTamanho; i++ {
			for j := 0; j < novoTamanho; j++ {

				a11[i][j] = matrix1[i][j]
				a12[i][j] = matrix1[i][j+novoTamanho]
				a21[i][j] = matrix1[i+novoTamanho][j]
				a22[i][j] = matrix1[i+novoTamanho][j+novoTamanho]

				b11[i][j] = matrix2[i][j]
				b12[i][j] = matrix2[i][j+novoTamanho]
				b21[i][j] = matrix2[i+novoTamanho][j]
				b22[i][j] = matrix2[i+novoTamanho][j+novoTamanho]

			}
		}

		aResult = add(a11, a22)
		bResult = add(b11, b22)

		p1 := Multiplica(aResult, bResult)

		aResult = add(a21, a22)
		p2 := Multiplica(aResult, b11)

		bResult = minus(b12, b22)
		p3 := Multiplica(a11, bResult)

		bResult = minus(b21, b11)

		p4 := Multiplica(a22, bResult)

		aResult = add(a11, a12)
		p5 := Multiplica(aResult, b22)

		aResult = minus(a21, a11)
		bResult = add(b11, b12)
		p6 := Multiplica(aResult, bResult)

		aResult = minus(a12, a22)
		bResult = add(b21, b22)

		p7 := Multiplica(aResult, bResult)

		c12 := add(p3, p5)
		c21 := add(p2, p4)

		aResult = add(p1, p4)
		bResult = add(aResult, p7)
		c11 := minus(bResult, p5)

		aResult = add(p1, p3)
		bResult = add(aResult, p6)
		c22 := minus(bResult, p2)

		c := NovaMatrix(tamanho)

		for i := 0; i < novoTamanho; i++ {
			for j := 0; j < novoTamanho; j++ {

				c[i][j] = c11[i][j]

				c[i][j+novoTamanho] = c12[i][j]

				c[i+novoTamanho][j] = c21[i][j]

				c[i+novoTamanho][j+novoTamanho] = c22[i][j]

			}
		}
		return c
	}

}

func add(matrix1, matrix2 [][]int) [][]int {
	tam := len(matrix1)
	sum := NovaMatrix(tam)
	for i := 0; i < tam; i++ {
		for j := 0; j < tam; j++ {
			sum[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}

	return sum

}

func minus(matrix1, matrix2 [][]int) [][]int {
	tam := len(matrix1)
	difference := NovaMatrix(tam)
	for i := 0; i < tam; i++ {
		for j := 0; j < tam; j++ {
			difference[i][j] = matrix1[i][j] - matrix2[i][j]
		}
	}

	return difference

}
