package main

import (
	"github.com/vectaport/flowgraph"
	"math/rand"
	"time"
)

func tbiFire(n *flowgraph.Node) {
	x := n.Dsts[0]
	var vec = make([]complex128, 128, 128)
	rand.Seed(0x1515)
	for i := range vec {
		vec[i] = complex(rand.Float64(), rand.Float64())
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
			if av[i] != bv[i] { break }
		}
		n.Tracef("SAME\n")
		return
	} 
	n.Tracef("!SAME\n")
}

func tbo(a, b flowgraph.Edge) {
	node:=flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a, &b}, nil, nil, tboFire)
	node.Run()
}

func main() {

	flowgraph.TraceLevel = flowgraph.V
	flowgraph.Indent = false

	e0 := flowgraph.MakeEdge("e0",nil)
	e1 := flowgraph.MakeEdge("e1",nil)
	e2 := flowgraph.MakeEdge("e2",nil)
	e3 := flowgraph.MakeEdge("e3",nil)
	e4 := flowgraph.MakeEdge("e4",nil)
	e5 := flowgraph.MakeEdge("e5",nil)
	e6 := flowgraph.MakeEdge("e6",nil)

	go tbi(e0)

	go flowgraph.FuncFork(e0, e1, e2)

	go flowgraph.FuncFft(e1, e3)
	go flowgraph.FuncPass(e2, e4)

//	go flowgraph.FuncFft(e3, e5)
	go flowgraph.FuncPass(e3, e5)
	go flowgraph.FuncPass(e4, e6)

	go tbo(e5, e6)

	time.Sleep(1000000000)
	flowgraph.StdoutLog.Printf("\n")

}

