package main

import (
	"github.com/vectaport/flowgraph"
	"reflect"
	"time"
)

func tbi(a, x, y flowgraph.Edge) {

	node := flowgraph.MakeNode("tbi", []*flowgraph.Edge{&a}, []*flowgraph.Edge{&x, &y}, nil)

	x.Val = 0
	y.Val = 0
	
	for {
		node.PrintStatus(false)
		
		if node.Rdy() {
			node.PrintVals()
			a.Ack <- true
			node.Printf("a.Ack written\n");
			x.Data <- x.Val
			node.Printf("x.Data written\n");
			y.Data <- y.Val
			node.Printf("y.Data written\n");
			x.Val = x.Val.(int) + 1
			y.Val = y.Val.(int) + 1
			x.Rdy = false
			y.Rdy = false
			a.Rdy = false
		}
		
		node.Printf("select\n")
		select {
		case x.Rdy = <-x.Ack: {
			node.Printf("x.Ack read\n")
		}
			
		case y.Rdy = <-y.Ack: {
			node.Printf("y.Ack read\n")
		}
			
		case a.Val = <-a.Data: {
			node.Printf("a.Data read\n")
			flowgraph.Sink(a.Val)
			a.Rdy = true
		}
		}
		
	}
}

func tbo(a, x flowgraph.Edge) {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, []*flowgraph.Edge{&x}, nil)

	for {
		node.PrintStatus(false)
		if node.Rdy() {
			x.Val = true
			node.PrintVals()
			node.Printf("writing x.Data and a.Ack\n")
			x.Data <- x.Val
			node.Printf("done writing x.Data\n")
			a.Ack <- true
			node.Printf("done writing a.Ack\n")
			a.Rdy = false
			x.Rdy = false
		}

		node.Printf("select\n")
		select {
		case a.Val = <-a.Data:
			{
				node.Printf("a read %v --  %v\n", reflect.TypeOf(a.Val), a.Val)
				a.Rdy = true
			}
		case x.Rdy = <-x.Ack:
			node.Printf("x.Ack read\n")
		}

	}

}

func main() {

	flowgraph.Debug = false
	flowgraph.Indent = false

	a := flowgraph.MakeEdge("a",false,true,int(0))
	b := flowgraph.MakeEdge("b",false,true,int(0))
	x := flowgraph.MakeEdge("x",false,true,nil)
	g := flowgraph.MakeEdge("g",true,false,nil)

	go tbi(g, a, b)
	go flowgraph.AddNode(a, b, x)
	go tbo(x, g)

	time.Sleep(1000000000)

}

