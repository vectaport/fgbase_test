package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbi(x flowgraph.Edge) {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, nil)

	x.Aux = 0

	var i int = 0
	for {
		if (i>1000000) { break }

		if node.RdyAll(){
			x.Val = x.Aux
			x.Aux = x.Aux.(int) + 1
			node.SendAll()
			i = i + 1
		}

		node.RecvOne()

	}
}

func tbo(a flowgraph.Edge) {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	node.Run()

}

func main() {

	flowgraph.Debug = false
	flowgraph.Indent = false

	e0 := flowgraph.MakeEdge("e0",nil)
	e1 := flowgraph.MakeEdge("e1",nil)
	e2 := flowgraph.MakeEdge("e2",nil)

	go tbi(e0)
	go flowgraph.FuncConst(e1, 1000)
	go flowgraph.FuncAdd(e0, e1, e2)
	go tbo(e2)

	time.Sleep(1000000000)
	flowgraph.StdoutLog.Printf("\n")

}

