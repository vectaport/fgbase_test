package main

import (
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

	flowgraph.ConfigByFlag(map[string]interface{} {"sec": 1})

	e,n := flowgraph.MakeGraph(2,3)

	n[0] = tbi(e[0])
	n[1] = flowgraph.FuncPrint(e[0], e[1], "%v\n")
	n[2] = tbo(e[1])

	flowgraph.RunAll(n)

}

