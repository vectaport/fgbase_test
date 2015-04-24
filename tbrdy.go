package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbi(x flowgraph.Edge) {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			x.Val = x.Aux
			x.Aux = x.Aux.(int) + 1
		})
	node.Run()
}

func tbo(a flowgraph.Edge) {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	node.Run()
}

func main() {

	flowgraph.TraceLevel = flowgraph.V
	flowgraph.TraceIndent = false

	e0 := flowgraph.MakeEdge("e0",nil)
	e0.Aux = 0
	e1 := flowgraph.MakeEdge("e1",nil)
	e1.Aux = 1000
	e2 := flowgraph.MakeEdge("e2",nil)

	go tbi(e0)
	go tbi(e1)
	go flowgraph.FuncRdy(e0, e1, e2)
	go tbo(e2)

	time.Sleep(time.Second)
	flowgraph.StdoutLog.Printf("\n")

}

