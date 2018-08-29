package main

import (
	"math/rand"

	"github.com/vectaport/fgbase"
	"github.com/vectaport/fgbase/imglab"
)

func tbiFire(n *fgbase.Node) {
	x := n.Dsts[0]
	l := 1024
	buf := make([]complex128, l, l)
	for i := 0; i < l; i++ {
		buf[i] = complex(rand.Float64(), rand.Float64())
	}
	x.DstPut(buf)
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

	e, n := fgbase.MakeGraph(3, 3)

	e[1].Const(false)

	n[0] = tbi(e[0])
	n[1] = imglab.FuncFFT(e[0], e[1], e[2])
	n[2] = tbo(e[2])

	fgbase.RunAll(n)

}
