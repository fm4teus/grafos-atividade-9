package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"./grafo"
)

func leitura() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	rawInput, err := reader.ReadString('\n')
	if err != nil {
		return -1, err
	}
	rawInput = rawInput[:len(rawInput)-1]

	input, err := strconv.Atoi(rawInput)
	if err != nil {
		return -1, err
	}

	return input, nil
}

func codifica(n int) string {
	return string(rune('a' + n))
}

func geraPalavras(res []string, r, s int) []string {
	if r > 1 {
		palavras := geraPalavras(res, r-1, s)
		var aux []string
		copy(aux, palavras)
		for _, palavra := range palavras {
			for j := 0; j < s; j++ {
				aux = append(aux, fmt.Sprintf("%s%s", codifica(j), palavra))
			}
		}
		return aux

	} else if r == 1 {
		for i := 0; i < s; i++ {
			res[i] = codifica(i)
		}
		return res
	}

	return nil
}

func main() {

	fmt.Println("Insira o tamanho da palavra")
	r, err := leitura()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Insira o número de símbolos")
	s, err := leitura()
	if err != nil {
		fmt.Println(err)
	}

	size := int(math.Pow(float64(s), float64(r)))
	res := make([]string, size)
	palavras := geraPalavras(res, r, s)[:size]

	fmt.Println("res", len(palavras))
	fmt.Println("len", len(palavras))

	vertices := []grafo.Vertice{}
	for _, palavra := range palavras {
		vertices = append(vertices, grafo.Vertice{Label: palavra})
	}

	graf := grafo.Grafo{
		Vertices: vertices,
	}
	graf.ImprimeVertices()

}
