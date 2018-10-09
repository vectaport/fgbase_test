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

func tbiRun(n *fgbase.Node) error {
	x := n.Dsts[0]
	y := n.Dsts[1]

	n.Aux = 0
	var i int = 0
	for {
		if i > 10 {
			break
		}
		if n.RdyAll() {
			x.DstPut(n.Aux)
			y.DstPut(n.Aux)
			n.Aux = n.Aux.(int) + 1
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	n.Aux = float32(0)
	i = 0
	for {
		if i > 9 {
			break
		}
		if n.RdyAll() {
			x.DstPut(n.Aux)
			y.DstPut(n.Aux)
			n.Aux = n.Aux.(float32) + 1
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	n.Aux = []interface{}{uint64(math.MaxUint64), -1}
	i = 0
	for {
		if i > 0 {
			break
		}
		if n.RdyAll() {
			x.DstPut(n.Aux.([]interface{})[0])
			y.DstPut(n.Aux.([]interface{})[1])
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	n.Aux = []interface{}{int8(0), uint64(0)}
	i = 0
	for {
		if i > 0 {
			break
		}
		if n.RdyAll() {
			x.DstPut(n.Aux.([]interface{})[0])
			y.DstPut(n.Aux.([]interface{})[1])
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	n.Aux = []interface{}{int8(0), int16(0)}
	i = 0
	for {
		if i > 0 {
			break
		}
		if n.RdyAll() {
			x.DstPut(n.Aux.([]interface{})[0])
			y.DstPut(n.Aux.([]interface{})[1])
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	n.Aux = []interface{}{"Can you add an int to a string?", int8(77)}
	i = 0
	for {
		if i > 0 {
			break
		}
		if n.RdyAll() {
			x.DstPut(n.Aux.([]interface{})[0])
			y.DstPut(n.Aux.([]interface{})[1])
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	n.Aux = []interface{}{[4]complex128{0 + 0i, 0 + 0i, 0 + 0i, 0 + 0i}, int8(77)}
	i = 0
	for {
		if i > 0 {
			break
		}
		if n.RdyAll() {
			x.DstPut(n.Aux.([]interface{})[0])
			y.DstPut(n.Aux.([]interface{})[1])
			n.SendAll()
			i = i + 1
		}
		n.RecvOne()
	}

	// read all the acks to clean up
	for {
		n.RecvOne()
	}
	return nil

}

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil, nil)
	return node

}

func main() {

	fgbase.ConfigByFlag(nil)

	e, n := fgbase.MakeGraph(3, 3)

	n[0] = tbi(e[0], e[1])
	n[1] = fgbase.FuncAdd(e[0], e[1], e[2])
	n[2] = tbo(e[2])

	fgbase.RunAll(n)

}
