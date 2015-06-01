package main

import (
	"flag"
//	"math/rand"
	"runtime"
	"strconv"
//	"time"

	"github.com/vectaport/flowgraph"
)

func main() {

	nCorep := flag.Int("ncore", 1 /*runtime.NumCPU()-1*/, "num cores to use, max is "+strconv.Itoa(runtime.NumCPU()))
	flag.Parse()
	runtime.GOMAXPROCS(*nCorep)

	flowgraph.TraceLevel = flowgraph.Q
	flowgraph.TraceSeconds = false
	flowgraph.ChannelSize = 1024

	e,n := flowgraph.MakeGraph(1,2)
	quitChan := make(chan flowgraph.Nada)
 
	n[0] = flowgraph.FuncHttp(e[0], ":8080", quitChan)
	n[1] = flowgraph.FuncKprod(e[0])

	flowgraph.RunAll(n, 0)

	<- quitChan

}

