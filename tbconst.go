package main

import (
	"github.com/vectaport/flowgraph"
	"reflect"
	"time"
)

func tbi(x flowgraph.Edge) {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil)

	var i int = 0
	for {
		if (i>10) { break }

		if node.Rdy(){
			node.Printf("writing x.Data: %d\n", x.Val.(int))
			node.PrintVals()
			x.Data <- x.Val
			x.Rdy = false
			x.Val = x.Val.(int) + 1
			i = i + 1
		}

		node.Printf("select\n")
		select {
		case x.Rdy = <-x.Ack:
			node.Printf("x.Ack read\n")

		}

	}
}

func tbo(a flowgraph.Edge) {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil)

	for {
		if node.Rdy() {
			node.Printf("writing a.Ack\n")
			node.PrintVals()
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

	a := flowgraph.MakeEdge("a", false,true,nil)
	b := flowgraph.MakeEdge("b", false,true,nil)
	x := flowgraph.MakeEdge("x", false,true,nil)

	a.Val = 0
	go tbi(a)
	b.Val = 1000
	go flowgraph.ConstNode(b)
	go flowgraph.AddNode(a, b, x)
	go tbo(x)

	time.Sleep(1000000000)

}

