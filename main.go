package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/yourbasic/graph"
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
	return string(rune('0' + n))
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

func main() {

	fmt.Print("Insira o tamanho da palavra: ")
	r, err := leitura()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Insira o número de símbolos: ")
	s, err := leitura()
	if err != nil {
		fmt.Println(err)
	}

	palavras := geraPalavras(r-1, s)

	g := graph.New(len(palavras))

	for indiceA, palavraA := range palavras {
		for indiceB, palavraB := range palavras {
			// se retirando primeiro simbolo de A e último de B temos mesma sequencia então cria aresta A --> B
			if palavraA[1:] == palavraB[:r-2] {
				g.Add(indiceA, indiceB)
			}
		}
	}

	caminhoEuleriano, _ := graph.EulerDirected(g)

	var deBrujin string
	for _, indiceVertice := range caminhoEuleriano {
		deBrujin = fmt.Sprintf("%s%s", deBrujin, palavras[indiceVertice][:1])
	}

	fmt.Println("sequencia: ", deBrujin[:len(deBrujin)-1])

}
