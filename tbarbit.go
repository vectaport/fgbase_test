package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbi(x flowgraph.Edge) flowgraph.Node {


	node:=flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func(n *flowgraph.Node) {
			x.Val = x.Aux
			x.Aux = (x.Aux.(int) + 1)
		})

	x.Aux = 1
	return node
	
}

func tbo(a flowgraph.Edge) flowgraph.Node {
	
	node:=flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node

}

func main() {

	flowgraph.TraceLevel = flowgraph.V

	e,n := flowgraph.MakeGraph(3,4)
	e[0].Val = 0
	e[1].Val = 1000

	n[0] = tbi(e[0])
	n[1] = tbi(e[1])
	n[2] = flowgraph.FuncArbit(e[0], e[1], e[2])
	n[3] = tbo(e[2])

	flowgraph.RunAll(n[:], time.Second)

}

