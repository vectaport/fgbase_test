package main

import (
	"flag"
//	"math/rand"
	"runtime"
	"strconv"
//	"time"

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

	nCorep := flag.Int("ncore", 1 /*runtime.NumCPU()-1*/, "num cores to use, max is "+strconv.Itoa(runtime.NumCPU()))
	flag.Parse()
	runtime.GOMAXPROCS(*nCorep)

	flowgraph.TraceLevel = flowgraph.Q
	flowgraph.TraceSeconds = true

	e,n := flowgraph.MakeGraph(1,2)
	quitChan := make(chan flowgraph.Nada)
 
	n[0] = flowgraph.FuncHttp(e[0], ":8080", quitChan)
	n[1] = tbo(e[0])

	flowgraph.RunAll(n, 0)

	<- quitChan

}

