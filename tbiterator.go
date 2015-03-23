package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbi(x flowgraph.Edge) {

	node := flowgraph.NewNode("tbi", nil, []*flowgraph.Edge{&x}, nil)

	x.Val = 10

	var i int = 0
	for {
		if (i>10) { break }

		if node.Rdy(){
			node.Tracef("writing x.Data: %d\n", x.Val.(int))
			node.PrintVals()
			x.Data <- x.Val
			x.Rdy = false
			x.Val = x.Val.(int) + 1
			i = i + 1
		}

		node.Tracef("select\n")
		select {
		case x.Rdy = <-x.Ack:
			node.Tracef("x.Ack read\n")

		}

	}
}

/*
func tbo(a flowgraph.Edge) {

	node := flowgraph.NewNode("tbo", []*flowgraph.Edge{&a}, nil, nil)

	for {
		if node.Rdy() {
			node.Tracef("writing a.Ack\n")
			node.PrintVals()
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
*/

func main() {

	flowgraph.Debug = false
	flowgraph.Indent = false

	start := flowgraph.NewEdge("start",nil)
	rdy0 := flowgraph.NewEdge("rdy0",nil)
	arbit0 := flowgraph.NewEdge("arbit0",nil)
	const1 := flowgraph.NewEdge("const1",nil)
	const1.Val = 1
	sub0 := flowgraph.NewEdge("sub0",nil)
	strcnd0 := flowgraph.NewEdge("strcnd0",int(0))
	strcnd1 := flowgraph.NewEdge("strcnd1",nil)

	go tbi(start)
	go flowgraph.FuncRdy(start, strcnd0, rdy0)
        go flowgraph.FuncArbit(rdy0, strcnd1, arbit0)
	go flowgraph.FuncSub(arbit0, const1, sub0)
	go flowgraph.FuncStrCnd(sub0, strcnd0, strcnd1)
	go flowgraph.FuncConst(const1)

	time.Sleep(1000000000)

}

