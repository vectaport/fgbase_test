package main

import (
	"flag"
//	"math/rand"
	"runtime"
	"strconv"
//	"time"

	"github.com/vectaport/flowgraphbase"
)

func main() {

	flowgraph.ConfigByFlag(map[string]interface{} {"trace": "Q", "chansz": 1024, "sec": 0, "ncore": 1} )

	flowgraph.TraceLevel = flowgraph.Q
	flowgraph.TraceSeconds = false
	flowgraph.ChannelSize = 1024

	e,n := flowgraph.MakeGraph(1,2)
	quitChan := make(chan struct{})
 
	n[0] = flowgraph.FuncHTTP(e[0], ":8080", quitChan)
	n[1] = flowgraph.FuncKprod(e[0])

	flowgraph.RunAll(n)

	<- quitChan

}

