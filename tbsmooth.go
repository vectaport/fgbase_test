package main

import (
	"flag"
	"runtime"
	"strconv"
	"time" 

	"github.com/vectaport/flowgraph"
)

func main() {

	ncorep := flag.Int("ncore", runtime.NumCPU()-1, "num cores to use, max is "+strconv.Itoa(runtime.NumCPU()))
	secp := flag.Int("sec", 100, "seconds to run")
	postp := flag.Bool("post", false, "post run dump of nodes")
	nsmoothp := flag.Int("nsmooth", 1, "number of smoothing operations in a pipeline")
	tracep := flag.String("trace", "Q", "trace level, Q|V|VV|VVV|VVVV")
	flag.Parse()
	runtime.GOMAXPROCS(*ncorep)
	flowgraph.PostDump = *postp
	sec := *secp
	nsmooth := *nsmoothp

	flowgraph.TraceLevel = flowgraph.TraceLevels[*tracep]

	e,n := flowgraph.MakeGraph(nsmooth+1,nsmooth+2)
 
	n[0] = flowgraph.FuncCapture(e[0])
	for i:= 0; i<nsmooth; i++ {
		n[i+1] = flowgraph.FuncSmooth(e[i], e[i+1])
	}
	n[nsmooth+1] = flowgraph.FuncDisplay(e[nsmooth], nil)

	flowgraph.RunAll(n, time.Duration(sec)*time.Second)

}

