package main

import (
	"flag"

	"github.com/vectaport/flowgraph"
)

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, 
		func (n *flowgraph.Node) {
			// time.Sleep(time.Duration(rand.Intn(150000)))
			if n.Cnt%10000==0 {
				flowgraph.StdoutLog.Printf("%.2f: %d (rps=%.2f)\n", flowgraph.TimeSinceStart(), n.Cnt, float64(n.Cnt)/flowgraph.TimeSinceStart())
			}
		})
	return node
}

func main() {

	topicp := flag.String("topic", "test", "Kafka topic")
	flag.Parse()
	topic := *topicp

	flowgraph.TraceLevel = flowgraph.Q
	flowgraph.TraceSeconds = false
	flowgraph.ChannelSize = 1024

	e,n := flowgraph.MakeGraph(1,2)
	quitChan := make(chan flowgraph.Nada)
 
	n[0] = flowgraph.FuncKcons(e[0], topic)
	n[1] = tbo(e[0])

	flowgraph.RunAll(n, 0)

	<- quitChan

}

