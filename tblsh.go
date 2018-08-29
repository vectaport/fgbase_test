package main

import (
	"math"

	"github.com/vectaport/fgbase"
)

func tbi(x, y fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x, &y}, nil, nil)
	node.RunFunc = tbiRun
	return node
}

func tbiRun(node *fgbase.Node) {
	x := node.Dsts[0]
	y := node.Dsts[1]

	var i uint = 0
	for {
		if i > 10 {
			break
		}
		if node.RdyAll() {
			x.DstPut(i)
			y.DstPut(i)
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	var xv interface{} = float32(0)
	i = 0
	for {
		if i > 9 {
			break
		}
		if node.RdyAll() {
			x.DstPut(xv)
			y.DstPut(xv)
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	xv = uint64(math.MaxUint64)
	var yv interface{} = -1
	i = 0
	for {
		if i > 0 {
			break
		}
		if node.RdyAll() {
			x.DstPut(xv)
			y.DstPut(yv)
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	xv = uint8(0)
	yv = uint64(0)
	i = 0
	for {
		if i > 0 {
			break
		}
		if node.RdyAll() {
			x.DstPut(xv)
			y.DstPut(yv)
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	xv = uint8(0)
	yv = uint16(0)
	i = 0
	for {
		if i > 0 {
			break
		}
		if node.RdyAll() {
			x.DstPut(xv)
			y.DstPut(yv)
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	xv = "Can you left shift a string by an int?"
	yv = uint8(77)
	i = 0
	for {
		if i > 0 {
			break
		}
		if node.RdyAll() {
			x.DstPut(xv)
			y.DstPut(yv)
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	xv = [4]complex128{0 + 0i, 0 + 0i, 0 + 0i, 0 + 0i}
	yv = uint8(77)
	i = 0
	for {
		if i > 0 {
			break
		}
		if node.RdyAll() {
			x.DstPut(xv)
			y.DstPut(yv)
			node.SendAll()
			i = i + 1
		}
		node.RecvOne()
	}

	// read all the acks to clean up
	for {
		node.RecvOne()
	}

}

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil, nil)
	return node

}

func main() {

	fgbase.ConfigByFlag(nil)

	e, n := fgbase.MakeGraph(3, 3)

	n[0] = tbi(e[0], e[1])
	n[1] = fgbase.FuncLsh(e[0], e[1], e[2])
	n[2] = tbo(e[2])

	fgbase.RunAll(n)

}
