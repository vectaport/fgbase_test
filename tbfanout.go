package main

import (
	"time"

	"github.com/vectaport/flowgraph"
)

func tbi(x flowgraph.Edge) flowgraph.Node {

	x.Aux = 0
	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			x.Val = x.Aux
			x.Aux = x.Aux.(int) + 1
		})
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	flowgraph.TraceLevel = flowgraph.V

	e,n := flowgraph.MakeGraph(2,4)

	n[0] = tbi(e[0])
	n[1] = flowgraph.FuncPass(e[0], e[1])
	n[2] = tbo(e[1])
	n[3] = tbo(e[1])

	flowgraph.RunAll(n, time.Second)
}

