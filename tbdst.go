package main

import (
	"github.com/vectaport/flowgraph"
	"net"
	"time"
)

func tbi(x flowgraph.Edge) {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, nil)

	x.Aux = 0

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

	flowgraph.TraceLevel = flowgraph.V
	flowgraph.TraceIndent = false

	time.Sleep(1*time.Second)
	conn, err := net.Dial("tcp", "localhost:37777")
	if err != nil {
		flowgraph.StderrLog.Printf("%v\n", err)
		return
	}

	e0 := flowgraph.MakeEdge("e0",nil)

	go tbi(e0)
	go flowgraph.FuncDst(e0, conn)

	time.Sleep(2*time.Second)
	flowgraph.StdoutLog.Printf("\n")

}

