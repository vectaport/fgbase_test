package main

import (
	"flag"
	"net"

	"github.com/vectaport/fgbase"
)

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	nodeid := flag.Int("nodeid", 0, "base for node ids")
	fgbase.ConfigByFlag(map[string]interface{}{"sec": 2})
	fgbase.NodeID = int64(*nodeid)

	ln, err := net.Listen("tcp", "localhost:37777")
	if err != nil {
		fgbase.StderrLog.Printf("%v\n", err)
		return
	}
	conn, err := ln.Accept()
	if err != nil {
		fgbase.StderrLog.Printf("%v\n", err)
		return
	}

	e, n := fgbase.MakeGraph(1, 2)

	n[0] = fgbase.FuncSrc(e[0], conn)
	n[1] = tbo(e[0])

	fgbase.RunAll(n)

}
