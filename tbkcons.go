package main

import (
	"flag"

	"github.com/vectaport/fgbase"
)

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, 
		func (n *flowgraph.Node) {
			// time.Sleep(time.Duration(rand.Intn(150000)))
			if n.Cnt%10000==0 {
				flowgraph.StdoutLog.Printf("%.2f: %d (=%.2f hz)\n", flowgraph.TimeSinceStart(), n.Cnt, float64(n.Cnt)/flowgraph.TimeSinceStart())
			}
		})
	return node
}

func main() {

	topicp := flag.String("topic", "test", "Kafka topic")
	flowgraph.ConfigByFlag(map[string]interface{} {"trace": "Q", "chansz": 1024, "sec": 0, "ncore": 1} )
	topic := *topicp

	flowgraph.TraceSeconds = false

	e,n := flowgraph.MakeGraph(1,2)
	quitChan := make(chan struct{})
 
	n[0] = flowgraph.FuncKcons(e[0], topic)
	n[1] = tbo(e[0])

	flowgraph.RunAll(n)

	<- quitChan

}

