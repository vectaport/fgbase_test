package main

import (
	"flag"
	"math/rand"
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

var MaxChanLen = 0

func tbi(x flowgraph.Edge) flowgraph.Node {

	x.Ack = make(chan flowgraph.Nada, flowgraph.ChannelSize)
	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			time.Sleep(time.Second/100000)
			l:=len((*x.Data)[0])
			if MaxChanLen < l {
				MaxChanLen = l
			}
			x.Val = n.NodeWrap(randSeq(16), x.Ack)})
	return node
}

var tboHz []float64
var tboBase int64

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, 
		func (n *flowgraph.Node) { 
			if n.Cnt%100==0 {
				tboHz[n.ID-tboBase] = float64(n.Cnt)/flowgraph.TimeSinceStart()
			}})
	return node
}

func testOrder(n *flowgraph.Node, dict []string) {
	for i := 0; i<len(dict)-1; i++ {
		if dict[i]>=dict[i+1] {
			n.LogError("out of order %v\n", dict)
		}
	}
}

func reducer(n *flowgraph.Node, s, d flowgraph.Datum) flowgraph.Datum {
	dict := d.([]string)
	ss := s.(string)
	lo := 0
	hi := len(dict)-1
	for lo<=hi {
		mid := (lo+hi)/2
		if ss < dict[mid] {
			hi = mid-1
		} else if ss > dict[mid] {
			lo = mid+1
		} else {
			break	
		}
	}
	i := lo
	
	if i==len(dict) {
		dict = append(dict, ss)
	} else {
		dict = append(dict, "")
		copy(dict[i+1:], dict[i:])
		dict[i] = ss

		testOrder(n, dict)
	}
	return dict
	
}


func main() {

	
	nreducep := flag.Int("nreduce", 26, "number of reducers")
	nmapp := flag.Int("nmap", 128, "number of mappers")
	flowgraph.ConfigByFlag(map[string]interface{}{ "ncore":6, "chansz":128, "trace": "Q", "sec":4})
	nreduce := *nreducep
	nmap := *nmapp

	e,n := flowgraph.MakeGraph(nmap+nreduce*2,nmap*2+nreduce*2)


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
		n[2*nmap+i] = flowgraph.FuncReduce(e[nmap+i], e[nmap+nreduce+i], reducer)
	}
	for i:= 0; i<nreduce; i++ {
		n[2*nmap+nreduce+i] = tbo(e[nmap+nreduce+i])
	}

	tboBase = int64(2*nmap+nreduce)
	tboHz = make([]float64, nreduce)

	flowgraph.RunAll(n)

	sum := 0.0
	for i:=0; i<len(tboHz); i++ {
		sum += tboHz[i]
	}
	if sum>1000*1000 {
		flowgraph.StdoutLog.Printf("%.2f Mhz\n", sum/1000/1000)
	} else {
		flowgraph.StdoutLog.Printf("%.2f Khz\n", sum/1000)
	}
	flowgraph.StdoutLog.Printf("(maxchansz=%d)\n", MaxChanLen)
}

