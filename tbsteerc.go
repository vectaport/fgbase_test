package main

import (
	"github.com/vectaport/fgbase"
)

func tbiFire(n *fgbase.Node) {
	x := n.Dsts[0]
	x.DstPut(n.Aux)
	n.Aux = (n.Aux.(int) + 1) % 2
}

func tbi(x fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil, tbiFire)
	node.Aux = 0
	return node

}

func tbo(a fgbase.Edge) fgbase.Node {
	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	fgbase.ConfigByFlag(nil)

	e, n := fgbase.MakeGraph(3, 4)

	n[0] = tbi(e[0])
	n[1] = fgbase.FuncSteerc(e[0], e[1], e[2])
	n[2] = tbo(e[1])
	n[3] = tbo(e[2])

	fgbase.RunAll(n)

}
