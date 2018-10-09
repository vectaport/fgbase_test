package main

import (
	"github.com/vectaport/fgbase"
)

func tbiFire(n *fgbase.Node) error {
	x := n.Dsts[0]
	x.DstPut(n.Aux)
	if n.Aux.(int) <= 1 {
		n.Aux = (n.Aux.(int) + 1) % 2
	} else {
		n.Aux = n.Aux.(int) + 1
	}
	return nil
}

func tbi(x fgbase.Edge) fgbase.Node {
	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil, tbiFire)
	return node
}

func tbo(a fgbase.Edge) fgbase.Node {
	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	fgbase.ConfigByFlag(nil)

	e, n := fgbase.MakeGraph(4, 5)

	n[0] = tbi(e[0])
	n[1] = tbi(e[1])
	n[2] = fgbase.FuncSteerv(e[0], e[1], e[2], e[3])
	n[3] = tbo(e[2])
	n[4] = tbo(e[3])

	// initialize different state in the two source testbenches (tbi)
	n[0].Aux = 0
	n[1].Aux = 1000

	fgbase.RunAll(n)

}
