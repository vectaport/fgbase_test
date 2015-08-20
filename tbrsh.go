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

func tbiRun (node *flowgraph.Node) {
	x := node.Dsts[0]
	y := node.Dsts[1]

	n.Aux = uint(0)
	y.Aux = uint(0)
	var i uint = 0
	for {
		if (i>10) { break }
		if node.RdyAll(){
			x.Val = n.Aux
			y.Val = y.Aux
			n.Aux = x.Aux.(uint) + 1
			y.Aux = y.Aux.(uint) + 1
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	n.Aux = float32(0)
	y.Aux = float32(0)
	i = 0
	for {
		if (i>9) { break }
		if node.RdyAll(){
			x.Val = n.Aux
			y.Val = y.Aux
			n.Aux = x.Aux.(float32) + 1
			y.Aux = y.Aux.(float32) + 1
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	n.Aux = uint64(math.MaxUint64)
	y.Aux = -1
	i = 0
	for {
		if (i > 0) { break }
		if node.RdyAll(){
			x.Val = n.Aux
			y.Val = y.Aux
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	n.Aux = uint8(0)
	y.Aux = uint64(0)
	i = 0
	for  {
		if (i > 0) { break }
		if node.RdyAll() {
			x.Val = n.Aux
			y.Val = y.Aux
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	n.Aux = uint8(0)
	y.Aux = uint16(0)
	i = 0
	for  {
		if (i > 0) { break }
		if node.RdyAll() {
			x.Val = n.Aux
			y.Val = y.Aux
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}


	n.Aux = "Can you right shift a string by an int?"
	y.Aux = uint8(77)
	i = 0
	for  {
		if (i > 0) { break }
		if node.RdyAll() {
			x.Val = n.Aux
			y.Val = y.Aux
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	n.Aux = [4]complex128 {0+0i,0+0i,0+0i,0+0i}
	y.Aux = uint8(77)
	i = 0
	for  {
		if (i > 0) { break }
		if node.RdyAll() {
			x.Val = n.Aux
			y.Val = y.Aux
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}
	

	// read all the acks to clean up
	for  {
		node.RecvOne()
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
	n[1] = flowgraph.FuncRsh(e[0], e[1], e[2])
	n[2] = tbo(e[2])

	flowgraph.RunAll(n)

}

