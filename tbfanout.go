package main

import (
	"github.com/vectaport/fgbase"
)

func tbi(x fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil,
		func(n *fgbase.Node) error {
			x.DstPut(n.Aux)
			n.Aux = n.Aux.(int) + 1
			return nil
		})
	node.Aux = 0
	return node
}

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	fgbase.ConfigByFlag(nil)

	e, n := fgbase.MakeGraph(2, 4)

	n[0] = tbi(e[0])
	n[1] = fgbase.FuncPass(e[0], e[1])
	n[2] = tbo(e[1])
	n[3] = tbo(e[1])

	fgbase.RunAll(n)
}
