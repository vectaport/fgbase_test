package main

import (
	"flag"

	"github.com/vectaport/fgbase"
)

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil,
		func(n *fgbase.Node) {
			// time.Sleep(time.Duration(rand.Intn(150000)))
			if n.Cnt%10000 == 0 {
				fgbase.StdoutLog.Printf("%.2f: %d (=%.2f hz)\n", fgbase.TimeSinceStart(), n.Cnt, float64(n.Cnt)/fgbase.TimeSinceStart())
			}
		})
	return node
}

func main() {

	topicp := flag.String("topic", "test", "Kafka topic")
	fgbase.ConfigByFlag(map[string]interface{}{"trace": "Q", "chansz": 1024, "sec": 0, "ncore": 1})
	topic := *topicp

	fgbase.TraceSeconds = false

	e, n := fgbase.MakeGraph(1, 2)
	quitChan := make(chan struct{})

	n[0] = fgbase.FuncKcons(e[0], topic)
	n[1] = tbo(e[0])

	fgbase.RunAll(n)

	<-quitChan

}
