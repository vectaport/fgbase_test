package main

import (
	"math"

	"github.com/vectaport/fgbase"
)

func tbi(x, y flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x, &y}, nil, nil)
	node.RunFunc = tbiRun
	return node
}

func tbiRun (n *flowgraph.Node) {
	x := n.Dsts[0]
	y := n.Dsts[1]

	var i uint = 0
	for {
		if (i>10) { break }
		if n.RdyAll(){
			x.DstPut(i)
			y.DstPut(uint(2))
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	i = 0
	for {
		if (i>9) { break }
		if n.RdyAll(){
			x.DstPut(uint(i))
			y.DstPut(uint(3))
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	var xv interface{} = uint64(math.MaxUint64)
	var yv interface{} = uint64(1)
	i = 0
	for {
		if (i > 0) { break }
		if n.RdyAll(){
			x.DstPut(xv)
			y.DstPut(yv)
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	xv = uint8(0)
	yv = uint(0)
	i = 0
	for  {
		if (i > 0) { break }
		if n.RdyAll() {
			x.DstPut(xv)
			y.DstPut(yv)
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	xv = uint8(0)
	yv = uint(0)
	i = 0
	for  {
		if (i > 0) { break }
		if n.RdyAll() {
			x.DstPut(xv)
			y.DstPut(yv)
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}


	xv = "Can you right shift a string by an int?"
	yv = uint8(77)
	i = 0
	for  {
		if (i > 0) { break }
		if n.RdyAll() {
			x.DstPut(xv)
			y.DstPut(yv)
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	xv = [4]complex128 {0+0i,0+0i,0+0i,0+0i}
	yv = uint8(77)
	i = 0
	for  {
		if (i > 0) { break }
		if n.RdyAll() {
			x.DstPut(xv)
			y.DstPut(yv)
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}
	

	// read all the acks to clean up
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

	e,n := flowgraph.MakeGraph(3,3)

	n[0] = tbi(e[0], e[1])
	n[1] = flowgraph.FuncRsh(e[0], e[1], e[2])
	n[2] = tbo(e[2])

	flowgraph.RunAll(n)

}

