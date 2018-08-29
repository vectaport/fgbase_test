package main

import (
	"github.com/vectaport/flowgraphbase"
)

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			x.DstPut(n.Aux)
			n.Aux = n.Aux.(int) + 1
		})
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	flowgraph.ConfigByFlag(nil)

	e,n := flowgraph.MakeGraph(3,4)
 
	n[0] = tbi(e[0])
	n[1] = tbi(e[1])
	n[2] = flowgraph.FuncRdy(e[0], e[1], e[2])
	n[3] = tbo(e[2])

	n[0].Aux = 0
	n[1].Aux = 1000

	flowgraph.RunAll(n)

}

