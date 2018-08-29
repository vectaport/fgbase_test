package main

import (
	"github.com/vectaport/flowgraphbase"
)

func tbiFire(n *flowgraph.Node) {
	a := n.Srcs[0]
	x := n.Dsts[0]
	y := n.Dsts[1]
	a.Flow = true
	x.DstPut(n.Aux)
	y.DstPut(n.Aux)
	n.Aux = n.Aux.(int) + 1
}

func tbi(a, x, y flowgraph.Edge) flowgraph.Node {
	node := flowgraph.MakeNode("tbi", []*flowgraph.Edge{&a}, []*flowgraph.Edge{&x, &y}, nil, tbiFire)
	node.Aux = 1
	return node
}

func tboFire(n *flowgraph.Node) {
	a := n.Srcs[0]
	x := n.Dsts[0]
	a.Flow = true
	x.DstPut(true)
}

func tbo(a, x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, []*flowgraph.Edge{&x}, nil, tboFire)
	return node

}

func main() {

	flowgraph.ConfigByFlag(nil)

	e,n := flowgraph.MakeGraph(4,3)

	e[3].Val = true // initialize data wavefront

	n[0] = tbi(e[3], e[0], e[1])
	n[1] = flowgraph.FuncAdd(e[0], e[1], e[2])
	n[2] = tbo(e[2], e[3])

	flowgraph.RunAll(n)

}

