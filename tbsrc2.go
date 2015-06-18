package main

import (
	"flag"
	"net"
	"time"

	"github.com/vectaport/flowgraph"
)

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	nodeid := flag.Int("nodeid", 0, "base for node ids")
	flag.Parse()
	flowgraph.NodeID = int64(*nodeid)

	flowgraph.TraceLevel = flowgraph.V
	flowgraph.ChannelSize = 16

	ln, err := net.Listen("tcp", "localhost:37777")
	if err != nil {
		flowgraph.StderrLog.Printf("%v\n", err)
		return
	}
	conn, err := ln.Accept()
	if err != nil {
		flowgraph.StderrLog.Printf("%v\n", err)
		return
	}

	e,n := flowgraph.MakeGraph(1,1)

	n[0] = tbo(e[0])
	e[0].Src(&n[0], conn)

	flowgraph.RunAll(n, 4*time.Second)

}

