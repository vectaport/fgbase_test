package main

import (
	"flag"
	"fmt"
//	"math/rand"
	"runtime"
	"strconv"
	"time"

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

	nCorep := flag.Int("ncore", runtime.NumCPU()-1, "num cores to use, max is "+strconv.Itoa(runtime.NumCPU()))
	nPortp := flag.Int("nport", 1, "number of server ports")
	testp := flag.Bool("test", false, "test mode")
	flag.Parse()
	runtime.GOMAXPROCS(*nCorep)
	nPort := *nPortp
	test := *testp

	flowgraph.TraceLevel = flowgraph.Q
	flowgraph.TraceSeconds = true
	flowgraph.ChannelSize = 1024

	var quitChan chan flowgraph.Nada
	var wait time.Duration
	if !test {
		quitChan = make(chan flowgraph.Nada)
	} else {
		wait = 5
	}

	e,n := flowgraph.MakeGraph(1,nPort+1)

	for i := 0; i<nPort; i++ {
		n[i] = flowgraph.FuncHttp(e[0], fmt.Sprintf(":%d", 8080+i), quitChan)
	}

	n[nPort] = tbo(e[0])

	flowgraph.RunAll(n, time.Duration(wait*time.Second))

	if !test {
		<- quitChan
	}
}

