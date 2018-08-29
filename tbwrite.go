package main

import (
	"flag"
	"os"

	"github.com/vectaport/fgbase"
)

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) {
			x.DstPut(n.Aux)
			n.Aux = n.Aux.(int) + 1
		})

	node.Aux = 0
	return node

}

func check(e error) {
	if e != nil {
		flowgraph.StderrLog.Printf("%v\n", e)
		os.Exit(1)
	}
}
		
func main() {

	flowgraph.ConfigByFlag(map[string]interface{} {"sec": 2})
	if len(flag.Args()) == 0  { 
		flag.Usage()
		os.Exit(1)
	}
	fileName := flag.Arg(0)

	f, err := os.Create(fileName)
	check(err)

	e,n := flowgraph.MakeGraph(1,2)

	n[0] = tbi(e[0])
	n[1] = flowgraph.FuncWrite(e[0], f)

	flowgraph.RunAll(n)

}

