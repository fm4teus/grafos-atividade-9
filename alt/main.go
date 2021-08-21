package main

import (
	"fmt"

	"github.com/yourbasic/graph"
)

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

func geraGrafoDeBruijn(palavras []string, r int) *graph.Mutable {
	g := graph.New(len(palavras))

	for indiceA, palavraA := range palavras {
		for indiceB, palavraB := range palavras {
			// se retirando primeiro simbolo de A e último de B temos mesma sequencia então cria aresta A --> B
			if palavraA[1:] == palavraB[:r-2] {
				g.Add(indiceA, indiceB)
			}
		}
	}

	return g
}

func imprime(caminhoEuleriano []int, palavras []string) {
	var deBrujin string
	for _, indiceVertice := range caminhoEuleriano {
		deBrujin = fmt.Sprintf("%s%s", deBrujin, palavras[indiceVertice][:1])
	}
	fmt.Println("sequencia: ", deBrujin[:len(deBrujin)-1])
	fmt.Println("    ----    ")
}

type teste struct {
	r int
	s int
}

func main() {

	pares := []teste{{r: 2, s: 2}, {r: 3, s: 2}, {r: 4, s: 2}, {r: 5, s: 2}, {r: 2, s: 3}, {r: 2, s: 4}, {r: 3, s: 3}, {r: 4, s: 3}, {r: 3, s: 4}, {r: 6, s: 2}, {r: 3, s: 7}}

	for _, par := range pares {
		fmt.Println("Tamanho da palavra: ", par.r)

		fmt.Println("Número de símbolos: ", par.s)

		palavras := geraPalavras(par.r-1, par.s)

		g := geraGrafoDeBruijn(palavras, par.r)

		caminhoEuleriano, _ := graph.EulerDirected(g)

		imprime(caminhoEuleriano, palavras)

	}

}
