package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbiFire(n *flowgraph.Node) {
	x := n.Dsts[0]
	x.Val = x.Aux
	x.Aux = (x.Aux.(int) + 1)%2
}

func tbi(x flowgraph.Edge) flowgraph.Node {

	node:=flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, tbiFire)
	x.Aux = 0
	return node
	
}

func tbo(a flowgraph.Edge) flowgraph.Node {
	node:=flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	flowgraph.TraceLevel = flowgraph.V

	e,n := flowgraph.MakeGraph(3,4)

	n[0] = tbi(e[0])
	n[1] = flowgraph.FuncSteerc(e[0], e[1], e[2])
	n[2] = tbo(e[1])
	n[3] = tbo(e[2])

	flowgraph.RunAll(n, time.Second)

}

