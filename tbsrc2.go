package main

import (
	"flag"

	"github.com/vectaport/flowgraphbase"
)

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, 
		func (n *flowgraph.Node) {
			a.Flow = true
			if n.Cnt%10000==0 {
				flowgraph.StdoutLog.Printf("%2.f: %d (%.2f hz)\n", flowgraph.TimeSinceStart(), n.Cnt, float64(n.Cnt)/flowgraph.TimeSinceStart())
			}
		})

	return node
}

func main() {

	nodeidp := flag.Int("nodeid", 0, "base for node ids")
	flowgraph.ConfigByFlag(map[string]interface{} { "sec": 4 })
	flowgraph.NodeID = int64(*nodeidp)

	flowgraph.TraceSeconds = true

	e,n := flowgraph.MakeGraph(1,1)

	n[0] = tbo(e[0])
	e[0].SrcJSON(&n[0], "localhost:37777")

	flowgraph.RunAll(n)

}

