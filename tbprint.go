package main

import (
	"os"

	"github.com/vectaport/fgbase"
)

func tbi(x fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil,
		func(n *fgbase.Node) {
			x.DstPut(n.Aux)
			n.Aux = n.Aux.(int) + 1
		})

	node.Aux = 0
	return node

}

func tbo(a fgbase.Edge) fgbase.Node {
	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil, nil)
	return node
}

func check(e error) {
	if e != nil {
		fgbase.StderrLog.Printf("%v\n", e)
		os.Exit(1)
	}
}

func main() {

	fgbase.ConfigByFlag(map[string]interface{}{"sec": 1})

	e, n := fgbase.MakeGraph(2, 3)

	n[0] = tbi(e[0])
	n[1] = fgbase.FuncPrint(e[0], e[1], "%v\n")
	n[2] = tbo(e[1])

	fgbase.RunAll(n)

}
