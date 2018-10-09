package main

import (
	"math/rand"

	"github.com/vectaport/fgbase"
)

func tbi(x fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil,
		func(n *fgbase.Node) error { n.Dsts[0].DstPut(rand.Intn(7) + 1); return nil })
	return node
}

func main() {

	fgbase.ConfigByFlag(nil)

	e, n := fgbase.MakeGraph(7, 5)

	e[3].Const(1)
	e[5].Val = 0

	n[0] = tbi(e[0])
	n[1] = fgbase.FuncRdy(e[0], e[5], e[1])
	n[2] = fgbase.FuncEither(e[1], e[6], e[2])
	n[3] = fgbase.FuncSub(e[2], e[3], e[4])
	n[4] = fgbase.FuncSteerc(e[4], e[5], e[6])

	fgbase.RunAll(n)

}
