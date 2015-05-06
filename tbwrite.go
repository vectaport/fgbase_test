package main

import (
	"flag"
	"os"
	"time"

	"github.com/vectaport/flowgraph"
)

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) {
			x.Val = x.Aux
			x.Aux = x.Aux.(int) + 1
		})

	x.Aux = 0
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

	f, err := os.Create(fileName)
	check(err)

	e,n := flowgraph.MakeGraph(1,2)

	n[0] = tbi(e[0])
	n[1] = flowgraph.FuncWrite(e[0], f)

	flowgraph.RunAll(n, 2*time.Second)

}

