package main

import (
	"math/rand"

	"github.com/vectaport/fgbase"
	"github.com/vectaport/fgbase/imglab"
)

func tbiFire(n *flowgraph.Node) {
	x := n.Dsts[0]
	l := 1024
	buf := make([]complex128, l, l)
	for i := 0; i<l; i++ {
		buf[i] = complex(rand.Float64(), rand.Float64())
	}
	x.DstPut(buf)
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

	e,n := flowgraph.MakeGraph(3,3)

	e[1].Const(false)

	n[0] = tbi(e[0])
	n[1] = imglab.FuncFFT(e[0], e[1], e[2])
	n[2] = tbo(e[2])

	flowgraph.RunAll(n)

}

