package main

import (
	"flag"
	"runtime"
	"strconv"
	"time"

	"github.com/vectaport/flowgraph"
)

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			// time.Sleep(1)
			if x.Aux==nil { x.Aux = 0 }
			x.Val = x.Aux
			x.Aux = x.Aux.(int) + 1
		})
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, 
		func (n *flowgraph.Node) {
			// time.Sleep(time.Duration(rand.Intn(150000)))
			if n.Cnt%100000==0 {
				flowgraph.StdoutLog.Printf("%.2f: %d (rps=%.2f)\n", flowgraph.TimeSinceStart(), n.Cnt, float64(n.Cnt)/flowgraph.TimeSinceStart())
			}
		})
	return node
}

func main() {

	nCorep := flag.Int("ncore", runtime.NumCPU()-1, "num cores to use, max is "+strconv.Itoa(runtime.NumCPU()))
	flag.Parse()
	runtime.GOMAXPROCS(*nCorep)

	flowgraph.TraceLevel = flowgraph.Q
	flowgraph.TraceSeconds = true

	e,n := flowgraph.MakeGraph(2,3)
 
	n[0] = tbi(e[0])
	n[1] = flowgraph.FuncFIFO(e[0], e[1], 4096)
	n[2] = tbo(e[1])

	flowgraph.RunAll(n, time.Duration(60)*time.Second)

}

