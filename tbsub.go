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

func tbiRun(n *flowgraph.Node) {
	x := n.Dsts[0]
	y := n.Dsts[1]

	n.Aux = []int{0, 0}
	var i int = 0
	for {
		if (i>10) { break }
		if n.RdyAll() {
			x.Val = n.Aux.([]int)[0]
			y.Val = n.Aux.([]int)[1]
			n.Aux = []int{x.Val.(int)+1, y.Val.(int)+2}
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}


	n.Aux = []float32{0,0}
	i = 0
	for {
		if (i>9) { break }
		if n.RdyAll(){
			x.Val = n.Aux.([]float32)[0]
			y.Val = n.Aux.([]float32)[1]
			n.Aux = []float32{x.Val.(float32) - 1., y.Val.(float32) + 1.}
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}
	
	n.Aux = []interface{}{uint64(math.MaxUint64), -1}
	i = 0
	for {
		if (i > 0) { break }
		if n.RdyAll(){
			x.Val = n.Aux.([]interface{})[0]
			y.Val = n.Aux.([]interface{})[1]
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	n.Aux = []interface{}{int8(-1), uint64(math.MaxUint64)}
	i = 0
	for  {
		if (i > 0) { break }
		if n.RdyAll(){
			x.Val = n.Aux.([]interface{})[0]
			y.Val = n.Aux.([]interface{})[1]
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	n.Aux = []interface{}{int8(-1), uint32(math.MaxUint32)}
	i = 0
	for  {
		if (i > 0) { break }
		if n.RdyAll(){
			x.Val = n.Aux.([]interface{})[0]
			y.Val = n.Aux.([]interface{})[1]
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

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

	e := flowgraph.MakeEdges(3)

	var n [3]flowgraph.Node
	n[0] = tbi(e[0], e[1])
	n[1] = flowgraph.FuncSub(e[0], e[1], e[2])
	n[2] = tbo(e[2])

	flowgraph.RunAll(n[:])

}

