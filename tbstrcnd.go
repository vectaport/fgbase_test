package main

import (
	"github.com/vectaport/flowgraph"
	"fmt"
	"reflect"
	"time"
)

func tbi(a flowgraph.Edge) {

	node:=flowgraph.MakeNode("tbi")

	var _a flowgraph.Datum = 0
	_a_rdy := a.Ack_init

	for {


		if _a_rdy {
			fmt.Printf("tbi(%d):  writing a.Data: %d\n", node.Id, _a.(int))
			_a_rdy = false
			a.Data <- _a
			_a = (_a.(int) + 1)%2
		}

		fmt.Printf("tbi(%d):  select", node.Id)
		select {
		case _a_rdy = <-a.Ack:
			fmt.Println("tbi(%d):  a_req read", node.Id)
			
			
		}
	}
	
}

func tbo(x flowgraph.Edge) {
	
	node:=flowgraph.MakeNode("tbo")
	
	var _x flowgraph.Datum
	_x_rdy := x.Data_init

	for {
		// fmt.Println("		tbo:  _x_rdy", _x_rdy)
		if _x_rdy {
			fmt.Printf("		tbo(%d):  writing x.Ack\n", node.Id)
			x.Ack <- true
			_x_rdy = false
		}

		fmt.Println("		tbo:  select")
		select {
		case _x = <-x.Data:
			{
				fmt.Printf("		tbo(%d):  x read %v --  %v\n", node.Id, reflect.TypeOf(_x), _x)
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

