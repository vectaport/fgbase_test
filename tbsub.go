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

	var _i int = 0
	for {
		if (_i>10) { break }

		if node.RdyAll() {
			x.Val = x.Aux
			y.Val = y.Aux
			x.Aux = x.Aux.(int) + 2
			y.Aux = y.Aux.(int) + 1
			node.SendAll()
			_i = _i + 1
		}

		node.RecvOne()

	}

	x.Aux = float32(0)
	y.Aux = float32(0)
	_i = 0

	for {
		if (_i>9) { break }
		if node.RdyAll(){
			x.Val = x.Aux
			y.Val = y.Aux
			x.Aux = x.Aux.(float32) - 1.
			y.Aux = y.Aux.(float32) + 1.
			node.SendAll()
			_i = _i + 1
		}

		node.RecvOne()
		
	}
	
	x.Aux = uint64(math.MaxUint64)
	y.Aux = -1
	_i = 0

	for {
		if (_i > 0) { break }

		if node.RdyAll(){
			x.Val = x.Aux
			y.Val = y.Aux
			node.SendAll()
			_i = _i + 1
		}

		node.RecvOne()

	}

	x.Aux = int8(-1)
	y.Aux = uint64(math.MaxUint64)
	_i = 0

	for  {
		if (_i > 0) { break }

		if node.RdyAll(){
			x.Val = x.Aux
			y.Val = y.Aux
			node.SendAll()
			_i = _i + 1
		}

		node.RecvOne()

	}

	x.Aux = int8(-1)
	y.Aux = uint32(math.MaxUint32)
	_i = 0

	for  {
		if (_i > 0) { break }

		if node.RdyAll(){
			x.Val = x.Aux
			y.Val = y.Aux
			node.SendAll()
			_i = _i + 1
		}

		node.RecvOne()

	}

}

func tbo(a flowgraph.Edge) {
	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	node.Run()
}

func main() {

	e0 := flowgraph.MakeEdge("e0",nil)
	e1 := flowgraph.MakeEdge("e1",nil)
	e2 := flowgraph.MakeEdge("e2",nil)

	go tbi(e0, e1)
	go flowgraph.FuncSub(e0, e1, e2)
	go tbo(e2)

	time.Sleep(1000000000)

}

