package main

import (
	"flag"
	"math/rand"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/vectaport/flowgraph"
)

var bushelCnt int64

type bushel struct {
	Slic []int
	Orig []int
	depth int64
	bushelID int64
}

// borrowed from Golang 1.4.2 sort example, copyright notice in flowgraph/GO-LICENSE
func (a bushel) Len() int           { return len(a.Slic) }
func (a bushel) Swap(i, j int)      { a.Slic[i], a.Slic[j] = a.Slic[j], a.Slic[i] }
func (a bushel) Less(i, j int) bool { return a.Slic[i] < a.Slic[j] }

func (a bushel) SubSlice(n, m int) flowgraph.Datum {
	a.Slic = a.Slic[n:m]
	a.depth += 1
	return a
}

func (a bushel) Slice() []int {
	return a.Slic
}

func (a bushel) SliceSorted() bool {
	l := len(a.Slic)
	for i:= 0; i<l-1; i++ {
		if a.Slic[i] > a.Slic[i+1] {
			return false
	}
	}
	return true
}

func (a bushel) Original() []int {
	return a.Orig
}

func (a bushel) OriginalSorted() bool {
	l := len(a.Orig)
	for i:= 0; i<l-1; i++ {
		if a.Orig[i] > a.Orig[i+1] {
			return false
		}
	}
	return true
}

func (a bushel) Depth() int64 { 
	return a.depth
}

func (a bushel) ID() int64 {
	return a.bushelID
}

func tbiRand() flowgraph.RecursiveSort {
	var s bushel
	s.bushelID = atomic.AddInt64(&bushelCnt, 1)-1
	n := rand.Intn(1<<20)
	l := rand.Intn(n)
	for i:=0; i<l; i++ {
		s.Orig = append(s.Orig, rand.Intn(l))
	}
	s.Slic = s.Orig
	return s
}

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil,
		func(n *flowgraph.Node) { x.Val = tbiRand() })
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, 
		func(n *flowgraph.Node) {
			switch v := a.Val.(type) {
			case flowgraph.RecursiveSort: {
				n.Tracef("Original(%p) sorted %t, Slice sorted %t, depth=%d, id=%d, len=%d\n", v.Original(), v.OriginalSorted(), v.SliceSorted(), v.Depth(), v.ID(), v.Len())
			}
			default: {
				n.Tracef("not of type flowgraph.RecursiveSort\n")
			}
			}})
	return node
}

func main() {

	poolSzp := flag.Int("poolsz", 64, "qsort pool size")
	numCorep := flag.Int("numcore", 1, "num cores to use")
	secp := flag.Int("sec", 1, "seconds to run")
	flag.Parse()
	poolSz := *poolSzp
	runtime.GOMAXPROCS(*numCorep)
	sec := *secp

	flowgraph.TraceLevel = flowgraph.VV

	e,n := flowgraph.MakeGraph(2, poolSz+2)

	n[0] = tbi(e[0])
	n[1] = tbo(e[1])

	p := n[2:poolSz+2]
	copy(p, flowgraph.FuncQsort(e[0], e[1], poolSz))

	flowgraph.RunAll(n, time.Duration(sec)*time.Second)

}
