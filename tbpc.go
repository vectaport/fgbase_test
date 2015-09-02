package main

import (
	"github.com/vectaport/flowgraph"
)

var pcHz []float64
var pcBase int64 = 1

// 
const (
	inc = iota
	rel
	abs
)

var path = "pc.csv"

func tbi(incRail,relRail,absRail,addrIn flowgraph.Edge) flowgraph.Node {

	var nada flowgraph.Nada

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&incRail,&relRail,&absRail,&addrIn}, nil,
		func (n *flowgraph.Node) { 
			incRail.Val = nada
			relRail.NoOut = true
			absRail.NoOut = true
			addrIn.NoOut = true
		})
	return node
}

func pc(incRail,relRail,absRail,addrIn,addrOut flowgraph.Edge) flowgraph.Node {

	var rdyFunc = func (n *flowgraph.Node) bool {
		if addrOut.DstRdy(n) {
			if incRail.SrcRdy(n) {
				n.RdyState = inc
				relRail.NoOut = true
				absRail.NoOut = true
				addrIn.NoOut = true
				return true
			}
			if addrIn.SrcRdy(n) { 
				if relRail.SrcRdy(n) {
					n.RdyState = rel
					incRail.NoOut = true
					absRail.NoOut = true
					return true
				}
				if absRail.SrcRdy(n) {
					n.RdyState = abs
					incRail.NoOut = true
					relRail.NoOut = true
					return true
				}
			}			
		}
		return false
	}

	var fireFunc = func (n *flowgraph.Node) { 
		if n.RdyState==abs {
			p := addrIn.Val.(int)
			addrOut.Val = p
			n.Aux = p
		} else {
			addrOut.Val = n.Aux.(int)
			if n.RdyState==inc {
				n.Aux = addrOut.Val.(int)+1
			} else {
				n.Aux = addrOut.Val.(int)+addrIn.Val.(int)
			}
		}
		
		if n.Cnt%1024==0 {
			pcHz[n.ID-pcBase] = float64(n.Cnt)/flowgraph.TimeSinceStart()
		}}
	
	node := flowgraph.MakeNode("pc", []*flowgraph.Edge{&incRail,&relRail,&absRail,&addrIn}, []*flowgraph.Edge{&addrOut}, rdyFunc, fireFunc)
	node.Aux = 0
	return node
}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}

func main() {
	
	
	flowgraph.ConfigByFlag(map[string]interface{}{ "ncore":4, "trace":"Q", "sec":4, "trsec":true})


	e,n := flowgraph.MakeGraph(5, 3)

	e[0].Name = "incRail"
	e[1].Name = "relRail"
	e[2].Name = "absRail"
        e[3].Name = "addrIn"
	e[4].Name = "addrOut"

	n[0] = tbi(e[0], e[1], e[2], e[3])
	n[1] = pc(e[0], e[1], e[2], e[3], e[4])
	n[2] = tbo(e[4])

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

