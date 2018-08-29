package main

import (
	"github.com/vectaport/flowgraphbase"
)

func tbiFire(n *flowgraph.Node) {
	x := n.Dsts[0]
	x.DstPut(n.Aux)
	if (n.Aux.(int)<=1) {
		n.Aux = (n.Aux.(int) + 1)%2
	} else {
		n.Aux = n.Aux.(int) + 1
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

	flowgraph.ConfigByFlag(nil)

	e,n := flowgraph.MakeGraph(4,5)

	n[0] = tbi(e[0])
	n[1] = tbi(e[1])
	n[2] = flowgraph.FuncSteerv(e[0], e[1], e[2], e[3])
	n[3] = tbo(e[2])
	n[4] = tbo(e[3])

	// initialize different state in the two source testbenches (tbi)
	n[0].Aux = 0
	n[1].Aux = 1000

	flowgraph.RunAll(n)

}

