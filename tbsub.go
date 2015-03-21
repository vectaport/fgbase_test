package main

import (
	"github.com/vectaport/flowgraph"
	"math"
	"reflect"
	"time"
)

func tbi(x, y flowgraph.Edge) {

	node := flowgraph.NewNode("tbi", nil, []*flowgraph.Edge{&x, &y}, nil)

	x.Val = 0
	y.Val = 0

	var _i int = 0
	for {
		if (_i>10) { break }

		if node.Rdy() {
			node.PrintVals()
			node.Printf("writing x.Data and y.Data: %d,%d\n", x.Val.(int), y.Val.(int))
			x.Rdy = false
			y.Rdy = false
			x.Data <- x.Val
			y.Data <- y.Val
			x.Val = x.Val.(int) + 2
			y.Val = y.Val.(int) + 1
			_i = _i + 1
		}

		node.Printf("select\n")
		select {
		case x.Rdy = <-x.Ack:
			node.Printf("x.Ack read\n")

		case y.Rdy = <-y.Ack:
			node.Printf("y.Ack read\n")
		}

	}

	x.Val = float32(0)
	y.Val = float32(0)
	_i = 0

	for {
		if (_i>9) { break }
		if node.Rdy(){
			node.PrintVals()
			node.Printf("writing x.Data and y.Data: %f,%f\n", x.Val.(float32), y.Val.(float32))
			x.Rdy = false
			y.Rdy = false
			x.Data <- x.Val
			y.Data <- y.Val
			x.Val = x.Val.(float32) - 1.
			y.Val = y.Val.(float32) + 1.
			_i = _i + 1
		}

		node.Printf("select\n")
		select {
		case x.Rdy = <-x.Ack:
			node.Printf("x.Ack read\n")

		case y.Rdy = <-y.Ack:
			node.Printf("y.Ack read\n")
		}

	}

	x.Val = uint64(math.MaxUint64)
	y.Val = -1
	_i = 0

	for {
		if (_i > 0) { break }

		if node.Rdy(){
			node.PrintVals()
			node.Printf("writing x.Data and y.Data: %v,%v\n", x.Val, y.Val)
			x.Rdy = false
			y.Rdy = false
			x.Data <- x.Val
			y.Data <- y.Val
			// x.Val = x.Val.(float32) + 1.
			// y.Val = y.Val.(float32) + 1.
			_i = _i + 1
		}

		node.Printf("select\n")
		select {
		case x.Rdy = <-x.Ack:
			node.Printf("x.Ack read\n")

		case y.Rdy = <-y.Ack:
			node.Printf("y.Ack read\n")
		}

	}

	x.Val = int8(0)
	y.Val = uint64(0)
	_i = 0

	for  {
		if (_i > 0) { break }

		if node.Rdy(){
			node.PrintVals()
			node.Printf("writing x.Data and y.Data: %v,%v\n", x.Val, y.Val)
			x.Rdy = false
			y.Rdy = false
			x.Data <- x.Val
			y.Data <- y.Val
			// x.Val = x.Val.(float32) + 1.
			// y.Val = y.Val.(float32) + 1.
			_i = _i + 1
		}

		node.Printf("select\n")
		select {
		case x.Rdy = <-x.Ack:
			node.Printf("x.Ack read\n")

		case y.Rdy = <-y.Ack:
			node.Printf("y.Ack read\n")
		}

	}

}

func tbo(a flowgraph.Edge) {

	node := flowgraph.NewNode("tbo", []*flowgraph.Edge{&a}, nil, nil)

	for {
		if node.Rdy() {
			node.PrintVals()
			node.Printf("writing a.Ack\n")
			a.Ack <- true
			a.Rdy = false
		}

		node.Printf("select\n")
		select {
		case a.Val = <-a.Data:
			{
				node.Printf("a read %v --  %v\n", reflect.TypeOf(a.Val), a.Val)
				a.Rdy = true
			}
		}

	}

}

func main() {

	a := flowgraph.NewEdge("a",nil)
	b := flowgraph.NewEdge("b",nil)
	x := flowgraph.NewEdge("x",nil)

	go tbi(a, b)
	go flowgraph.FuncSub(a, b, x)
	go tbo(x)

	time.Sleep(1000000000)

}

