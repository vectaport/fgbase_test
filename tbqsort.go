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

func (a bushel) SubSlice(n, m int) flowgraph.Datum {return a[n:m]}

func tbiRand() sort.Interface {
	var s bushel
	n := 1024*1024
	l := rand.Intn(n)
	for i:=0; i<l; i++ {
		s = append(s, rand.Intn(n))
	}
	return s
}

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil,
		func(n *flowgraph.Node) { n.Dsts[0].Val = tbiRand() })
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, 
		func(n *flowgraph.Node) {
			switch v := a.Val.(type) {
			case flowgraph.Interface2: {
				n.Tracef("Fragment sorted? %v\n", v.Sorted())
			}
			default: {
				n.Tracef("not of type flowgraph.Interface2\n")
			}
			}})
	return node
}

func main() {

	flowgraph.TraceLevel = flowgraph.V

	const poolSz = 128
	e,n := flowgraph.MakeGraph(2, poolSz+2)

	n[0] = tbi(e[0])

	p := n[1:poolSz+1]
	copy(p, flowgraph.FuncQsort(e[0], e[1], poolSz))
	n[poolSz+1] = tbo(e[1])

	flowgraph.RunAll(n, 1*time.Second)

}
