package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbi(a, x, y flowgraph.Edge) {

	node := flowgraph.NewNode("tbi", []*flowgraph.Edge{&a}, []*flowgraph.Edge{&x, &y}, nil)

	x.Val = 0
	y.Val = 0
	
	for {
		node.TraceValRdy(false)
		
		if node.Rdy() {
			node.TraceVals()
			a.Ack <- true
			node.Tracef("a.Ack written\n");
			x.Data <- x.Val
			node.Tracef("x.Data written\n");
			y.Data <- y.Val
			node.Tracef("y.Data written\n");
			x.Val = x.Val.(int) + 1
			y.Val = y.Val.(int) + 1
			x.Rdy = false
			y.Rdy = false
			a.Rdy = false
		}
		
		node.Select()
	}
}

func tbo(a, x flowgraph.Edge) {

	node := flowgraph.NewNode("tbo", []*flowgraph.Edge{&a}, []*flowgraph.Edge{&x}, nil)

	for {
		node.TraceValRdy(false)
		if node.Rdy() {
			x.Val = true
			node.TraceVals()
			node.Tracef("writing x.Data and a.Ack\n")
			x.Data <- x.Val
			node.Tracef("done writing x.Data\n")
			a.Ack <- true
			node.Tracef("done writing a.Ack\n")
			a.Rdy = false
			x.Rdy = false
		}

		node.Select()

	}

}

func main() {

	flowgraph.Debug = false
	flowgraph.Indent = false

	a := flowgraph.NewEdge("a",nil)
	b := flowgraph.NewEdge("b",nil)
	x := flowgraph.NewEdge("x",nil)
	g := flowgraph.NewEdge("g",true)

	go tbi(g, a, b)
	go flowgraph.FuncAdd(a, b, x)
	go tbo(x, g)

	time.Sleep(1000000000)

}

