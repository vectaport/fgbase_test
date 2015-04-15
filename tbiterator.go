package main

import (
	"github.com/vectaport/flowgraph"
	"math/rand"
	"time"
)

func tbi(x flowgraph.Edge) {

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
	e3 := flowgraph.MakeEdgeConst("e3", 1)
	e4 := flowgraph.MakeEdge("e4", nil)
	e5 := flowgraph.MakeEdge("e5", int(0))
	e6 := flowgraph.MakeEdge("e6", nil)

	go tbi(e0)
	go flowgraph.FuncRdy(e0, e5, e1)
	go flowgraph.FuncArbit(e1, e6, e2)
	go flowgraph.FuncSub(e2, e3, e4)
	go flowgraph.FuncSteerc(e4, e5, e6)

	time.Sleep(time.Second)
	flowgraph.StdoutLog.Printf("\n")

}
