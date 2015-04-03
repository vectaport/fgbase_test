package main

import (
	"github.com/vectaport/flowgraph"
	"math"
	"time"
)

func tbi(x, y flowgraph.Edge) {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x, &y}, nil, nil)

	x.Aux = 0
	y.Aux = 0

	var i int = 0
	for {
		if (i>10) { break }

		if node.RdyAll(){
			x.Val = x.Aux
			y.Val = y.Aux
			x.Aux = x.Aux.(int) + 1
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
			x.Aux = x.Aux.(float32) + 1
			y.Aux = y.Aux.(float32) + 1
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
			y.Val = x.Aux
			node.SendAll()
			i = i + 1
		}

		node.RecvOne()

	}

	x.Aux = int8(0)
	y.Aux = uint64(0)
	i = 0

	for  {
		if (i > 0) { break }

		if node.RdyAll() {
			x.Val = x.Aux
			y.Val = x.Aux
			node.SendAll()
			i = i + 1
		}

		node.RecvOne()

	}

	x.Aux = int8(0)
	y.Aux = int16(0)
	i = 0

	for  {
		if (i > 0) { break }

		if node.RdyAll() {
			x.Val = x.Aux
			y.Val = x.Aux
			node.SendAll()
			i = i + 1
		}

		node.RecvOne()

	}

}

func tbo(a flowgraph.Edge) {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)

	for {
		if node.RdyAll() {
			node.SendAll()
		}

		node.RecvOne()

	}

}

func main() {

	flowgraph.Indent = false
	flowgraph.Debug = false

	e0 := flowgraph.MakeEdge("e0",nil)
	e1 := flowgraph.MakeEdge("e1",nil)
	e2 := flowgraph.MakeEdge("e2",nil)

	go tbi(e0, e1)
	go flowgraph.FuncAdd(e0, e1, e2)
	go tbo(e2)

	time.Sleep(1000000000)

}

