package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbi(x flowgraph.Edge) {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil)

	x.Val = 10

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

/*
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
*/

func main() {

	flowgraph.Debug = false
	flowgraph.Indent = false

	start := flowgraph.MakeEdge("start",nil)
	rdy0 := flowgraph.MakeEdge("rdy0",nil)
	arbit0 := flowgraph.MakeEdge("arbit0",nil)
	const1 := flowgraph.MakeEdge("const1",nil)
	const1.Val = 1
	sub0 := flowgraph.MakeEdge("sub0",nil)
	strcnd0 := flowgraph.MakeEdge("strcnd0",int(0))
	strcnd1 := flowgraph.MakeEdge("strcnd1",nil)

	go tbi(start)
	go flowgraph.RdyNode(start, strcnd0, rdy0)
        go flowgraph.ArbitNode(rdy0, strcnd1, arbit0)
	go flowgraph.SubNode(arbit0, const1, sub0)
	go flowgraph.StrCndNode(sub0, strcnd0, strcnd1)
	go flowgraph.ConstNode(const1)

	time.Sleep(1000000000)

}

