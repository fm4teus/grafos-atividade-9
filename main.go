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

func geraPalavras(r, s int) []string {
	if r > 1 {
		palavras := geraPalavras(r-1, s)
		res := []string{}
		for _, palavra := range palavras {
			for j := 0; j < s; j++ {
				res = append(res, fmt.Sprintf("%s%s", codifica(j), palavra))
			}
		}
		return res

	} else if r == 1 {
		res := []string{}
		for i := 0; i < s; i++ {
			res = append(res, codifica(i))
		}
		return res
	}

	return nil
}

func geraVertices(palavras []string, size int) []grafo.Vertice {
	vertices := make([]grafo.Vertice, 0, size)
	for _, palavra := range palavras {
		vertices = append(vertices, grafo.Vertice{Label: palavra})
	}
	return vertices
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

	palavras := geraPalavras(r, s)

	size := int(math.Pow(float64(r), float64(s)))
	vertices := geraVertices(palavras, size)
	g := grafo.Grafo{
		Vertices: vertices,
	}

	for _, verticeA := range g.Vertices {
		for _, verticeB := range g.Vertices {
			// se retirando primeiro simbolo de A e último de B temos mesma sequencia então cria aresta
			if verticeA.Label[1:] == verticeB.Label[:len(verticeB.Label)-1] {
				g.Arestas = append(g.Arestas, grafo.Aresta{Origem: verticeA, Destino: verticeB})
			}
		}
	}

	g.ImprimeVertices()

}
