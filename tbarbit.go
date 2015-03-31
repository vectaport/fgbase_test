package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbi(x flowgraph.Edge) {


	node:=flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil)

	for {


		if node.Rdy() {
			node.TraceVals()
			node.Tracef("writing x.Data: %d\n", x.Val.(int))
			x.Rdy = false
			x.Data <- x.Val
			x.Val = (x.Val.(int) + 1)
		}

		node.Select()

	}
	
}

func tbo(a flowgraph.Edge) {
	
	node:=flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil)
	
	for {
		if node.Rdy() {
			node.TraceVals()
			node.Tracef("writing a.Ack\n")
			a.Ack <- true
			a.Rdy = false
		}

		node.Select()

	}

}

func main() {

	flowgraph.Debug = false
	flowgraph.Indent = false

	a := flowgraph.MakeEdge("a",nil)
	a.Val = 0
	b := flowgraph.MakeEdge("b",nil)
	b.Val = 1000
	x := flowgraph.MakeEdge("x",nil)

	go tbi(a)
	go tbi(b)
	go flowgraph.FuncArbit(a, b, x)
	go tbo(x)

	time.Sleep(1000000000)

}

