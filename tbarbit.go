package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbi(x flowgraph.Edge) {


	node:=flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, nil)
	
	x.Aux = 1

	for {
		

		if node.RdyAll() {
			x.Val = x.Aux
			node.SendAll()
			x.Aux = (x.Aux.(int) + 1)
		}

		node.RecvOne()

	}
	
}

func tbo(a flowgraph.Edge) {
	
	node:=flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	
	for {
		if node.RdyAll() {
			node.SendAll()
		}

		node.RecvOne()

	}

}

func main() {

	flowgraph.Debug = false
	flowgraph.Indent = false

	e0 := flowgraph.MakeEdge("e0",0)
	e1 := flowgraph.MakeEdge("e1",1000)
	e2 := flowgraph.MakeEdge("e2",nil)

	go tbi(e0)
	go tbi(e1)
	go flowgraph.FuncArbit(e0, e1, e2)
	go tbo(e2)

	time.Sleep(1000000000)
	flowgraph.StdoutLog.Printf("\n")

}

