package main

import (
	"flag"
	"math/rand"

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

	x.Ack = make(chan struct{}, fgbase.ChannelSize)
	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil,
		func(n *fgbase.Node) error {
			l := len((*x.Data)[0])
			if MaxChanLen < l {
				MaxChanLen = l
			}
			x.DstPut(n.NodeWrap(randSeq(16), x.Ack))
			return nil
		})
	return node
}

var reduceHz []float64
var reduceBase int64

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil,
		func(n *fgbase.Node) error {
			a.Flow = true
			if n.Cnt%1000 == 0 {
				reduceHz[n.ID-reduceBase] = float64(n.Cnt) / fgbase.TimeSinceStart()
			}
			return nil
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

func main() {

	nreducep := flag.Int("nreduce", 26, "number of reducers")
	nmapp := flag.Int("nmap", 6, "number of mappers")
	fgbase.ConfigByFlag(map[string]interface{}{"ncore": 3, "chansz": 32768, "trace": "Q", "sec": 4})
	nreduce := *nreducep
	nmap := *nmapp

	e, n := fgbase.MakeGraph(nmap+nreduce, nmap*2+nreduce)

	for i := 0; i < nmap; i++ {
		n[i] = tbi(e[i])
	}

	p := fgbase.FuncMap(e[0:nmap], e[nmap:nmap+nreduce], mapper)
	copy(n[nmap:2*nmap], p.Nodes())

	for i := 0; i < nreduce; i++ {
		n[2*nmap+i] = tbo(e[nmap+i])
	}

	reduceBase = int64(2 * nmap)
	reduceHz = make([]float64, nreduce)

	fgbase.RunAll(n)

	sum := 0.0
	for i := 0; i < len(reduceHz); i++ {
		sum += reduceHz[i]
	}
	if sum > 1000*1000 {
		fgbase.StdoutLog.Printf("%.2f Mhz\n", sum/1000/1000)
	} else {
		fgbase.StdoutLog.Printf("%.2f Khz\n", sum/1000)
	}
	fgbase.StdoutLog.Printf("(maxchansz=%d)\n", MaxChanLen)
}
