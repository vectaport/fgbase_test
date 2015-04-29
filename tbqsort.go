package main

import (
	"math/rand"
	"sort"
	"time"

	"github.com/vectaport/flowgraph"
)

type bushel []int

// borrowed from Golang 1.4.2 sort example, copyright notice in flowgraph/GO-LICENSE
func (a bushel) Len() int           { return len(a) }
func (a bushel) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bushel) Less(i, j int) bool { return a[i] < a[j] }


func tbiRand() sort.Interface {
	var s bushel
	for i:=0; i<1024; i++ {
		s = append(s, rand.Intn(1024))
	}
	return s
}

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil,
		func(n *flowgraph.Node) { n.Dsts[0].Val = tbiRand() })
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	flowgraph.TraceLevel = flowgraph.V

	e,n := flowgraph.MakeGraph(2, 3)

	n[0] = tbi(e[0])

	n[1] = flowgraph.FuncQsort(e[0], e[1])
	n[2] = tbo(e[1])

	flowgraph.RunAll(n, time.Second)

}
