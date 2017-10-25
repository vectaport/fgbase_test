package main

import (
	"flag"
	"math/rand"

	"github.com/vectaport/flowgraph"
	"github.com/vectaport/flowgraph/grid"
)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var MaxChanLen = 0

var tbbHz []float64
var tbbBase int64 = 0

func tbb(a flowgraph.Edge, x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbb", []*flowgraph.Edge{&a}, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			x.DstPut(n.NodeWrap(randSeq(16), x.Ack))
			if n.Cnt%100==0 {
				tbbHz[n.ID-tbbBase] = float64(n.Cnt)/flowgraph.TimeSinceStart()
			}})
	return node
}

func main() {
	
	nrowp := flag.Int("nrow", 4, "number of rows")
	ncolp := flag.Int("ncol", 4, "number of columns")
	flowgraph.ConfigByFlag(nil)
	nrow := *nrowp
	ncol := *ncolp
	
        fieldNodes := flowgraph.MakeNodes(nrow*ncol)
	topNodes := flowgraph.MakeNodes(ncol)
	botNodes := flowgraph.MakeNodes(ncol)
	lftNodes := flowgraph.MakeNodes(nrow)
	rgtNodes := flowgraph.MakeNodes(nrow)

        souEdges := flowgraph.MakeEdges((nrow+1)*ncol)
        norEdges := flowgraph.MakeEdges((nrow+1)*ncol)
        easEdges := flowgraph.MakeEdges((ncol+1)*nrow)
        wesEdges := flowgraph.MakeEdges((ncol+1)*nrow)

        for j:=0; j<nrow; j++ {
                for i:=0; i<ncol; i++ {
		        srcn := souEdges[i*(nrow+1)+j]
		        dsts := souEdges[i*(nrow+1)+j+1]
			
		        srcw := easEdges[j*(ncol+1)+i]
		        dste := easEdges[j*(ncol+1)+i+1]
			
		        srcs := norEdges[i*(nrow+1)+j+1]
		        dstn := norEdges[i*(nrow+1)+j]
			
		        srce := wesEdges[j*(ncol+1)+i+1]
		        dstw := wesEdges[j*(ncol+1)+i]
			
		        fieldNodes[i*ncol+j] = grid.FuncGrid(srcn, srce, srcs, srcw, dstn, dste, dsts, dstw)
		}
	}

        for i:=0; i<ncol; i++ {
	        topNodes[i] = tbb(norEdges[i*(nrow+1)], souEdges[i*(nrow+1)])
	        botNodes[i] = tbb(norEdges[i*(nrow+1)+ncol], souEdges[i*(nrow+1)+ncol])
	}

        for i:=0; i<nrow; i++ {
	        lftNodes[i] = tbb(wesEdges[i*(ncol+1)], easEdges[i*(ncol+1)])
	        rgtNodes[i] = tbb(wesEdges[i*(ncol+1)+nrow], easEdges[i*(ncol+1)+nrow])
	}

        
	tbbHz = make([]float64, nrow*2+ncol*2)

	var allNodes [] flowgraph.Node
	for i:= range fieldNodes {
	        allNodes = append(allNodes, fieldNodes[i])
	}
	for i:= range topNodes {
	        allNodes = append(allNodes, topNodes[i])
	}
	for i:= range botNodes {
	        allNodes = append(allNodes, botNodes[i])
	}
	for i:= range lftNodes {
	        allNodes = append(allNodes, lftNodes[i])
	}
	for i:= range rgtNodes {
	        allNodes = append(allNodes, rgtNodes[i])
	}

	flowgraph.RunAll(allNodes)

	// generate total frequency for tbb
	sum := 0.0
	for i:=0; i<len(tbbHz); i++ {
		sum += tbbHz[i]
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

