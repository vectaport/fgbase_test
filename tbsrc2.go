package main

import (
	"flag"
	"time"

	"github.com/vectaport/flowgraph"
)

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, 
		func (n *flowgraph.Node) {
			if n.Cnt%10000==0 {
				flowgraph.StdoutLog.Printf("%2.f: %d (rps=%.2f)\n", flowgraph.TimeSinceStart(), n.Cnt, float64(n.Cnt)/flowgraph.TimeSinceStart())
			}
		})
	return node
}

func main() {

	tracep := flag.String("trace", "V", "trace level, Q|V|VV|VVV|VVVV")
	nodeidp := flag.Int("nodeid", 0, "base for node ids")
	chanszp := flag.Int("chansz", 1, "channel size")
	flag.Parse()
	flowgraph.NodeID = int64(*nodeidp)

	flowgraph.TraceLevel = flowgraph.TraceLevels[*tracep]
	flowgraph.TraceSeconds = true
	flowgraph.ChannelSize = *chanszp

	e,n := flowgraph.MakeGraph(1,1)

	n[0] = tbo(e[0])
	e[0].Src(&n[0], "localhost:37777")

	flowgraph.RunAll(n, 8*time.Second)

}

