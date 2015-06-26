package main

import (
	"flag"
	"time"

	"github.com/vectaport/flowgraph"
	"github.com/vectaport/flowgraph/imglab"
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
 
	n[0] = imglab.FuncCapture(e[0])
	n[1] = imglab.FuncDisplay(e[0], quitChan)

	flowgraph.TraceLevel = flowgraph.V
	flowgraph.RunAll(n, time.Duration(wait*time.Second))

	if !test {
		<- quitChan
	}

}

