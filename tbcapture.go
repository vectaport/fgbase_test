package main

import (
	"flag"
	"time"

	"github.com/vectaport/flowgraph"
)

func main() {

	testp := flag.Bool("test", false, "test mode")
	flag.Parse()
	test := *testp


	var quitChan chan flowgraph.Nada
	var wait time.Duration
	if !test {
		quitChan = make(chan flowgraph.Nada)
	} else {
		wait = 1
	}

	e,n := flowgraph.MakeGraph(1,2)
 
	n[0] = flowgraph.FuncCapture(e[0])
	n[1] = flowgraph.FuncDisplay(e[0], quitChan)

	flowgraph.TraceLevel = flowgraph.V
	flowgraph.RunAll(n, time.Duration(wait*time.Second))

	if !test {
		<- quitChan
	}

}
