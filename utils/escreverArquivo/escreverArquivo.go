package escreverArquivo

import (
	"bufio"
	"fmt"
	"os"
)

func EscreverTexto(linhas string, caminhoDoArquivo string) error {

	arquivo, err := os.Create("apresentacao/" + caminhoDoArquivo)
	if err != nil {
		return err
	}
	defer arquivo.Close()
	escritor := bufio.NewWriter(arquivo)

	fmt.Fprintln(escritor, linhas)

	return escritor.Flush()
}
