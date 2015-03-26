package main

import (
	"github.com/vectaport/flowgraph"
	"reflect"
	"time"
)

func tbi(x flowgraph.Edge) {

	node:=flowgraph.NewNode("tbi", nil, []*flowgraph.Edge{&x}, nil)
	
	for {

		if node.Rdy() {
			node.TraceVals()
			node.Tracef("writing x.Data: %v\n", x.Val)
			x.Data <- x.Val
			if (x.Val.(int)<=1) {
				x.Val = (x.Val.(int) + 1)%2
			} else {
				x.Val = x.Val.(int) + 1
			}
			x.Rdy = false
		}

		node.Tracef("select\n")
		select {
		case x.Rdy = <-x.Ack:
			node.Tracef("x.Ack read\n")
			
			
		}
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

		node.Tracef("select\n")
		select {
		case a.Val = <-a.Data:
			{
				node.Tracef("a read %v --  %v\n", reflect.TypeOf(a.Val), a.Val)
				a.Rdy = true
			}
		}

	}

}

func main() {

	flowgraph.Debug = false
	flowgraph.Indent = false

	a := flowgraph.NewEdge("a",nil)
	b := flowgraph.NewEdge("b",nil)
	x := flowgraph.NewEdge("x",nil)
	y := flowgraph.NewEdge("y",nil)

	a.Val = 0
	go tbi(a)
	b.Val = 1000
	go tbi(b)
	go flowgraph.FuncStrVal(a, b, x, y)
	go tbo(x)
	go tbo(y)

	time.Sleep(1000000000)

}

