package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type G struct {
	// collection if vertices and edges
	V []int   // list (slice) of  vertices
	E [][]int // adjacency list
}

func NewGraph(vertices int) *G {
	g := new(G)
	g.V = make([]int, vertices, vertices)

	for i := 0; i < vertices; i++ {
		g.V[i] = i
	}

	g.E = make([][]int, vertices)
	for i := range g.E {
		g.E[i] = make([]int, 0)
	}

	return g
}

func (g *G) Print() {
	for v := range g.V {
		fmt.Printf("%4d -> ", v)
		for i := range g.E[v] {
			fmt.Printf("%2d ", g.E[v][i])
		}
		fmt.Println()
	}
}

func (g *G) AddEdge(u, v int) {
	// add an edge from u to v
	g.E[u] = append(g.E[u], v)
}

func (g *G) DFS(s int) {
	// start a DFS from source s
	//TODO
}

func (g *G) BFS(s int) {
	// start a breadth first search from source vertex s
	// TODO
}

func (g *G) hasLoop() bool {
	// TODO
	return false
}

func main() {
	file, err := os.Open("edges.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sep := ","

	g := NewGraph(6)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		edge := strings.Split(strings.TrimSpace(line), sep)
		// fmt.Println(edge)
		if len(edge) != 2 {
			fmt.Println("Error in input")
			return
		}

		u, _ := strconv.Atoi(strings.TrimSpace(edge[0]))
		v, _ := strconv.Atoi(strings.TrimSpace(edge[1]))
		// fmt.Println(u, v)

		g.AddEdge(u, v)
		//fmt.Println(g.E)
	}
	g.Print()
}
