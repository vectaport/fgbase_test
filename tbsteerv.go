package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbiFire(n *flowgraph.Node) {
	x := n.Dsts[0]
	x.Val = x.Aux
	if (x.Aux.(int)<=1) {
		x.Aux = (x.Aux.(int) + 1)%2
	} else {
		x.Aux = x.Aux.(int) + 1
	}
}

func tbi(x flowgraph.Edge) flowgraph.Node {
	node:=flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, tbiFire)
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {
	node:=flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	flowgraph.TraceLevel = flowgraph.V

	e,n := flowgraph.MakeGraph(4,5)

	// initialize different state in the two source testbenches (tbi)
	e[0].Aux = 0
	e[1].Aux = 1000

	n[0] = tbi(e[0])
	n[1] = tbi(e[1])
	n[2] = flowgraph.FuncSteerv(e[0], e[1], e[2], e[3])
	n[3] = tbo(e[2])
	n[4] = tbo(e[3])

	flowgraph.RunAll(n, time.Second)

}

