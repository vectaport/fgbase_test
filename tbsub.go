package main

import (
	"github.com/vectaport/flowgraph"
	"math"
	"time"
)

func tbi(x, y flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x, &y}, nil, nil)
	node.RunFunc = tbiRun
	return node
}

func tbiRun(node *flowgraph.Node) {
	x := node.Dsts[0]
	y := node.Dsts[1]

	x.Aux = 0
	y.Aux = 0
	var i int = 0
	for {
		if (i>10) { break }
		if node.RdyAll() {
			x.Val = x.Aux
			y.Val = y.Aux
			x.Aux = x.Aux.(int) + 2
			y.Aux = y.Aux.(int) + 1
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	x.Aux = float32(0)
	y.Aux = float32(0)
	i = 0
	for {
		if (i>9) { break }
		if node.RdyAll(){
			x.Val = x.Aux
			y.Val = y.Aux
			x.Aux = x.Aux.(float32) - 1.
			y.Aux = y.Aux.(float32) + 1.
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}
	
	x.Aux = uint64(math.MaxUint64)
	y.Aux = -1
	i = 0
	for {
		if (i > 0) { break }
		if node.RdyAll(){
			x.Val = x.Aux
			y.Val = y.Aux
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	x.Aux = int8(-1)
	y.Aux = uint64(math.MaxUint64)
	i = 0
	for  {
		if (i > 0) { break }
		if node.RdyAll(){
			x.Val = x.Aux
			y.Val = y.Aux
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	x.Aux = int8(-1)
	y.Aux = uint32(math.MaxUint32)
	i = 0
	for  {
		if (i > 0) { break }
		if node.RdyAll(){
			x.Val = x.Aux
			y.Val = y.Aux
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	for  {
		node.RecvOne()
	}

}

func tbo(a flowgraph.Edge) flowgraph.Node {
	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	flowgraph.TraceLevel = flowgraph.V

	e := flowgraph.MakeEdges(3)

	var n [3]flowgraph.Node
	n[0] = tbi(e[0], e[1])
	n[1] = flowgraph.FuncSub(e[0], e[1], e[2])
	n[2] = tbo(e[2])

	flowgraph.RunAll(n[:], time.Second)

}

