package main

import (
	"flag"
	"net"
	"time"

	"github.com/vectaport/flowgraph"
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
	flowgraph.ChannelSize = 16

	time.Sleep(1*time.Second)
	conn, err := net.Dial("tcp", "localhost:37777")
	if err != nil {
		flowgraph.StderrLog.Printf("%v\n", err)
		return
	}

	e,n := flowgraph.MakeGraph(1,1)

	n[0] = tbi(e[0])
	e[0].Dst(&n[0], conn)


	flowgraph.RunAll(n, 4*time.Second)

}

