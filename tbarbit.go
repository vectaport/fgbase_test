package main

import (
	"github.com/vectaport/flowgraph"
	"fmt"
	"reflect"
	"time"
)

func tbi(a flowgraph.Edge) {

	nodeid:=flowgraph.MakeNode()

	var _a flowgraph.Datum = 0
	_a_rdy := a.Ack_init

	for {


		if _a_rdy {
			fmt.Printf("tbi(%d):  writing a.Data: %d\n", nodeid, _a.(int))
			_a_rdy = false
			a.Data <- _a
			_a = (_a.(int) + 1)
		}

		fmt.Printf("tbi(%d):  select", nodeid)
		select {
		case _a_rdy = <-a.Ack:
			fmt.Printf("tbi(%d):  a.Ack read\n", nodeid)
			
			
		}
	}
	
}

func tbo(x flowgraph.Edge) {
	
	nodeid:=flowgraph.MakeNode()
	
	var _x flowgraph.Datum
	_x_rdy := x.Data_init

	for {
		// fmt.Println("		tbo:  _x_rdy", _x_rdy)
		if _x_rdy {
			fmt.Printf("		tbo(%d):  writing x.Ack\n", nodeid)
			x.Ack <- true
			_x_rdy = false
		}

		fmt.Println("		tbo:  select")
		select {
		case _x = <-x.Data:
			{
				fmt.Printf("		tbo(%d):  x read %v --  %v\n", nodeid, reflect.TypeOf(_x), _x)
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
	go flowgraph.ArbitFunc(a, b, x)
	go tbo(x)

	time.Sleep(1000000000)

}

