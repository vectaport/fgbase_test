package main

import (
	"github.com/vectaport/flowgraph"
	"reflect"
	"time"
)

func tbi(g, a, b flowgraph.Edge) {

	node := flowgraph.MakeNode("tbi")
	
	var _g flowgraph.Datum
	var _a flowgraph.Datum = a.Init_val
	var _b flowgraph.Datum = b.Init_val
	
	_g_rdy := g.Data_init
	_a_rdy := a.Ack_init
	_b_rdy := b.Ack_init
	
	for {
		node.Printf("_g_rdy %v, _a_rdy,_b_rdy %v,%v\n",  _g_rdy, _a_rdy, _b_rdy);
		
		if _a_rdy && _b_rdy && _g_rdy {
			node.ExecCnt()
			_a_rdy = false
			_b_rdy = false
			_g_rdy = false
			g.Ack <- true
			node.Printf("g.Ack written\n");
			a.Data <- _a
			node.Printf("a.Data written\n");
			b.Data <- _b
			node.Printf("b.Data written\n");
			_a = _a.(int) + 1
			_b = _b.(int) + 1
		}
		
		node.Printf("select\n")
		select {
		case _a_rdy = <-a.Ack: {
			node.Printf("a.Ack read\n")
		}
			
		case _b_rdy = <-b.Ack: {
			node.Printf("b.Ack read\n")
		}
			
		case _g = <-g.Data: {
			node.Printf("g.Data read\n")
			flowgraph.Sink(_g)
			_g_rdy = true
		}
		}
		
	}
}

func tbo(x, g flowgraph.Edge) {

	node := flowgraph.MakeNode("tbo")

	var _x flowgraph.Datum
	_x_rdy := x.Data_init
	_g_rdy := g.Ack_init

	for {
		node.Printf("_x_rdy %v, _g_rdy %v\n", _x_rdy, _g_rdy);
		if _x_rdy && _g_rdy {
			node.ExecCnt()
			node.Printf("writing g.Data and x.Ack\n")
			g.Data <- true
			node.Printf("done writing g.Data\n")
			x.Ack <- true
			node.Printf("done writing x.Ack\n")
			_x_rdy = false
			_g_rdy = false
		}

		node.Printf("select\n")
		select {
		case _x = <-x.Data:
			{
				node.Printf("x read %v --  %v\n", reflect.TypeOf(_x), _x)
				_x_rdy = true
			}
		case _g_rdy = <-g.Ack:
			node.Printf("g.Ack read\n")
		}

	}

}

func main() {

	a := flowgraph.MakeEdge(false,true,int(0))
	b := flowgraph.MakeEdge(false,true,int(0))
	x := flowgraph.MakeEdge(false,true,nil)
	g := flowgraph.MakeEdge(true,false,nil)

	go tbi(g, a, b)
	go flowgraph.AddNode(a, b, x)
	go tbo(x, g)

	time.Sleep(1000000000)

}

