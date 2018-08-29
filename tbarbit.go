package main

import (
	"math/rand"
	"time"

	"github.com/vectaport/fgbase"
)

func tbi(x fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil,
		func(n *fgbase.Node) {
			x.DstPut(n.Aux)
			n.Aux = (n.Aux.(int) + 1)
			time.Sleep(time.Duration(rand.Intn(10000)) * time.Microsecond)
		})

	return node

}

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil,
		func(n *fgbase.Node) {
			a.Flow = true
			time.Sleep(time.Duration(rand.Intn(10000)) * time.Microsecond)
		})
	return node

}

func main() {

	fgbase.ConfigByFlag(nil)

	e, n := fgbase.MakeGraph(3, 4)

	n[0] = tbi(e[0])
	n[1] = tbi(e[1])
	n[2] = fgbase.FuncArbit(e[0], e[1], e[2])
	n[3] = tbo(e[2])

	n[0].Aux = 0
	n[1].Aux = 1000

	fgbase.RunAll(n)

}
