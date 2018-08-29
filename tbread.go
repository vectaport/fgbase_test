package main

import (
	"flag"
	"os"

	"github.com/vectaport/fgbase"
)

func tbo(a fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil, nil)
	return node
}

func check(e error) {
	if e != nil {
		fgbase.StderrLog.Printf("%v\n", e)
		os.Exit(1)
	}
}

func main() {

	fgbase.ConfigByFlag(map[string]interface{}{"sec": 2})
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	fileName := flag.Arg(0)

	f, err := os.Open(fileName)
	check(err)

	e, n := fgbase.MakeGraph(1, 2)

	n[0] = fgbase.FuncRead(e[0], f)
	n[1] = tbo(e[0])

	fgbase.RunAll(n)

}
