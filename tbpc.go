package main

import (
	"github.com/vectaport/flowgraph"
)

var pcHz []float64
var pcBase int64 = 0

func pc(incFlag,addrIn,absFlag,addrOut flowgraph.Edge) flowgraph.Node {

	node := flowgraph.FuncFunc("pc", []*flowgraph.Edge{&incFlag,&absFlag,&addrIn}, []*flowgraph.Edge{&addrOut}, 

		func (n *flowgraph.Node) bool {
			incFlag := n.Srcs[0]
			absFlag := n.Srcs[1]
			addrIn := n.Srcs[2]
			addrOut := n.Dsts[0]
			if addrOut.DstRdy(n) {
				if incFlag.SrcRdy(n) {
					absFlag.NoOut = true
					addrIn.NoOut = true
					return addrOut.DstRdy(n)
				}
				if addrIn.SrcRdy(n) && absFlag.SrcRdy(n) {
					incFlag.NoOut = true
					return addrOut.DstRdy(n)
				}			
			}
			return false
		},
		func (n *flowgraph.Node) { 
			addrOut := n.Dsts[0]
			addrOut.Val = n.Aux.(int)
			
			n.Aux = addrOut.Val.(int)+1
			if n.Cnt%100==0 {
				pcHz[n.ID-pcBase] = float64(n.Cnt)/flowgraph.TimeSinceStart()
			}})
	node.Aux = 0
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {
	
	
	flowgraph.ConfigByFlag(map[string]interface{}{ "ncore":4, "trace":"Q", "sec":4, "trsec":true})


	e,n := flowgraph.MakeGraph(4, 2)

	e[0].Const(true)
	e[0].Name = "incFlag"
        e[1].Name = "addrIn"
	e[2].Name = "absFlag"
	e[3].Name = "addrOut"

	n[0] = pc(e[0], e[1], e[2], e[3])
	n[1] = tbo(e[3])

	pcHz = make([]float64, 1)

	flowgraph.RunAll(n)

	// generate total frequency for pc
	sum := 0.0
	for i:=0; i<len(pcHz); i++ {
		sum += pcHz[i]
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

