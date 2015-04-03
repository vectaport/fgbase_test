package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbi_func(n *flowgraph.Node) {
	x := n.Dsts[0]
	x.Val = x.Aux
	if (x.Aux.(int)<=1) {
		x.Aux = (x.Aux.(int) + 1)%2
	} else {
		x.Aux = x.Aux.(int) + 1
	}
}

func tbi(x flowgraph.Edge) {
	node:=flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, tbi_func)
	node.Run()
}

func tbo(a flowgraph.Edge) {
	node:=flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	node.Run()
}

func main() {

	flowgraph.Debug = false
	flowgraph.Indent = false

	e0 := flowgraph.MakeEdge("e0",nil)
	e1 := flowgraph.MakeEdge("e1",nil)
	e2 := flowgraph.MakeEdge("e2",nil)
	e3 := flowgraph.MakeEdge("e3",nil)

	e0.Aux = 0
	e1.Aux = 1000

	go tbi(e0)
	go tbi(e1)
	go flowgraph.FuncSteerv(e0, e1, e2, e3)
	go tbo(e2)
	go tbo(e3)

	time.Sleep(1000000000)

}

