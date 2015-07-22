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

var MaxChanLen = 0

func tbi(x flowgraph.Edge) flowgraph.Node {

	x.Ack = make(chan flowgraph.Nada, flowgraph.ChannelSize)
	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			l:=len((*x.Data)[0])
			if MaxChanLen < l {
				MaxChanLen = l
			}
			x.Val = n.NodeWrap(randSeq(16), x.Ack)})
	return node
}

var reduceHz []float64
var reduceBase int64

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, 
		func (n *flowgraph.Node) { 
			if n.Cnt%1000==0 {
				reduceHz[n.ID-reduceBase] = float64(n.Cnt)/flowgraph.TimeSinceStart()
			}})
	return node
}

func main() {

	
	nreducep := flag.Int("nreduce", 4, "number of reducers")
	nmapp := flag.Int("nmap", 2, "number of mappers")
	flowgraph.ConfigByFlag(nil)
	nreduce := *nreducep
	nmap := *nmapp

	e,n := flowgraph.MakeGraph(nmap+nreduce,nmap*2+nreduce)


	for i:= 0; i<nmap; i++ {
		n[i] = tbi(e[i])
	}

	var mapSel = func(n *flowgraph.Node) int {
		i,ok := n.Srcs[0].Val.(int)
		if ok {return i%nreduce}
		s,ok := n.Srcs[0].Val.(string)
		if ok { 
			return int(s[0]-'A')*nreduce/26
		}
		return -1
	}
	p := flowgraph.FuncMap(e[0:nmap], e[nmap:nmap+nreduce], mapSel)
	copy(n[nmap:2*nmap], p.Nodes())

	for i:= 0; i<nreduce; i++ {
		n[2*nmap+i] = tbo(e[nmap+i])
	}

	reduceBase = int64(2*nmap)
	reduceHz = make([]float64, nreduce)

	flowgraph.RunAll(n)

	sum := 0.0
	for i:=0; i<len(reduceHz); i++ {
		sum += reduceHz[i]
	}
	if sum>1000*1000 {
		flowgraph.StdoutLog.Printf("%.2f Mhz\n", sum/1000/1000)
	} else {
		flowgraph.StdoutLog.Printf("%.2f Khz\n", sum/1000)
	}
	flowgraph.StdoutLog.Printf("(maxchansz=%d)\n", MaxChanLen)
}

