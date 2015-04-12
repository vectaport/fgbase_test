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

func tbi(x flowgraph.Edge) {
	node:=flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, tbiFire)
	node.Run()
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

func tbo(a, b flowgraph.Edge) {
	node:=flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a, &b}, nil, nil, tboFire)
	node.Run()
}

func main() {

	flowgraph.TraceLevel = flowgraph.V
	flowgraph.TraceIndent = false

	e0 := flowgraph.MakeEdge("e0",nil)
	e1 := flowgraph.MakeEdge("e1",nil)
	e2 := flowgraph.MakeEdge("e2",nil)
	e3 := flowgraph.MakeEdge("e3",nil)
	e4 := flowgraph.MakeEdge("e4",nil)
	e5 := flowgraph.MakeEdge("e5",nil)
	e6 := flowgraph.MakeEdge("e6",nil)

	cfalse := flowgraph.MakeEdgeConst("cfalse", false)
	ctrue := flowgraph.MakeEdgeConst("ctrue", true)

	go tbi(e0)

	go flowgraph.FuncFork(e0, e1, e2)

	go flowgraph.FuncFft(e1, cfalse, e3)
	go flowgraph.FuncPass(e2, e4)

	go flowgraph.FuncFft(e3, ctrue, e5)
	go flowgraph.FuncPass(e4, e6)

	go tbo(e5, e6)

	time.Sleep(time.Second)
	flowgraph.StdoutLog.Printf("\n")

}

