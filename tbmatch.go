package main

import (
	"github.com/vectaport/fgbase"
	"github.com/vectaport/fgbase/regexp"
)

var teststrings = []string{
	"test string",
	"test",
	"wrong",
	"testX",
}

func tbi(x fgbase.Edge) fgbase.Node {

	i := 0

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x},
		func(n *fgbase.Node) bool {
			return i <= len(teststrings) && n.DefaultRdyFunc()
		},
		func(n *fgbase.Node) error {
			if i < len(teststrings) {
				x.DstPut(regexp.Search{Orig: teststrings[i], Curr: teststrings[i], ID: regexp.NextID()})
			} else {
				if i == len(teststrings) {
					x.DstPut(regexp.Search{ID: regexp.NextID()})
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

	e, n := fgbase.MakeGraph(3, 3)

	e[2].Const("test")

	n[0] = tbi(e[0])
	n[1] = regexp.FuncMatch(e[0], e[2], e[1], false)
	n[2] = tbo(e[1])

	fgbase.RunAll(n)

}
