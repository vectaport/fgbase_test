package main

import (
	"math/rand"

	"github.com/vectaport/flowgraph"
)

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil,
		func(n *flowgraph.Node) { n.Dsts[0].Val = rand.Intn(7)+1 })
	return node
}

func main() {

	flowgraph.ConfigByFlag(nil)

	e,n := flowgraph.MakeGraph(7,5)

	e[3].Const(1)
	e[5].Val = 0

	n[0] = tbi(e[0])
	n[1] = flowgraph.FuncRdy(e[0], e[5], e[1])
	n[2] = flowgraph.FuncEither(e[1], e[6], e[2])
	n[3] = flowgraph.FuncSub(e[2], e[3], e[4])
	n[4] = flowgraph.FuncSteerc(e[4], e[5], e[6])

	flowgraph.RunAll(n)

}
