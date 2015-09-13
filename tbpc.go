package main

import (
	"os"

	"github.com/vectaport/flowgraph"
)

var pcHz []float64
var pcBase int64 = 1

const (
	incRail = iota
	absRail
	relRail
)

var pathIn = "pc_inputs.csv"
var pathOut = "pc_outputs.csv"

func check(e error) {
	if e != nil {
		flowgraph.StderrLog.Printf("%v\n", e)
		os.Exit(1)
	}
}
		
func pc(pcCtrl,addrIn,addrOut flowgraph.Edge) flowgraph.Node {

	var rdyFunc = func (n *flowgraph.Node) bool {
		if addrOut.DstRdy(n) {
			if pcCtrl.SrcRdy(n) {
				if pcCtrl.Val.(int)==incRail {
					n.RdyState = incRail
					addrIn.NoOut = true
					return true
				}
				if addrIn.SrcRdy(n) {
					n.RdyState = pcCtrl.Val.(int)
					return true
				}
			}
		}
		return false
	}

	var fireFunc = func (n *flowgraph.Node) { 
		if n.RdyState==absRail {
			p := addrIn.Val.(int)
			addrOut.Val = p
			n.Aux = p
		} else {
			addrOut.Val = n.Aux.(int)
			if n.RdyState==incRail {
				n.Aux = addrOut.Val.(int)+1
			} else {
				n.Aux = addrOut.Val.(int)+addrIn.Val.(int)
			}
		}
		
		if n.Cnt%1024==0 {
			pcHz[n.ID-pcBase] = float64(n.Cnt)/flowgraph.TimeSinceStart()
		}}

	node := flowgraph.MakeNode("pc", []*flowgraph.Edge{&pcCtrl,&addrIn}, []*flowgraph.Edge{&addrOut}, rdyFunc, fireFunc)
	node.Aux = 0
	return node
}

func main() {
	
	
	flowgraph.ConfigByFlag(map[string]interface{}{ "ncore":4, "trace":"Q", "sec":4})

	fi, err := os.Open(pathIn)
	check(err)
	fo, err := os.Open(pathOut)
	check(err)

	e,n := flowgraph.MakeGraph(3, 3)

	e[0].Name = "pcCtrl"
        e[1].Name = "addrIn"
	e[2].Name = "addrOut"

	n[0] = flowgraph.FuncCSVI(e[0:2], fi)
	n[1] = pc(e[0], e[1], e[2])
	n[2] = flowgraph.FuncCSVO(e[2:3], fo)

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
	// flowgraph.StdoutLog.Printf("%.2f%s", speed, hzstr)
}

