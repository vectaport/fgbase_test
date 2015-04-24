package main

import (
	"time"

	"github.com/vectaport/flowgraph"
)

func tbiFire(n *flowgraph.Node) {
	x := n.Dsts[0]
	y := n.Dsts[1]
	x.Val = x.Aux
	y.Val = y.Aux
	x.Aux = x.Aux.(int) + 1
	y.Aux = y.Aux.(int) + 1
}

func tbi(a, x, y flowgraph.Edge) flowgraph.Node {
	node := flowgraph.MakeNode("tbi", []*flowgraph.Edge{&a}, []*flowgraph.Edge{&x, &y}, nil, tbiFire)
	x.Aux = 1
	y.Aux = 1
	return node
}

func tboFire(n *flowgraph.Node) {
	x := n.Dsts[0]
	x.Val = true
}

func tbo(a, x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, []*flowgraph.Edge{&x}, nil, tboFire)
	return node

}

func main() {

	flowgraph.TraceLevel = flowgraph.V

	e,n := flowgraph.MakeGraph(4,3)

	e[3].Val = true // initialize data wavefront

	n[0] = tbi(e[3], e[0], e[1])
	n[1] = flowgraph.FuncAdd(e[0], e[1], e[2])
	n[2] = tbo(e[2], e[3])

	flowgraph.RunAll(n, time.Second)

}

