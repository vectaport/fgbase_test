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

	flowgraph.ConfigByFlag(map[string]interface{} {"trace": "Q", "chansz": 1024, "sec": 0, "ncore": 1} )

	flowgraph.TraceLevel = flowgraph.Q
	flowgraph.TraceSeconds = false
	flowgraph.ChannelSize = 1024

	e,n := flowgraph.MakeGraph(1,2)
	quitChan := make(chan flowgraph.Nada)
 
	n[0] = flowgraph.FuncHttp(e[0], ":8080", quitChan)
	n[1] = flowgraph.FuncKprod(e[0])

	flowgraph.RunAll(n)

	<- quitChan

}

