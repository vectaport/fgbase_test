package main

import (
	"flag"
	"time"

	"github.com/vectaport/fgbase"
)

func tbi(x fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil,
		func(n *fgbase.Node) {
			if n.Aux.(int)%3 == 0 {
				s := []int{0, 1, 2, 3, 4, 5, 6, 7}
				x.DstPut(s)
			} else if n.Aux.(int)%3 == 1 {
				x.DstPut(float32(n.Aux.(int)) + .5)
			} else {
				x.DstPut(n.Aux)
			}

			n.Aux = n.Aux.(int) + 1
			if n.Cnt%10000 == 0 {
				fgbase.StdoutLog.Printf("%2.f: %d (%.2f hz)\n", fgbase.TimeSinceStart(), n.Cnt, float64(n.Cnt)/fgbase.TimeSinceStart())

			}
		})
	node.Aux = 0
	return node

}

func main() {

	nodeid := flag.Int("nodeid", 0, "base for node ids")
	fgbase.ConfigByFlag(map[string]interface{}{"sec": 4})
	fgbase.NodeID = int64(*nodeid)

	fgbase.TraceSeconds = true

	time.Sleep(1 * time.Second)

	e, n := fgbase.MakeGraph(1, 1)

	n[0] = tbi(e[0])
	e[0].DstJSON(&n[0], "localhost:37777")

	fgbase.RunAll(n)

}
