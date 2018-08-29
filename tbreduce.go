package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/vectaport/fgbase"
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

func tbi(x fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil,
		func(n *fgbase.Node) {
			l := len((*x.Data)[0])
			if MaxChanLen < l {
				MaxChanLen = l
			}
			if n.Cnt%100 == 0 {
				tbiHz[n.ID-tbiBase] = float64(n.Cnt) / fgbase.TimeSinceStart()
			}
			x.DstPut(n.NodeWrap(randSeq(16), x.Ack))
		})
	return node
}

var tbiHz []float64
var tbiBase int64

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil,
		func(n *fgbase.Node) {
			a.Flow = true
			time.Sleep(100000000)
		})
	return node
}

func mapper(n *fgbase.Node, datum interface{}) int {
	nreduce := len(n.Dsts)
	i, ok := datum.(int)
	if ok {
		return i % nreduce
	}
	s, ok := datum.(string)
	if ok {
		return int(s[0]-'A') * nreduce / 26
	}
	return -1
}

func testOrder(n *fgbase.Node, dict []string) {
	for i := 0; i < len(dict)-1; i++ {
		if dict[i] >= dict[i+1] {
			n.LogError("out of order %v\n", dict)
		}
	}
}

func reducer(n *fgbase.Node, datum, collection interface{}) interface{} {
	s := datum.(string)
	c := collection.([]string)
	lo := 0
	hi := len(c) - 1
	for lo <= hi {
		mid := (lo + hi) / 2
		if s < c[mid] {
			hi = mid - 1
		} else if s > c[mid] {
			lo = mid + 1
		} else {
			lo = mid
			break
		}
	}
	i := lo

	c = append(c, s)
	if i < len(c)-1 {
		copy(c[i+1:], c[i:])
		c[i] = s
	}

	testOrder(n, c)

	return c

}

func main() {

	nreducep := flag.Int("nreduce", 26, "number of reducers")
	nmapp := flag.Int("nmap", 4, "number of mappers")
	fgbase.ConfigByFlag(map[string]interface{}{"ncore": 4, "trace": "Q", "sec": 4})
	nreduce := *nreducep
	nmap := *nmapp

	e, n := fgbase.MakeGraph(nmap+nreduce*2, nmap*2+nreduce*2)

	for i := 0; i < nmap; i++ {
		n[i] = tbi(e[i])
	}

	p := fgbase.FuncMap(e[0:nmap], e[nmap:nmap+nreduce], mapper)
	copy(n[nmap:2*nmap], p.Nodes())

	for i := 0; i < nreduce; i++ {
		n[2*nmap+i] = fgbase.FuncReduce(e[nmap+i], e[nmap+nreduce+i], reducer, true)
	}
	for i := 0; i < nreduce; i++ {
		n[2*nmap+nreduce+i] = tbo(e[nmap+nreduce+i])
	}

	// tboBase = int64(2*nmap+nreduce)
	// tboHz = make([]float64, nreduce)
	tbiBase = int64(0)
	tbiHz = make([]float64, nmap)

	fgbase.RunAll(n)

	sum := 0.0
	for i := 0; i < len(tbiHz); i++ {
		sum += tbiHz[i]
	}

	speed := sum / 1000
	hzstr := "Khz\n"
	if sum > 1000*1000 {
		speed = speed / 1000
		hzstr = "Mhz\n"
	}
	if fgbase.TraceLevel == fgbase.QQ {
		hzstr = ""
	}
	fgbase.StdoutLog.Printf("%.2f%s", speed, hzstr)
	// fgbase.StdoutLog.Printf("(maxchansz=%d)\n", MaxChanLen)
}
