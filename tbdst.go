package main

import (
	"flag"
	"net"
	"time"

	"github.com/vectaport/fgbase"
)

func tbi(x fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil,
		func(n *fgbase.Node) error {
			x.DstPut(n.Aux)
			n.Aux = n.Aux.(int) + 1
			return nil
		})

	node.Aux = 0
	return node

}

func main() {

	nodeid := flag.Int("nodeid", 0, "base for node ids")
	fgbase.ConfigByFlag(map[string]interface{}{"sec": 2})
	fgbase.NodeID = int64(*nodeid)

	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "localhost:37777")
	if err != nil {
		fgbase.StderrLog.Printf("%v\n", err)
		return
	}

	e, n := fgbase.MakeGraph(1, 2)

	n[0] = tbi(e[0])
	n[1] = fgbase.FuncDst(e[0], conn)

	fgbase.RunAll(n)

}
