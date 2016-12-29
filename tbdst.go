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
			x.DstPut(n.Aux)
			n.Aux = n.Aux.(int) + 1
		})

	node.Aux = 0
	return node

}

func main() {

	nodeid := flag.Int("nodeid", 0, "base for node ids")
	flowgraph.ConfigByFlag(map[string]interface{} {"sec": 2})
	flowgraph.NodeID = int64(*nodeid)

	time.Sleep(1*time.Second)
	conn, err := net.Dial("tcp", "localhost:37777")
	if err != nil {
		flowgraph.StderrLog.Printf("%v\n", err)
		return
	}

	e,n := flowgraph.MakeGraph(1,2)

	n[0] = tbi(e[0])
	n[1] = flowgraph.FuncDst(e[0], conn)

	flowgraph.RunAll(n)

}

