package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/vectaport/flowgraphbase"
	"github.com/vectaport/flowgraphbase/weblab"
)

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, 
		func (n *flowgraph.Node) {
			// time.Sleep(time.Duration(rand.Intn(150000)))
			a.Flow = true
			if n.Cnt%10000==0 {
				flowgraph.StdoutLog.Printf("%.2f: %d (%.2f hz)\n", flowgraph.TimeSinceStart(), n.Cnt, float64(n.Cnt)/flowgraph.TimeSinceStart())
			}
		})
	return node
}

func main() {

	nPortp := flag.Int("nport", 1, "number of server ports")
	testp := flag.Bool("test", false, "test mode")
	flowgraph.ConfigByFlag(map[string]interface{} {"trace": "Q", "chansz": 1024})
	nPort := *nPortp
	test := *testp

	flowgraph.TraceSeconds = true

	var quitChan chan struct{}
	if !test {
		quitChan = make(chan struct{})
		flowgraph.RunTime = 0
	} else {
		flowgraph.RunTime = 10*time.Second
	}

	e,n := flowgraph.MakeGraph(1,nPort+1)

	for i := 0; i<nPort; i++ {
		n[i] = weblab.FuncHTTP(e[0], fmt.Sprintf(":%d", 8080+i), quitChan)
	}

	n[nPort] = tbo(e[0])

	flowgraph.RunAll(n)

	if !test {
		<- quitChan
	}
}

