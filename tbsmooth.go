package main

import (
	"flag"
	"runtime"
	"strconv"
	"time" 

	"github.com/vectaport/flowgraph"
	"github.com/vectaport/flowgraph/imagelab"
)

func main() {

	ncorep := flag.Int("ncore", runtime.NumCPU()-1, "num cores to use, max is "+strconv.Itoa(runtime.NumCPU()))
	secp := flag.Int("sec", 100, "seconds to run")
	nsmoothp := flag.Int("nsmooth", 1, "number of smoothing operations in a pipeline")
	tracep := flag.String("trace", "V", "trace level, Q|V|VV|VVV|VVVV")
	flag.Parse()
	runtime.GOMAXPROCS(*ncorep)
	sec := *secp
	nsmooth := *nsmoothp

	flowgraph.TraceLevel = flowgraph.TraceLevels[*tracep]

	e,n := flowgraph.MakeGraph(nsmooth+1,nsmooth+2)
 
	n[0] = imagelab.FuncCapture(e[0])
	for i:= 0; i<nsmooth; i++ {
		n[i+1] = imagelab.FuncSmooth(e[i], e[i+1])
	}
	n[nsmooth+1] = imagelab.FuncDisplay(e[nsmooth], nil)

	flowgraph.RunAll(n, time.Duration(sec)*time.Second)

}

