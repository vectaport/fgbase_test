package main

import (
	"github.com/vectaport/flowgraph"
	"reflect"
	"time"
)

func tbi(a flowgraph.Edge) {

	node:=flowgraph.MakeNode("tbi")

	var _a flowgraph.Datum = 0
	_a_rdy := a.Ack_init

	for {


		if _a_rdy {
			node.Printf("writing a.Data: %d\n", _a.(int))
			_a_rdy = false
			a.Data <- _a
			_a = (_a.(int) + 1)%2
		}

		node.Printf("select\n")
		select {
		case _a_rdy = <-a.Ack:
			node.Printf("a_req read\n")
			
			
		}
	}
	
}

func tbo(x flowgraph.Edge) {
	
	node:=flowgraph.MakeNode("tbo")
	
	var _x flowgraph.Datum
	_x_rdy := x.Data_init

	for {
		if _x_rdy {
			node.Printf("writing x.Ack\n")
			x.Ack <- true
			_x_rdy = false
		}

		node.Printf("select\n")
		select {
		case _x = <-x.Data:
			{
				node.Printf("x read %v --  %v\n", reflect.TypeOf(_x), _x)
				_x_rdy = true
			}
		}

	}

}

func main() {

	a := flowgraph.MakeEdge(false,true,nil)
	x := flowgraph.MakeEdge(false,true,nil)
	y := flowgraph.MakeEdge(false,true,nil)

	go tbi(a)
	go flowgraph.StrCndNode(a, x, y)
	go tbo(x)
	go tbo(y)

	time.Sleep(1000000000)

}

