package grafo

import "fmt"

type Vertice struct {
	Label string
}

type Aresta struct {
	Origem  Vertice
	Destino Vertice
	Peso    int
}

type Grafo struct {
	Vertices []Vertice
	Arestas  []Aresta
}

func (g Grafo) ImprimeVertices() {
	for _, vertice := range g.Vertices {
		fmt.Println(vertice)
	}
}
