package main

import (
	"bufio"
	"github.com/vectaport/flowgraph"
	"net"
	"time"
)

func tbo(a flowgraph.Edge) {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	node.Run()
}

func main() {

	flowgraph.TraceLevel = flowgraph.V
	flowgraph.TraceIndent = false

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
	reader:= bufio.NewReader(conn)

	e0 := flowgraph.MakeEdge("e0",nil)

	go flowgraph.FuncSrc(e0, reader)
	go tbo(e0)

	time.Sleep(time.Hour)
	flowgraph.StdoutLog.Printf("\n")

}

