package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbi_func(n *flowgraph.Node) {
	x := n.Dsts[0]
	x.Val = x.Aux
	x.Aux = (x.Aux.(int) + 1)%2
}

func tbi(x flowgraph.Edge) {

	node:=flowgraph.MakeNode2("tbi", nil, []*flowgraph.Edge{&x}, nil, tbi_func)
	x.Aux = 0
	node.Run()
	
}

func tbo(a flowgraph.Edge) {
	node:=flowgraph.MakeNode2("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	node.Run()
}

func main() {

	flowgraph.Indent = false
	flowgraph.Debug = false

	e0 := flowgraph.MakeEdge("e0",nil)
	e1 := flowgraph.MakeEdge("e1",nil)
	e2 := flowgraph.MakeEdge("e2",nil)

	go tbi(e0)
	go flowgraph.FuncStrCnd(e0, e1, e2)
	go tbo(e1)
	go tbo(e2)

	time.Sleep(1000000000)

}

