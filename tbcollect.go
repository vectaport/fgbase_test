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

var tbiHz []float64
var tbiBase int64 = 0

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			l:=len((*x.Data)[0])
			if MaxChanLen < l {
				MaxChanLen = l
			}
			x.DstPut(n.NodeWrap(randSeq(16), x.Ack))
			if n.Cnt%100==0 {
				tbiHz[n.ID-tbiBase] = float64(n.Cnt)/flowgraph.TimeSinceStart()
			}})
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func tbc(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbc", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			time.Sleep(10000000)
			x.DstPut(true)
		})
	return node
}

func mapper(n *flowgraph.Node, datum interface{}) int {
	nreduce := len(n.Dsts)
	i,ok := datum.(int)
	if ok {return i%nreduce}
	s,ok := datum.(string)
	if ok { 
		return int(s[0]-'A')*nreduce/26
	}
	return -1
}

func testOrder(n *flowgraph.Node, dict []string) {
	for i := 0; i<len(dict)-1; i++ {
		if dict[i]>dict[i+1] {
			n.LogError("out of order %v\n", dict)
		}
		if dict[i]==dict[i+1] {
			n.Tracef("WARNING:  duplicate string found %v\n", dict[i])
		}
	}
}

func reducer(n *flowgraph.Node, datum,collection interface{}) interface{} {
	c,ok := collection.([]string)
	if !ok {
		c = []string{}
	}
	s,ok := datum.(string)
	if !ok {
		return c
	}
	lo := 0
	hi := len(c)-1
	for lo<=hi {
		mid := (lo+hi)/2
		if s < c[mid] {
			hi = mid-1
		} else if s > c[mid] {
			lo = mid+1
		} else {
			lo = mid
			break	
		}
	}
	i := lo

	c = append(c, s)
	if i<len(c)-1 {
		copy(c[i+1:], c[i:])
		c[i] = s
	}
	
	testOrder(n, c)

	return c
	
}

const (
	rdcRdy = iota
	cllRdy 
)

func reduce2(rdc,cll,snd flowgraph.Edge, reducer func(n *flowgraph.Node, datum,collection interface{}) interface{}) flowgraph.Node {

	var rdyFunc = func(n *flowgraph.Node) bool {
		lastRdy := n.RdyState
		if lastRdy!=rdcRdy && rdc.SrcRdy(n) {
			n.RdyState = rdcRdy
			return true
		}
		if cll.SrcRdy(n) && snd.DstRdy(n) {
			n.RdyState = cllRdy
			return true
		}
		if rdc.SrcRdy(n) {
			n.RdyState = rdcRdy
			return true
		}
		return false
	}

	var fireFunc = func(n *flowgraph.Node) {
		c := n.Aux
		lastRdy := n.RdyState
		if lastRdy == rdcRdy {
			n.Aux = reducer(n, rdc.SrcGet(), c)
			return
		}
		snd.DstPut(c)
		return
	}

	node := flowgraph.MakeNode("reduce2", []*flowgraph.Edge{&rdc,&cll}, []*flowgraph.Edge{&snd}, rdyFunc, fireFunc)
	return node
}

func main() {
	
	nreducep := flag.Int("nreduce", 26, "number of reducers")
	nmapp := flag.Int("nmap", 4, "number of mappers")
	flowgraph.ConfigByFlag(map[string]interface{}{ "ncore":4, "trace":"Q", "sec":4, "trsec":true})
	nreduce := *nreducep
	nmap := *nmapp

	e,n := flowgraph.MakeGraph(nmap+nreduce*3,nmap*2+nreduce*3)

	tboNodeBase := 2*nmap+nreduce
	tboEdgeBase := nmap+nreduce
	tbcNodeBase := tboNodeBase+nreduce
	tbcEdgeBase := tboEdgeBase+nreduce

	for i:= 0; i<nmap; i++ {
		n[i] = tbi(e[i])
	}

	p := flowgraph.FuncMap(e[0:nmap], e[nmap:nmap+nreduce], mapper)
	copy(n[nmap:2*nmap], p.Nodes())
	
	for i:= 0; i<nreduce; i++ {
		n[2*nmap+i] = reduce2(e[nmap+i],e[tbcEdgeBase+i], e[tboEdgeBase+i], reducer)
	}

	for i:= 0; i<nreduce; i++ {
		n[tboNodeBase+i] = tbo(e[tboEdgeBase+i])
	}

	for i:= 0; i<nreduce; i++ {
		n[tbcNodeBase+i] = tbc(e[tbcEdgeBase+i])
	}

	tbiHz = make([]float64, nreduce)

	flowgraph.RunAll(n)

	// generate total frequency for tbi
	sum := 0.0
	for i:=0; i<len(tbiHz); i++ {
		sum += tbiHz[i]
	}

	speed := sum/1000
	hzstr := "Khz\n"
	if sum>1000*1000 {
		speed = speed/1000
		hzstr = "Mhz\n"
	}
	if flowgraph.TraceLevel==flowgraph.QQ {
		hzstr = ""
	}
	flowgraph.StdoutLog.Printf("%.2f%s", speed, hzstr)
}

