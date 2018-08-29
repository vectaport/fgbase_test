package main

import (
	"flag"
	//	"math/rand"
	"runtime"
	"strconv"
	//	"time"

	"github.com/vectaport/fgbase"
)

func main() {

	fgbase.ConfigByFlag(map[string]interface{}{"trace": "Q", "chansz": 1024, "sec": 0, "ncore": 1})

	fgbase.TraceLevel = fgbase.Q
	fgbase.TraceSeconds = false
	fgbase.ChannelSize = 1024

	e, n := fgbase.MakeGraph(1, 2)
	quitChan := make(chan struct{})

	n[0] = fgbase.FuncHTTP(e[0], ":8080", quitChan)
	n[1] = fgbase.FuncKprod(e[0])

	fgbase.RunAll(n)

	<-quitChan

}
