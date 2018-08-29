package main

import (
	"flag"

	"github.com/vectaport/fgbase"
)

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil,
		func(n *fgbase.Node) {
			a.Flow = true
			if n.Cnt%10000 == 0 {
				fgbase.StdoutLog.Printf("%2.f: %d (%.2f hz)\n", fgbase.TimeSinceStart(), n.Cnt, float64(n.Cnt)/fgbase.TimeSinceStart())
			}
		})

	return node
}

func main() {

	nodeidp := flag.Int("nodeid", 0, "base for node ids")
	fgbase.ConfigByFlag(map[string]interface{}{"sec": 4})
	fgbase.NodeID = int64(*nodeidp)

	fgbase.TraceSeconds = true

	e, n := fgbase.MakeGraph(1, 1)

	n[0] = tbo(e[0])
	e[0].SrcJSON(&n[0], "localhost:37777")

	fgbase.RunAll(n)

}
