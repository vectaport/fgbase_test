package main

import (
	"github.com/vectaport/flowgraph"
	"reflect"
	"time"
)

func tbi(x flowgraph.Edge) {


	node:=flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x})

	x.Val = 0

	for {


		if x.Rdy {
			node.ExecCnt()
			node.PrintVals()
			node.Printf("writing x.Data: %d\n", x.Val.(int))
			x.Rdy = false
			x.Data <- x.Val
			x.Val = (x.Val.(int) + 1)
		}

		node.Printf("select\n")
		select {
		case x.Rdy = <-x.Ack:
			node.Printf("x.Ack read\n")
			
			
		}
	}
	
}

func tbo(a flowgraph.Edge) {
	
	node:=flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil)
	
	for {
		if a.Rdy {
			node.ExecCnt()
			node.PrintVals()
			node.Printf("writing a.Ack\n")
			a.Ack <- true
			a.Rdy = false
		}

		node.Printf("select\n")
		select {
		case a.Val = <-a.Data:
			{
				node.Printf("a read %v --  %v\n", reflect.TypeOf(a.Val), a.Val)
				a.Rdy = true
			}
		}

	}

}

func main() {

	flowgraph.Debug = false
	flowgraph.Indent = false

	a := flowgraph.MakeEdge2("a", false,true,nil)
	b := flowgraph.MakeEdge2("b", false,true,nil)
	x := flowgraph.MakeEdge2("x", false,true,nil)

	go tbi(a)
	go tbi(b)
	go flowgraph.ArbitNode(a, b, x)
	go tbo(x)

	time.Sleep(1000000000)

}

