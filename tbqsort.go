package main

import (
	"flag"
	"math/rand"
	"runtime"
	"sort"
	"sync/atomic"
	"time"

	"github.com/vectaport/flowgraph"
)

var bushelCnt int64

type bushel struct {
	depth int64
	bushelID int64
	Original []int
	Sliced []int
}

// borrowed from Golang 1.4.2 sort example, copyright notice in flowgraph/GO-LICENSE
func (a bushel) Len() int           { return len(a.Sliced) }
func (a bushel) Swap(i, j int)      { a.Sliced[i], a.Sliced[j] = a.Sliced[j], a.Sliced[i] }
func (a bushel) Less(i, j int) bool { return a.Sliced[i] < a.Sliced[j] }

func (a bushel) Sorted() bool {
	l := len(a.Sliced)
	for i:= 0; i<l-1; i++ {
		if a.Sliced[i] > a.Sliced[i+1] {
			return false
	}
	}
	return true
}

func (a bushel) SubSlice(n, m int) flowgraph.Datum {
	a.Sliced = a.Sliced[n:m]
	a.depth += 1
	return a
}

func (a bushel) OrigSorted() bool {
	l := len(a.Original)
	for i:= 0; i<l-1; i++ {
		if a.Original[i] > a.Original[i+1] {
			return false
		}
	}
	return true
}

func (a bushel) Orig() []int {
	return a.Original
}

func (a bushel) Slic() []int {
	return a.Sliced
}

func (a bushel) Depth() int64 { 
	return a.depth
}

func (a *bushel) DepthIncr() { 
	
	a.depth += 1
}

func (a bushel) ID() int64 {
	return a.bushelID
}

func tbiRand() sort.Interface {
	var s bushel
	s.bushelID = atomic.AddInt64(&bushelCnt, 1)-1
	n := 1024*1024
	l := rand.Intn(n)
	for i:=0; i<l; i++ {
		s.Original = append(s.Original, rand.Intn(n))
	}
	s.Sliced = s.Original
	return s
}

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil,
		func(n *flowgraph.Node) { n.Dsts[0].Val = tbiRand() })
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, 
		func(n *flowgraph.Node) {
			switch v := a.Val.(type) {
			case flowgraph.Interface2: {
				n.Tracef("Original(%p) sorted %t, Sliced sorted %t, depth=%d, id=%d, len=%d\n", v.Orig(), v.OrigSorted(), v.Sorted(), v.Depth(), v.ID(), v.Len())
			}
			default: {
				n.Tracef("not of type flowgraph.Interface2\n")
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

	flowgraph.TraceLevel = flowgraph.V

	e,n := flowgraph.MakeGraph(2, poolSz+2)

	n[0] = tbi(e[0])
	n[1] = tbo(e[1])

	p := n[2:poolSz+2]
	copy(p, flowgraph.FuncQsort(e[0], e[1], poolSz))

	flowgraph.RunAll(n, time.Duration(sec)*time.Second)

}
