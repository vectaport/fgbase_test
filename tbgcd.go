package main

import (
	"math/rand"

	"github.com/vectaport/fgbase"
)

func tbm(x fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil,
		func(n *fgbase.Node) error { n.Dsts[0].DstPut(rand.Intn(15) + 1); return nil })
	return node
}

func tbn(x fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil,
		func(n *fgbase.Node) error { n.Dsts[0].DstPut(rand.Intn(15) + 1); return nil })
	return node
}

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	fgbase.ConfigByFlag(nil)

	e, n := fgbase.MakeGraph(11, 10)

	e[7].Val = 0

	n[0] = tbm(e[0])
	n[1] = tbn(e[1])

	n[2] = fgbase.FuncRdy(e[0], e[7], e[2])
	n[3] = fgbase.FuncRdy(e[1], e[7], e[3])

	n[4] = fgbase.FuncEither(e[2], e[10], e[4])
	n[5] = fgbase.FuncEither(e[3], e[8], e[5])

	n[6] = fgbase.FuncMod(e[4], e[5], e[6])

	n[7] = fgbase.FuncSteerc(e[6], e[7], e[8])
	n[8] = fgbase.FuncSteerv(e[6], e[5], e[9], e[10])

	n[9] = tbo(e[9])

	fgbase.RunAll(n)

}
