package main

import (
	"time"

	"github.com/vectaport/flowgraph"
)

func tbiWork(n *flowgraph.Node) {
	x := n.Dsts[0]
	x.Val = make([]complex128, 32, 32)
}

func tbi(x flowgraph.Edge) flowgraph.Node {
	node:=flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, tbiWork)
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {
	node:=flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	flowgraph.TraceLevel = flowgraph.V

	e,n := flowgraph.MakeGraph(3,3)

	e[1].Const(false)

	n[0] = tbi(e[0])
	n[1] = flowgraph.FuncFFT(e[0], e[1], e[2])
	n[2] = tbo(e[2])

	flowgraph.RunAll(n, time.Second)

}

