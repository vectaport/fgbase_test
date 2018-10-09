package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/vectaport/fgbase"
	"github.com/vectaport/fgbase/weblab"
)

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil,
		func(n *fgbase.Node) error {
			// time.Sleep(time.Duration(rand.Intn(150000)))
			a.Flow = true
			if n.Cnt%10000 == 0 {
				fgbase.StdoutLog.Printf("%.2f: %d (%.2f hz)\n", fgbase.TimeSinceStart(), n.Cnt, float64(n.Cnt)/fgbase.TimeSinceStart())
			}
			return nil
		})
	return node
}

func main() {

	nPortp := flag.Int("nport", 1, "number of server ports")
	testp := flag.Bool("test", false, "test mode")
	fgbase.ConfigByFlag(map[string]interface{}{"trace": "Q", "chansz": 1024})
	nPort := *nPortp
	test := *testp

	fgbase.TraceSeconds = true

	var quitChan chan struct{}
	if !test {
		quitChan = make(chan struct{})
		fgbase.RunTime = 0
	} else {
		fgbase.RunTime = 10 * time.Second
	}

	e, n := fgbase.MakeGraph(1, nPort+1)

	for i := 0; i < nPort; i++ {
		n[i] = weblab.FuncHTTP(e[0], fmt.Sprintf(":%d", 8080+i), quitChan)
	}

	n[nPort] = tbo(e[0])

	fgbase.RunAll(n)

	if !test {
		<-quitChan
	}
}
