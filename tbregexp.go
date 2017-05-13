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
!    "AGGG[ACG]AAA|TTT[CGT]CCCT",
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

var prev map[int64]string;
		
func tbi(dnstreq flowgraph.Edge, newmatch flowgraph.Edge) flowgraph.Node {

        i := 0
	done := false

	f, err := os.Open("fasta.txt")
	check(err)
	r := bufio.NewReader(f)
	
	prev = make(map[int64]string)

	node := flowgraph.MakeNode("tbi", []*flowgraph.Edge{&dnstreq}, []*flowgraph.Edge{&newmatch},
		func (n *flowgraph.Node) bool {
			return !done && (dnstreq.SrcRdy(n) || newmatch.DstRdy(n) && i <= flowgraph.ChannelSize)
		},
		func (n *flowgraph.Node) {
			if dnstreq.SrcRdy(n) {
				match := dnstreq.SrcGet().(regexp.Search)
				if match.State == regexp.Done {
				        n.Tracef("DONE REQUEST FROM DOWNSTREAM %d\n", i-1);
				        delete(prev, match.ID)
					i--
				        return
				}
				n.Tracef("LIVE REQUEST FROM DOWNSTREAM %d\n", i);
				match.Curr = prev[match.ID][1:]
				prev[match.ID] = match.Curr
				newmatch.DstPut(match)
				return
			}
			xv,err := r.ReadString('\n')
			if err == io.EOF {
			        n.Tracef("EOF\n")
				done = true
			        return
			}
			id := regexp.NextID()
		        prev[id] = xv
			newmatch.DstPut(regexp.Search{Orig:xv, Curr:xv, State:regexp.Live, ID:id})
                        i++
		})
	return node
	
}

func tbo(oldmatch flowgraph.Edge, dnstreq flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&oldmatch}, []*flowgraph.Edge{&dnstreq}, nil,
		func (n *flowgraph.Node) {
			o := oldmatch.SrcGet().(regexp.Search)
			o.State = regexp.Done
			dnstreq.DstPut(o) // echo back
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


	// 1 match
        // e[test0].Const("AGGGTAAA")
        // e[test1].Const("TTTACCCT")
	
	// 0 match
	// e[test0].Const("[CGT]GGGTAAA")
        // e[test1].Const("TTTACCC[ACG]")
	
	// 0 match
	// e[test0].Const("A[ACT]GGTAAA")
        // e[test1].Const("TTTACC[AGT]T")
	
	// 0 match
	// e[test0].Const("AG[ACT]GTAAA")
        // e[test1].Const("TTTAC[AGT]CT")
	
	// 1 match
	// e[test0].Const("AGG[ACT]TAAA")
        // e[test1].Const("TTTA[AGT]CCT")
	
	// 0 match
	// e[test0].Const("AGGG[ACG]AAA")
        // e[test1].Const("TTT[CGT]CCCT")
	
	// 0 match
	// e[test0].Const("AGGGT[CGT]AA")
        // e[test1].Const("TT[ACG]ACCCT")
	
	// 0 match
	// e[test0].Const("AGGGTA[CGT]A")
        // e[test1].Const("T[ACG]TACCCT")
	
	// 2 match
	e[test0].Const("AGGGTAA[CGT]")
	e[test1].Const("[ACG]TTACCCT")
	
	n[0] = tbi(e[upstreq], e[newmatch])
	n[1] = regexp.FuncRepeat(e[newmatch], e[subsrc], e[dnstreq], e[oldmatch], e[subdst], e[upstreq], 1, -1)
	n[2] = regexp.FuncMatch(e[subdst], e[test0], e[subsrc0], /* ignoreCase = */ true)
	n[3] = regexp.FuncMatch(e[subdst], e[test1], e[subsrc1], /* ignoreCase = */ true)
	n[4] = regexp.FuncBar(e[subsrc0], e[subsrc1], e[subsrc], false)
	n[5] = tbo(e[oldmatch], e[dnstreq])
	
	flowgraph.RunAll(n)

}
