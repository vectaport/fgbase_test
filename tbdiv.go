package main

import (
	"math"

	"github.com/vectaport/flowgraph"
)

func tbi(x, y flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x, &y}, nil, nil)
	node.RunFunc = tbiRun
	return node
}

func tbiRun (n *flowgraph.Node) {
	x := n.Dsts[0]
	y := n.Dsts[1]

	i := 0
	for {
		if (i>10) { break }
		if n.RdyAll(){
			x.Val = i
			y.Val = 2
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	var xv interface{} = uint64(math.MaxUint64)
	var yv interface{} = -1
	i = 0
	for {
		if (i > 0) { break }
		if n.RdyAll(){
			x.Val = xv
			y.Val = yv
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	xv = int8(0)
	yv =  uint64(0)
	i = 0
	for  {
		if (i > 0) { break }
		if n.RdyAll() {
			x.Val = xv
			y.Val = yv
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	xv = int8(0)
	yv = int16(0)
	i = 0
	for  {
		if (i > 0) { break }
		if n.RdyAll() {
			x.Val = xv
			y.Val = yv
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}


	xv = "Can you divide an int to a string?"
	yv= int8(77)
	i = 0
	for  {
		if (i > 0) { break }
		if n.RdyAll() {
			x.Val = xv
			y.Val = yv
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	xv = [4]complex128 {0+0i,0+0i,0+0i,0+0i}
	yv = int8(77)
	i = 0
	for  {
		if (i > 0) { break }
		if n.RdyAll() {
			x.Val = xv
			y.Val = yv
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}
	

	// read all the acks to clean up
	for  {
		n.RecvOne()
	}
	

}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node

}

func main() {

	flowgraph.ConfigByFlag(nil)

	e,n := flowgraph.MakeGraph(3,3)

	n[0] = tbi(e[0], e[1])
	n[1] = flowgraph.FuncDiv(e[0], e[1], e[2])
	n[2] = tbo(e[2])

	flowgraph.RunAll(n)

}

