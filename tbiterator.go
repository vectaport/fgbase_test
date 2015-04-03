package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbi(x flowgraph.Edge) {

	node := flowgraph.MakeNode2("tbi", nil, []*flowgraph.Edge{&x}, nil, nil)

	x.Aux = 10

	var i int = 0
	for {
		if (i>1000000) { break }

		if node.RdyAll(){
			x.Val = x.Aux
			x.Aux = x.Aux.(int) + 1
			node.SendAll()
			i = i + 1
		}

		node.RecvOne()

	}
}


func main() {

	flowgraph.Debug = false
	flowgraph.Indent = false

	e0 := flowgraph.MakeEdge("e0",nil)
	e1 := flowgraph.MakeEdge("e1",nil)
	e2 := flowgraph.MakeEdge("e2",nil)
	e3 := flowgraph.MakeEdge("e3",nil)
	e4 := flowgraph.MakeEdge("e4",nil)
	e5 := flowgraph.MakeEdge("e5",int(0))
	e6 := flowgraph.MakeEdge("e6",nil)

	go tbi(e0)
	go flowgraph.FuncRdy(e0, e5, e1)
        go flowgraph.FuncArbit(e1, e6, e2)
	go flowgraph.FuncSub(e2, e3, e4)
	go flowgraph.FuncStrCnd(e4, e5, e6)
	go flowgraph.FuncConst(e3,1)

	time.Sleep(1000000000)

}

