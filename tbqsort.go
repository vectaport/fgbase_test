package main

import (
	"math/rand"
	"time"

	"github.com/vectaport/flowgraph"
)

type bushel []int

// borrowed from Golang 1.4.2 sort example, copyright notice in flowgraph/GO-LICENSE
func (a bushel) Len() int           { return len(a) }
func (a bushel) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bushel) Less(i, j int) bool { return a[i] < a[j] }

func (a bushel) Sorted() bool {
	l := len(a)
	for i:= 0; i<l-1; i++ {
		if a[i] > a[i+1] {
			// flowgraph.StdoutLog.Printf("a[%d] > a[%d]:  %d > %d\n", i, i+1, a[i], a[i+1])
			return false
		}
	}
	return true
}

func tbiRand() flowgraph.Interface {
	var s bushel
	l := rand.Intn(1024*1024)
	for i:=0; i<l; i++ {
		s = append(s, rand.Intn(8))
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

	e,n := flowgraph.MakeGraph(2, 10)

	n[0] = tbi(e[0])

	p := n[1:9]
	copy(p, flowgraph.FuncQsort(e[0], e[1], 8))
	n[9] = tbo(e[1])

	flowgraph.RunAll(n, time.Second)

}
