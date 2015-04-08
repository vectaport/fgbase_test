package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbiFire(n *flowgraph.Node) {
	x := n.Dsts[0]
	y := n.Dsts[1]
	x.Val = x.Aux
	y.Val = y.Aux
	x.Aux = x.Aux.(int) + 1
	y.Aux = y.Aux.(int) + 1
}

func tbi(a, x, y flowgraph.Edge) {
	node := flowgraph.MakeNode("tbi", []*flowgraph.Edge{&a}, []*flowgraph.Edge{&x, &y}, nil, tbiFire)
	x.Aux = 1
	y.Aux = 1
	node.Run()
}

func tboFire(n *flowgraph.Node) {
	x := n.Dsts[0]
	x.Val = true
}

func tbo(a, x flowgraph.Edge) {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, []*flowgraph.Edge{&x}, nil, tboFire)
	node.Run()

}

func main() {

	flowgraph.TraceLevel = flowgraph.V
	flowgraph.TraceIndent = false

	e0 := flowgraph.MakeEdge("e0",nil)
	e1 := flowgraph.MakeEdge("e1",nil)
	e2 := flowgraph.MakeEdge("e2",nil)
	e3 := flowgraph.MakeEdge("e3",true)

	go tbi(e3, e0, e1)
	go flowgraph.FuncAdd(e0, e1, e2)
	go tbo(e2, e3)

	time.Sleep(1000000000)
	flowgraph.StdoutLog.Printf("\n")

}

