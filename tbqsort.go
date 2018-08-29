package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"

	"github.com/vectaport/fgbase"
)

var bushelCnt int64
var qsortPool fgbase.Pool

type bushel struct {
	Slic     []int
	Orig     []int
	depth    int64
	bushelID int64
}

// borrowed from Golang 1.4.2 sort example, copyright notice in fgbase/GO-LICENSE
func (a bushel) Len() int           { return len(a.Slic) }
func (a bushel) Swap(i, j int)      { a.Slic[i], a.Slic[j] = a.Slic[j], a.Slic[i] }
func (a bushel) Less(i, j int) bool { return a.Slic[i] < a.Slic[j] }

func (a bushel) SubSlice(n, m int) interface{} {
	a.Slic = a.Slic[n:m]
	a.depth += 1
	return a
}

func (a bushel) Slice() []int {
	return a.Slic
}

func (a bushel) Original() []int {
	return a.Orig
}

func (a bushel) Depth() int64 {
	return a.depth
}

func (a bushel) ID() int64 {
	return a.bushelID
}

func tbiRand(pow2 uint) fgbase.RecursiveSort {
	var s bushel
	s.bushelID = bushelCnt
	bushelCnt += 1
	n := rand.Intn(1<<pow2) + 1
	l := rand.Intn(n)
	for i := 0; i < l; i++ {
		s.Orig = append(s.Orig, rand.Intn(l))
	}
	s.Slic = s.Orig
	fmt.Printf("RAND\n")
	return s
}

func tbi(x fgbase.Edge, pow2 uint) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil,
		func(n *fgbase.Node) { x.DstPut(tbiRand(pow2)) })
	return node
}

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil,
		func(n *fgbase.Node) {
			av := a.SrcGet()
			switch v := av.(type) {
			case fgbase.RecursiveSort:
				{
					if sort.IntsAreSorted(v.Original()) {
						n.Tracef("END for id=%d, depth=%d, len=%d\n", v.ID(), v.Depth(), v.Len())
					}
					n.Tracef("Original(%p) sorted %t, Slice sorted %t, depth=%d, id=%d, len=%d, poolsz=%d, ratio = %d\n", v.Original(), sort.IntsAreSorted(v.Original()), sort.IntsAreSorted(v.Slice()), v.Depth(), v.ID(), len(v.Original()), qsortPool.Size(), len(v.Original())/(1+int(v.Depth())))
				}
			default:
				{
					n.Tracef("not of type fgbase.RecursiveSort\n")
				}
			}
		})
	return node
}

func main() {

	poolSzp := flag.Int("poolsz", 64, "qsort pool size")
	pow2p := flag.Uint("pow2", 20, "power of 2 to scale random numbers")
	fgbase.ConfigByFlag(nil)
	poolSz := *poolSzp
	pow2 := *pow2p

	e, n := fgbase.MakeGraph(2, poolSz+2)

	n[0] = tbi(e[0], pow2)
	n[1] = tbo(e[1])

	p := fgbase.FuncQsort(e[0], e[1], poolSz)
	copy(n[2:poolSz+2], p.Nodes())
	p.Alloc(&n[2], 1) // reserve one for input

	fgbase.RunAll(n)

}
