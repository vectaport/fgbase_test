package main

import (
	"github.com/vectaport/flowgraph"
	"math"
	"math/rand"
	"time"
)

const infitesimal=1.e-15

func tbiFire(n *flowgraph.Node) {
	x := n.Dsts[0]
	const sz = 128
	var vec = make([]complex128, sz, sz)
	rand.Seed(0x1515)
	
	delta := 3*2*math.Pi/float64(sz)
	domain := float64(0)

	for i := range vec {
		vec[i] = complex(math.Sin(domain), 0.0)
		domain += delta
	}
	x.Val = vec
}

func tbi(x flowgraph.Edge) flowgraph.Node {
	node:=flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, tbiFire)
	return node
}

func tboFire(n *flowgraph.Node) {
	a := n.Srcs[0]
	b := n.Srcs[1]
	av := a.Val.([]complex128)
	bv := b.Val.([]complex128)
	if (len(av)==len(bv)) {
		for i := range av {
			if (real(av[i])-real(bv[i])) < -infitesimal || (real(av[i])-real(bv[i]))>infitesimal || 
				(imag(av[i])-imag(bv[i])) < -infitesimal || (imag(av[i])-imag(bv[i]))>infitesimal {
				n.Tracef("!SAME:  for %d delta is %v\n", i, av[i]-bv[i])
				n.Tracef("!SAME:  a = %v,  b = %v\n", av[i], bv[i])
				return
			}
		}
		n.Tracef("SAME all differences smaller than %v\n", infitesimal)
		return
	} 
	n.Tracef("!SAME:  different sizes\n")
}

func tbo(a, b flowgraph.Edge) flowgraph.Node {
	node:=flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a, &b}, nil, nil, tboFire)
	return node
}

func main() {

	flowgraph.TraceLevel = flowgraph.V
	
	e,n := flowgraph.MakeGraph(9,7)

	e[7].Const(false)
	e[8].Const(true)

	n[0] = tbi(e[0])

	n[1] = flowgraph.FuncFork(e[0], e[1], e[2])

	n[2] = flowgraph.FuncFft(e[1], e[7], e[3])
	n[3] = flowgraph.FuncPass(e[2], e[4])

	n[4] = flowgraph.FuncFft(e[3], e[8], e[5])
	n[5] = flowgraph.FuncPass(e[4], e[6])

	n[6] = tbo(e[5], e[6])

	flowgraph.RunAll(n, time.Second)

}

