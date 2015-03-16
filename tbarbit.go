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
			node.ExecCnt()
			node.Printf("writing a.Data: %d\n", _a.(int))
			_a_rdy = false
			a.Data <- _a
			_a = (_a.(int) + 1)
		}

		node.Printf("select\n")
		select {
		case _a_rdy = <-a.Ack:
			node.Printf("a.Ack read\n")
			
			
		}
	}
	
}

func tbo(x flowgraph.Edge) {
	
	node:=flowgraph.MakeNode("tbo")
	
	var _x flowgraph.Datum
	_x_rdy := x.Data_init

	for {
		if _x_rdy {
			node.ExecCnt()
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
	b := flowgraph.MakeEdge(false,true,nil)
	x := flowgraph.MakeEdge(false,true,nil)

	go tbi(a)
	go tbi(b)
	go flowgraph.ArbitNode(a, b, x)
	go tbo(x)

	time.Sleep(1000000000)

}

