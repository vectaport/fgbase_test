package main

import (
	"github.com/vectaport/fgbase"
	"github.com/vectaport/fgbase/regexp"
)

var teststrings = []string{
	/*
		"MMMapplesapplesapplesTAG",
	*/
	"applesapplesapples",
	/*
	           "applesX",
	           "orangesX",
	   	"apples",
	   	"oranges",
	   	"T",
	   	"",
	*/
}

func tbi(dnstreq fgbase.Edge, newmatch fgbase.Edge) fgbase.Node {

	i := 0

	Prev := make(map[string]string)

	node := fgbase.MakeNode("tbi", []*fgbase.Edge{&dnstreq}, []*fgbase.Edge{&newmatch},
		func(n *fgbase.Node) bool {
			return dnstreq.SrcRdy(n) || newmatch.DstRdy(n) && i < len(teststrings)
		},
		func(n *fgbase.Node) error {
			if dnstreq.SrcRdy(n) {
				match := dnstreq.SrcGet().(regexp.Search)
				if len(Prev[match.Orig]) > 1 {
					match.Curr = Prev[match.Orig][1:]
					Prev[match.Orig] = match.Curr
					newmatch.DstPut(match)
				}
				delete(Prev, match.Orig)
				return nil
			}
			if i < len(teststrings) {
				newmatch.DstPut(regexp.Search{Orig: teststrings[i], Curr: teststrings[i], State: regexp.Live, ID: regexp.NextID()})
			} else {
				if i == len(teststrings) {
					newmatch.DstPut(regexp.Search{ID: regexp.NextID()})
				}
			}
			i++
			return nil
		})
	return node

}

func tbo(oldmatch fgbase.Edge, dnstreq fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&oldmatch}, []*fgbase.Edge{&dnstreq}, nil,
		func(n *fgbase.Node) error {
			match := oldmatch.SrcGet().(regexp.Search)
			match.State = regexp.Done
			dnstreq.DstPut(match)
			return nil
		})
	return node

}

type edgeCnt int

const (
	newmatch edgeCnt = iota
	subsrc
	dnstreq
	oldmatch
	subdst
	upstreq
	apples
	edgeNum
)

var edgeNames []string = []string{
	"newmatch",
	"subsrc",
	"dnstreq",
	"oldmatch",
	"subdst",
	"upstreq",
	"apples",
}

func main() {

	fgbase.ConfigByFlag(nil)

	e, n := fgbase.MakeGraph(int(edgeNum), 4)
	fgbase.NameEdges(e, edgeNames)

	e[apples].Const("apples")

	n[0] = tbi(e[upstreq], e[newmatch])
	n[1] = regexp.FuncRepeat(e[newmatch], e[subsrc], e[dnstreq], e[oldmatch], e[subdst], e[upstreq], 2, -1)
	n[2] = regexp.FuncMatch(e[subdst], e[apples], e[subsrc], false)
	n[3] = tbo(e[oldmatch], e[dnstreq])

	fgbase.RunAll(n)

}
