package main

import (
	"github.com/vectaport/fgbase"
)

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			x.DstPut(n.Aux)
			n.Aux = n.Aux.(int) + 1
		})
	node.Aux = 0
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	flowgraph.ConfigByFlag(nil)
	
	e,n := flowgraph.MakeGraph(2,4)

	n[0] = tbi(e[0])
	n[1] = flowgraph.FuncPass(e[0], e[1])
	n[2] = tbo(e[1])
	n[3] = tbo(e[1])

	flowgraph.RunAll(n)
}

