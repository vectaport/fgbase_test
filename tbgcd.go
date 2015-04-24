package main

import (
	"github.com/vectaport/flowgraph"
	"math/rand"
	"time"
)

func tbm(x flowgraph.Edge) {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil,
		func(n *flowgraph.Node) { n.Dsts[0].Val = rand.Intn(15)+1 })
	node.Run()
}

func tbn(x flowgraph.Edge) {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil,
		func(n *flowgraph.Node) { n.Dsts[0].Val = rand.Intn(15)+1 })
	node.Run()
}

func main() {

	flowgraph.TraceLevel = flowgraph.V
	flowgraph.TraceIndent = false

	e0 := flowgraph.MakeEdge("e0", nil)
	e1 := flowgraph.MakeEdge("e1", nil)
	e2 := flowgraph.MakeEdge("e2", nil)
	e3 := flowgraph.MakeEdge("e3", nil)
	e4 := flowgraph.MakeEdge("e4", nil)
	e5 := flowgraph.MakeEdge("e5", nil)
	e6 := flowgraph.MakeEdge("e6", nil)
	e7 := flowgraph.MakeEdge("e7", 0)
	e8 := flowgraph.MakeEdge("e8", nil)
	e9 := flowgraph.MakeEdge("e9", nil)
	e10 := flowgraph.MakeEdge("e10", nil)

	go tbm(e0)
	go tbn(e1)

	go flowgraph.FuncRdy(e0, e7, e2)
	go flowgraph.FuncRdy(e1, e7, e3)

	go flowgraph.FuncEither(e2, e10, e4)
	go flowgraph.FuncEither(e3, e8, e5)

	go flowgraph.FuncMod(e4, e5, e6)

	go flowgraph.FuncSteerc(e6, e7, e8)
	go flowgraph.FuncSteerv(e6, e5, e9, e10)

	go tbo(e9)

	time.Sleep(time.Second)
	flowgraph.StdoutLog.Printf("\n")

}
