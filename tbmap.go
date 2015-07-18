package main

import (
	"flag"
	"math/rand"

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

	x.Ack = make(chan flowgraph.Nada, flowgraph.ChannelSize)
	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			x.Val = n.NodeWrap(randSeq(16), x.Ack)
			if n.Cnt%1000==0 {
				flowgraph.StdoutLog.Printf("%.2f: %d (%.2f Khz)\n", flowgraph.TimeSinceStart(), n.Cnt, float64(n.Cnt)/flowgraph.TimeSinceStart()/1000)
			}})
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	
	nmapp := flag.Int("nmap", 4, "number of mappers")
	nfeedp := flag.Int("nfeed", 2, "number of feeds")
	flowgraph.ConfigByFlag(nil)
	nmap := *nmapp
	nfeed := *nfeedp

	e,n := flowgraph.MakeGraph(1+nmap,nfeed*2+nmap)


	for i:= 0; i<nfeed; i++ {
		n[i] = tbi(e[0])
	}

	var mapSel = func(n *flowgraph.Node) int {
		i,ok := n.Srcs[0].Val.(int)
		if ok {return i%nmap}
		s,ok := n.Srcs[0].Val.(string)
		if ok { 
			return int(s[0]-'A')*nmap/26
		}
		return -1
	}
	p := flowgraph.FuncMap(e[0], e[1:1+nmap], nfeed, mapSel)
	copy(n[nfeed:2*nfeed], p.Nodes())

	for i:= 0; i<nmap; i++ {
		n[2*nfeed+i] = tbo(e[1+i])
	}

	flowgraph.RunAll(n)
}

