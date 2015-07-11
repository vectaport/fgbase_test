package main

import (
	"flag"
	"math/rand"
	"runtime"
	"strconv"
	"time"

	"github.com/vectaport/flowgraph"
)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			x.Val = randSeq(16)
		})
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	tracep := flag.String("trace", "V", "trace level, Q|V|VV|VVV|VVVV")
	nCorep := flag.Int("ncore", runtime.NumCPU()-1, "num cores to use, max is "+strconv.Itoa(runtime.NumCPU()))
	nforkp := flag.Int("nfork", 4, "number of forks")
	flag.Parse()
	flowgraph.TraceLevel = flowgraph.TraceLevels[*tracep]
	flowgraph.TraceSeconds = true
	nfork := *nforkp
	runtime.GOMAXPROCS(*nCorep)

	e,n := flowgraph.MakeGraph(1+nfork,1+nfork*2)

	n[0] = tbi(e[0])

	var forkSel = func(n *flowgraph.Node) int {
		i,ok := n.Srcs[0].Val.(int)
		if ok {return i%nfork}
		s,ok := n.Srcs[0].Val.(string)
		if ok { 
			return int(s[0]-'A')*nfork/26
		}
		return -1
	}
	p := flowgraph.FuncFork(e[0], e[1:1+nfork], forkSel)
	copy(n[1:1+nfork], p.Nodes())

	for i:= 0; i<nfork; i++ {
		n[1+nfork+i] = tbo(e[1+i])
	}

	flowgraph.RunAll(n, 60*time.Second)
}

