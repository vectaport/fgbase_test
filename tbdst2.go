package main

import (
	"flag"
	"time"

	"github.com/vectaport/flowgraph"
)

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) {
			x.Val = x.Aux
			x.Aux = x.Aux.(int) + 1
			if n.Cnt%10000==0 {
				flowgraph.StdoutLog.Printf("%2.f: %d (rps=%.2f)\n", flowgraph.TimeSinceStart(), n.Cnt, float64(n.Cnt)/flowgraph.TimeSinceStart())
				
			}
		})
	x.Aux = 0
	return node

}

func main() {

	tracep := flag.String("trace", "V", "trace level, Q|V|VV|VVV|VVVV")
	nodeid := flag.Int("nodeid", 0, "base for node ids")
	chanszp := flag.Int("chansz", 1, "channel size")
	flag.Parse()
	flowgraph.NodeID = int64(*nodeid)

	flowgraph.TraceLevel = flowgraph.TraceLevels[*tracep]
	flowgraph.TraceSeconds = true
	flowgraph.ChannelSize = *chanszp

	time.Sleep(1*time.Second)

	e,n := flowgraph.MakeGraph(1,1)

	n[0] = tbi(e[0])
	e[0].Dst(&n[0], "localhost:37777")


	flowgraph.RunAll(n, 8*time.Second)

}

