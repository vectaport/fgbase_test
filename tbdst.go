package main

import (
	"flag"
	"github.com/vectaport/flowgraph"
	"net"
	"time"
)

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) {
			x.Val = x.Aux
			x.Aux = x.Aux.(int) + 1
		})

	x.Aux = 0
	return node

}

func main() {

	nodeid := flag.Int("nodeid", 0, "base for node ids")
	flag.Parse()
	flowgraph.NodeID = int64(*nodeid)

	flowgraph.TraceLevel = flowgraph.V

	time.Sleep(1*time.Second)
	conn, err := net.Dial("tcp", "localhost:37777")
	if err != nil {
		flowgraph.StderrLog.Printf("%v\n", err)
		return
	}

	e,n := flowgraph.MakeGraph(1,2)

	n[0] = tbi(e[0])
	n[1] = flowgraph.FuncDst(e[0], conn)

	flowgraph.RunAll(n, 2*time.Second)

}

