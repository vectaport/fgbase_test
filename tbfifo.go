package main

import (
	"flag"
	// "math/rand"
	"runtime"
	"strconv"
	"time"

	"github.com/vectaport/fgbase"
)

var maxTbi = 0
var maxTbo = 0
var capTbi = 0
var capTbo = 0
var lenTbi = 0
var lenTbo = 0
var oneDelay = true

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			capTbi = cap((*x.Data)[0]) 
			lenTbi = len((*x.Data)[0]) 
			if len((*x.Data)[0])>maxTbi { maxTbi = len((*x.Data)[0])}
			// time.Sleep(1)
			if n.Aux==nil { x.Aux = 0 }
			x.Val = n.Aux
			n.Aux = x.Aux.(int) + 1
		})
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, 
		func (n *flowgraph.Node) {
			capTbo = cap(a.Ack) 
			lenTbo = len(a.Ack) 
			if len(a.Ack)>maxTbo { maxTbo = len(a.Ack) }
			if oneDelay && false {
				time.Sleep(time.Duration(time.Second*time.Duration(10)))
				oneDelay = false
			}
			if n.Cnt%100000==0 {
				flowgraph.StdoutLog.Printf("%.2f: %d (%.2f hz)  datamax=%d, ackmax=%d, datalen=%d, acklen=%d, datacap=%d, ackcap=%d\n", flowgraph.TimeSinceStart(), n.Cnt, float64(n.Cnt)/flowgraph.TimeSinceStart(),
					maxTbi, maxTbo, lenTbi, lenTbo, capTbi, capTbo)
				maxTbi = 0
				maxTbo = 0
			}
		})
	return node
}

func main() {

	flowgraph.ConfigByFlag(map[string]interface{} {"ncore": 1, "chansz": 1024, "sec": 60})

	flowgraph.TraceSeconds = true

	e,n := flowgraph.MakeGraph(1,2)
 
	n[0] = tbi(e[0])
	n[1] = tbo(e[0])

	flowgraph.RunAll(n)

}

