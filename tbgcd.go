package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/vectaport/flowgraph"
)

func tbm(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil,
		func(n *flowgraph.Node) { n.Dsts[0].Val = rand.Intn(15)+1 })
	return node
}

func tbn(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil,
		func(n *flowgraph.Node) { n.Dsts[0].Val = rand.Intn(15)+1 })
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	tracep := flag.String("trace", "V", "trace level, Q|V|VV|VVV|VVVV")
	flag.Parse()
	
	flowgraph.TraceLevel = flowgraph.TraceLevels[*tracep]

	e,n := flowgraph.MakeGraph(11, 10)

	e[7].Val = 0

	n[0] = tbm(e[0])
	n[1] = tbn(e[1])

	n[2] = flowgraph.FuncRdy(e[0], e[7], e[2])
	n[3] = flowgraph.FuncRdy(e[1], e[7], e[3])

	n[4] = flowgraph.FuncEither(e[2], e[10], e[4])
	n[5] = flowgraph.FuncEither(e[3], e[8], e[5])

	n[6] = flowgraph.FuncMod(e[4], e[5], e[6])

	n[7] = flowgraph.FuncSteerc(e[6], e[7], e[8])
	n[8] = flowgraph.FuncSteerv(e[6], e[5], e[9], e[10])

	n[9] = tbo(e[9])

	flowgraph.RunAll(n, time.Second)

}
