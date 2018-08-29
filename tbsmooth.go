package main

import (
	"flag"

	"github.com/vectaport/flowgraphbase"
	"github.com/vectaport/flowgraphbase/imglab"
)

func main() {

	nsmoothp := flag.Int("nsmooth", 1, "number of smoothing operations in a pipeline")
	flowgraph.ConfigByFlag(map[string]interface{} {"sec": 100})
	nsmooth := *nsmoothp

	e,n := flowgraph.MakeGraph(nsmooth+1,nsmooth+2)
 
	n[0] = imglab.FuncCapture(e[0])
	for i:= 0; i<nsmooth; i++ {
		n[i+1] = imglab.FuncSmooth(e[i], e[i+1])
	}
	n[nsmooth+1] = imglab.FuncDisplay(e[nsmooth], nil)

	flowgraph.RunAll(n)

}

