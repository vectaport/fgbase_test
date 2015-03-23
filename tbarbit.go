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
			node.PrintVals()
			node.Tracef("writing x.Data: %d\n", x.Val.(int))
			x.Rdy = false
			x.Data <- x.Val
			x.Val = (x.Val.(int) + 1)
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
		if a.Rdy {
			node.ExecCnt()
			node.PrintVals()
			node.Tracef("writing a.Ack\n")
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
	a.Val = 0
	b := flowgraph.NewEdge("b",nil)
	b.Val = 1000
	x := flowgraph.NewEdge("x",nil)

	go tbi(a)
	go tbi(b)
	go flowgraph.FuncArbit(a, b, x)
	go tbo(x)

	time.Sleep(1000000000)

}

