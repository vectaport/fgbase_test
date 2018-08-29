package main

import (
	"github.com/vectaport/fgbase"
)

func tbiFire(n *fgbase.Node) {
	a := n.Srcs[0]
	x := n.Dsts[0]
	y := n.Dsts[1]
	a.Flow = true
	x.DstPut(n.Aux)
	y.DstPut(n.Aux)
	n.Aux = n.Aux.(int) + 1
}

func tbi(a, x, y fgbase.Edge) fgbase.Node {
	node := fgbase.MakeNode("tbi", []*fgbase.Edge{&a}, []*fgbase.Edge{&x, &y}, nil, tbiFire)
	node.Aux = 1
	return node
}

func tboFire(n *fgbase.Node) {
	a := n.Srcs[0]
	x := n.Dsts[0]
	a.Flow = true
	x.DstPut(true)
}

func tbo(a, x fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, []*fgbase.Edge{&x}, nil, tboFire)
	return node

}

func main() {

	fgbase.ConfigByFlag(nil)

	e, n := fgbase.MakeGraph(4, 3)

	e[3].Val = true // initialize data wavefront

	n[0] = tbi(e[3], e[0], e[1])
	n[1] = fgbase.FuncAdd(e[0], e[1], e[2])
	n[2] = tbo(e[2], e[3])

	fgbase.RunAll(n)

}
