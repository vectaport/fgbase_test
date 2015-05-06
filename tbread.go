package main

import (
	"flag"
	"os"
	"time"

	"github.com/vectaport/flowgraph"
)

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func check(e error) {
	if e != nil {
		flowgraph.StderrLog.Printf("%v\n", e)
		os.Exit(1)
	}
}
		
func main() {

	flag.Parse()
	if len(flag.Args()) == 0  { 
		flag.Usage()
		os.Exit(1)
	}
	fileName := flag.Arg(0)

	flowgraph.TraceLevel = flowgraph.V

	f, err := os.Open(fileName)
	check(err)

	e,n := flowgraph.MakeGraph(1,2)

	n[0] = flowgraph.FuncRead(e[0], f)
	n[1] = tbo(e[0])

	flowgraph.RunAll(n, 2*time.Second)

}

