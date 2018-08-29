package main

import (
	"flag"
	"time"

	"github.com/vectaport/fgbase"
)

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) {
			if n.Aux.(int)%3==0 {
				s := []int{0,1,2,3,4,5,6,7}
				x.DstPut(s)
			} else if n.Aux.(int)%3==1 {
				x.DstPut(float32(n.Aux.(int))+.5)
			} else {
				x.DstPut(n.Aux)
			}

			n.Aux = n.Aux.(int) + 1
			if n.Cnt%10000==0 {
				flowgraph.StdoutLog.Printf("%2.f: %d (%.2f hz)\n", flowgraph.TimeSinceStart(), n.Cnt, float64(n.Cnt)/flowgraph.TimeSinceStart())
				
			}
		})
	node.Aux = 0
	return node

}

func main() {

	nodeid := flag.Int("nodeid", 0, "base for node ids")
	flowgraph.ConfigByFlag(map[string]interface{} { "sec": 4 })
	flowgraph.NodeID = int64(*nodeid)

	flowgraph.TraceSeconds = true

	time.Sleep(1*time.Second)

	e,n := flowgraph.MakeGraph(1,1)

	n[0] = tbi(e[0])
	e[0].DstJSON(&n[0], "localhost:37777")

	flowgraph.RunAll(n)

}

