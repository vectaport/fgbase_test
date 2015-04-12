package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbiFire(n *flowgraph.Node) {
	x := n.Dsts[0]
	x.Val = make([]complex128, 32, 32)
}

func tbi(x flowgraph.Edge) {
	node:=flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, tbiFire)
	node.Run()
}

func tbo(a flowgraph.Edge) {
	node:=flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	node.Run()
}

func main() {

	flowgraph.TraceLevel = flowgraph.V
	flowgraph.TraceIndent = false

	e0 := flowgraph.MakeEdge("e0",nil)
	e1 := flowgraph.MakeEdge("e1",nil)
	cfalse := flowgraph.MakeEdgeConst("cfalse", false)

	go tbi(e0)
	go flowgraph.FuncFft(e0, cfalse, e1)
	go tbo(e1)

	time.Sleep(time.Second)
	flowgraph.StdoutLog.Printf("\n")

}

