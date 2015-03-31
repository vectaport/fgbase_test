package main

import (
	"github.com/vectaport/flowgraph"
	"math"
	"time"
)

func tbi(x, y flowgraph.Edge) {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x, &y}, nil)

	x.Val = 0
	y.Val = 0

	var _i int = 0
	for {
		if (_i>10) { break }

		if node.Rdy() {
			node.TraceVals()
			node.Tracef("writing x.Data and y.Data: %d,%d\n", x.Val.(int), y.Val.(int))
			x.Rdy = false
			y.Rdy = false
			x.Data <- x.Val
			y.Data <- y.Val
			x.Val = x.Val.(int) + 2
			y.Val = y.Val.(int) + 1
			_i = _i + 1
		}

		node.Select()

	}

	x.Val = float32(0)
	y.Val = float32(0)
	_i = 0

	for {
		if (_i>9) { break }
		if node.Rdy(){
			node.TraceVals()
			node.Tracef("writing x.Data and y.Data: %f,%f\n", x.Val.(float32), y.Val.(float32))
			x.Rdy = false
			y.Rdy = false
			x.Data <- x.Val
			y.Data <- y.Val
			x.Val = x.Val.(float32) - 1.
			y.Val = y.Val.(float32) + 1.
			_i = _i + 1
		}

		node.Select()
		
	}
	
	x.Val = uint64(math.MaxUint64)
	y.Val = -1
	_i = 0

	for {
		if (_i > 0) { break }

		if node.Rdy(){
			node.TraceVals()
			node.Tracef("writing x.Data and y.Data: %v,%v\n", x.Val, y.Val)
			x.Rdy = false
			y.Rdy = false
			x.Data <- x.Val
			y.Data <- y.Val
			_i = _i + 1
		}

		node.Select()

	}

	x.Val = int8(-1)
	y.Val = uint64(math.MaxUint64)
	_i = 0

	for  {
		if (_i > 0) { break }

		if node.Rdy(){
			node.TraceVals()
			node.Tracef("writing x.Data and y.Data: %v,%v\n", x.Val, y.Val)
			x.Rdy = false
			y.Rdy = false
			x.Data <- x.Val
			y.Data <- y.Val
			_i = _i + 1
		}

		node.Select()

	}

	x.Val = int8(-1)
	y.Val = uint32(math.MaxUint32)
	_i = 0

	for  {
		if (_i > 0) { break }

		if node.Rdy(){
			node.TraceVals()
			node.Tracef("writing x.Data and y.Data: %v,%v\n", x.Val, y.Val)
			x.Rdy = false
			y.Rdy = false
			x.Data <- x.Val
			y.Data <- y.Val
			_i = _i + 1
		}

		node.Select()

	}

}

func tbo(a flowgraph.Edge) {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil)

	for {
		if node.Rdy() {
			node.TraceVals()
			node.Tracef("writing a.Ack\n")
			a.Ack <- true
			a.Rdy = false
		}

		node.Select()

	}

}

func main() {

	a := flowgraph.MakeEdge("a",nil)
	b := flowgraph.MakeEdge("b",nil)
	x := flowgraph.MakeEdge("x",nil)

	go tbi(a, b)
	go flowgraph.FuncSub(a, b, x)
	go tbo(x)

	time.Sleep(1000000000)

}

