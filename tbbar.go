package main

import (
	"github.com/vectaport/fgbase"
	"github.com/vectaport/fgbase/regexp"
)

var teststrings = []string{
	"test string",
	/*
	   	"apples",
	   	"oranges",
	   	"",
	   	"T",
	           "applesX",
	           "orangesX",
	*/
}

func tbi(x fgbase.Edge) fgbase.Node {

	i := 0

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x},
		func(n *fgbase.Node) bool {
			return i <= len(teststrings) && n.DefaultRdyFunc()
		},
		func(n *fgbase.Node) error {
			if i < len(teststrings) {
				x.DstPut(regexp.Search{Curr: teststrings[i], Orig: teststrings[i]})
			} else {
				if i == len(teststrings) {
					x.DstPut(regexp.Search{})
				}
			}
			i++
			return nil
		})
	return node

}

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil, nil)
	return node

}

func main() {

	fgbase.ConfigByFlag(nil)

	e, n := fgbase.MakeGraph(6, 5)

	e[4].Const("apples")
	e[5].Const("oranges")

	n[0] = tbi(e[0])
	n[1] = regexp.FuncMatch(e[0], e[4], e[1], false)
	n[2] = regexp.FuncMatch(e[0], e[5], e[2], false)
	n[3] = regexp.FuncBar(e[1], e[2], e[3], true)
	n[4] = tbo(e[3])

	fgbase.RunAll(n)

}
