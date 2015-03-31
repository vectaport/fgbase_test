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
			node.Tracef("writing x.Data: %d\n", x.Val.(int))
			node.TraceVals()
			x.Data <- x.Val
			x.Rdy = false
			x.Val = x.Val.(int) + 1
			i = i + 1
		}

		node.Select()

	}
}

/*
func tbo(a flowgraph.Edge) {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil)

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
	go flowgraph.FuncRdy(start, strcnd0, rdy0)
        go flowgraph.FuncArbit(rdy0, strcnd1, arbit0)
	go flowgraph.FuncSub(arbit0, const1, sub0)
	go flowgraph.FuncStrCnd(sub0, strcnd0, strcnd1)
	go flowgraph.FuncConst(const1)

	time.Sleep(1000000000)

}

