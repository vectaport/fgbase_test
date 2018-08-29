package main

import (
	"flag"
	"time"

	"github.com/vectaport/fgbase"
	"github.com/vectaport/fgbase/imglab"
)

func main() {

	testp := flag.Bool("test", false, "test mode")
	flowgraph.ConfigByFlag(nil)
	test := *testp


	var quitChan chan struct{}
	if !test {
		quitChan = make(chan struct{})
		flowgraph.RunTime = 0
	} else {
		flowgraph.RunTime = 1*time.Second
	}

	e,n := flowgraph.MakeGraph(1,2)
 
	n[0] = imglab.FuncCapture(e[0])
	n[1] = imglab.FuncDisplay(e[0], quitChan)

	flowgraph.RunAll(n)

	if !test {
		<- quitChan
	}

}

