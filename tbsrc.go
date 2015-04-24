package main

import (
	"flag"
	"github.com/vectaport/flowgraph"
	"net"
	"time"
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

	e,n := flowgraph.MakeGraph(1,2)

	n[0] = flowgraph.FuncSrc(e[0], conn)
	n[1] = tbo(e[0])

	flowgraph.RunAll(n, 2*time.Second)

}

