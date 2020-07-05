package iterativo

func Mergesort(vetor []int) []int {

	if len(vetor) == 1 {
		return vetor
	}

	l := Mergesort(vetor[:(len(vetor) / 2)])
	r := Mergesort(vetor[(len(vetor) / 2):])

	return Merge(l, r)
}

func Merge(esquerda []int, direita []int) []int {

	var saida []int
	var a = 0
	var b = 0

	for a < len(esquerda) && b < len(direita) {

		if esquerda[a] < direita[b] {
			saida = append(saida, esquerda[a])
			a++
		} else {
			saida = append(saida, direita[b])
			b++
		}

	}

	if a < len(esquerda) || b < len(direita) {
		saida = append(saida, esquerda[a:]...)
		saida = append(saida, direita[b:]...)
	}
	return saida
}
