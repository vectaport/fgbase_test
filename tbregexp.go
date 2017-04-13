package main

import(
	"bufio"
	"io"
	"os"
	
	"github.com/vectaport/flowgraph"
        "github.com/vectaport/flowgraph/regexp"
)

var teststrings = []string{
	"axyz",
	"dxyz",
}

/*
var variants = []string{
    "AGGGTAAA|TTTACCCT",
    "[CGT]GGGTAAA|TTTACCC[ACG]",
    "A[ACT]GGTAAA|TTTACC[AGT]T",
    "AG[ACT]GTAAA|TTTAC[AGT]CT",
    "AGG[ACT]TAAA|TTTA[AGT]CCT",
    "AGGG[ACG]AAA|TTT[CGT]CCCT",
    "AGGGT[CGT]AA|TT[ACG]ACCCT",
    "AGGGTA[CGT]A|T[ACG]TACCCT",
    "AGGGTAA[CGT]|[ACG]TTACCCT",
}
*/

func check(e error) {
	if e != nil {
		flowgraph.StderrLog.Printf("%v\n", e)
		os.Exit(1)
	}
}
		
func tbi(dnstreq flowgraph.Edge, newmatch flowgraph.Edge) flowgraph.Node {

        i := 0
	done := false

	f, err := os.Open("fasta.txt")
	check(err)
	r := bufio.NewReader(f)
	
	prev := make(map[string]string)

	node := flowgraph.MakeNode("tbi", []*flowgraph.Edge{&dnstreq}, []*flowgraph.Edge{&newmatch},
		func (n *flowgraph.Node) bool {
			return !done && (dnstreq.SrcRdy(n) || newmatch.DstRdy(n) && i <= flowgraph.ChannelSize)
		},
		func (n *flowgraph.Node) {
			if dnstreq.SrcRdy(n) {
				match := dnstreq.SrcGet().(regexp.Search)
				if match.State == regexp.Done {
				        delete(prev, match.Orig)
					i--
				        return
				}
				match.Curr = prev[match.Orig][1:]
				prev[match.Orig] = match.Curr
				newmatch.DstPut(match)
				return
			}
			xv,err := r.ReadString('\n')
			if err == io.EOF {
				done = true
			        return
			}
		        prev[xv] = xv
			newmatch.DstPut(regexp.Search{Orig:xv, Curr:xv, State:regexp.Live})
                        i++
		})
	return node
	
}

func tbo(oldmatch flowgraph.Edge, dnstreq flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&oldmatch}, []*flowgraph.Edge{&dnstreq}, nil,
		func (n *flowgraph.Node) {
			o := oldmatch.SrcGet().(regexp.Search)
			dnstreq.DstPut(regexp.Search{State:regexp.Done, Orig:o.Orig}) // echo back
		})
	return node
         
}

type edgeCnt int
const (
	newmatch edgeCnt = iota
	subsrc
	dnstreq
	oldmatch
	subdst
	upstreq
	test0
	test1
	subsrc0
	subsrc1
	edgeNum
)

var edgeNames []string = []string {
	"newmatch",
	"subsrc",
	"dnstreq",
	"oldmatch",
	"subdst",
	"upstreq",
	"test0",
	"test1",
	"subsrc0",
	"subsrc1",
}

func main() {
	
	
	flowgraph.ConfigByFlag(nil)
	
	e,n := flowgraph.MakeGraph(int(edgeNum),6)
	flowgraph.NameEdges(e,edgeNames)

	// e[test0].Const("AGGGTAAA")
	// e[test1].Const("TTTACCCT")
	e[test0].Const("GGGAGGCCGAGGCGGGCGGATCACCTGAGGTCAGGAGTTCGAGACCAGCCTGGCCAA")
	e[test1].Const("GGGAGGCCGAGGCGGGCGGATCACCTGAGGTCAGGAGTTCGAGACCAGCCTGGCCAA")
	
	n[0] = tbi(e[upstreq], e[newmatch])
	n[1] = regexp.FuncRepeat(e[newmatch], e[subsrc], e[dnstreq], e[oldmatch], e[subdst], e[upstreq], 1, -1)
	n[2] = regexp.FuncMatch(e[subdst], e[test0], e[subsrc0])
	n[3] = regexp.FuncMatch(e[subdst], e[test1], e[subsrc1])
	n[4] = regexp.FuncBar(e[subsrc0], e[subsrc1], e[subsrc], false)
	n[5] = tbo(e[oldmatch], e[dnstreq])
	
	flowgraph.RunAll(n)
	
}
