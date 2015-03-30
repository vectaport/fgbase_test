package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbi(x flowgraph.Edge) {

	node:=flowgraph.NewNode("tbi", nil, []*flowgraph.Edge{&x}, nil)
	
	x.Val = 0

	for {

		if node.Rdy() {
			node.TraceVals()
			node.Tracef("writing x.Data: %d\n", x.Val.(int))
			x.Data <- x.Val
			x.Val = (x.Val.(int) + 1)%2
			x.Rdy = false
		}

		node.Select()

	}
	
}

func tbo(a flowgraph.Edge) {
	node:=flowgraph.NewNode("tbo", []*flowgraph.Edge{&a}, nil, nil)
	
	for {
		if node.Rdy() {
			node.Tracef("writing a.Ack\n")
			node.TraceVals()
			a.Ack <- true
			a.Rdy = false
		}

		node.Select()

	}

}

func main() {

	a := flowgraph.NewEdge("a",nil)
	x := flowgraph.NewEdge("x",nil)
	y := flowgraph.NewEdge("y",nil)

	go tbi(a)
	go flowgraph.FuncStrCnd(a, x, y)
	go tbo(x)
	go tbo(y)

	time.Sleep(1000000000)

}

